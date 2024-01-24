package db

import (
	"fmt"
	"math"
	"regexp"
)

type builderItf interface {
	Filter() *filterData
	Paging() *pageData
	JoinArgs(args ...[]interface{}) []interface{}
}

type builderData struct {
	filter filterData
	paging pageData
}

type filter struct {
	key      string
	value    interface{}
	operator string
}

type search struct {
	value    interface{}
	target   []string
	operator string
}

type filterData struct {
	filter []filter
	search []search
	args   []interface{}
}

type PageData struct {
	Page      int    `json:"page" query:"page"`
	Size      int    `json:"size" query:"size"`
	OrderBy   string `json:"order_by" query:"orderBy"`
	OrderType string `json:"order_type" query:"orderType"`
}

type PaginationResponse struct {
	TotalRow  int    `json:"total_row"`
	TotalPage int    `json:"total_page"`
	Page      int    `json:"page"`
	Size      int    `json:"size"`
	HasMore   bool   `json:"has_more"`
	OrderBy   string `json:"order_by,omitempty"`
	OrderType string `json:"order_type,omitempty"`
}

type pageData struct {
	page                PageData
	defaultValue        PageData
	args                []interface{}
	isContainDangerChar bool
}

func (dbAcc DatabaseAccount) NewQueryBuilder() builderItf {
	return &builderData{}
}

func (b *builderData) Filter() *filterData {
	return &b.filter
}

func (f *filterData) AddFilter(key, operator string, value interface{}) *filterData {
	f.filter = append(f.filter, filter{
		key:      key,
		operator: operator,
		value:    value,
	})
	return f
}

func (f *filterData) AddSearch(operator string, value interface{}, target ...string) *filterData {
	f.search = append(f.search, search{
		operator: operator,
		value:    value,
		target:   append(make([]string, 0), target...),
	})
	return f
}

func (b *builderData) Paging() *pageData {
	return &b.paging
}

func (p *pageData) SetLimit(page, size int) *pageData {
	p.page.Page = page
	p.page.Size = size
	return p
}

func (p *pageData) SetOrder(orderBy, orderType string) *pageData {
	p.page.OrderBy = orderBy
	p.page.OrderType = orderType
	return p
}

func (p *pageData) SetDefaultLimit(page, size int) *pageData {
	p.defaultValue.Page = page
	p.defaultValue.Size = size
	return p
}

func (p *pageData) SetDefaultOrder(orderBy, orderType string) *pageData {
	p.defaultValue.OrderBy = orderBy
	p.defaultValue.OrderType = orderType
	return p
}

func (f *filterData) Build(isStandalone bool) string {

	var (
		filter   string
		search   string
		counter1 int
		counter2 int
		result   string
		where    string = "WHERE"
	)

	//* Filter
	fLoop := len(f.filter)
	if fLoop != 0 {
		isFirst := true
		isNull := true
		for _, ft := range f.filter {
			if ft.value == nil || fmt.Sprintf("%v", ft.value) == "<nil>" {
				continue
			} else {
				isNull = false
			}

			if ft.operator == "ILIKE" || ft.operator == "LIKE" {
				newValue, ok := ft.value.(*string)
				if !ok {
					continue
				}
				ft.value = ("%" + fmt.Sprintf("%v", newValue) + "%")
			}

			if isFirst {
				filter += fmt.Sprintf(` %s %s ?`, ft.key, ft.operator)
			} else {
				filter += fmt.Sprintf(` AND %s %s ?`, ft.key, ft.operator)
			}

			f.args = append(f.args, ft.value)

			if isFirst {
				isFirst = false
			}
		}
		if !isNull {
			counter1 += 1
		}
	}

	//* Search
	sLoop := len(f.search)
	if sLoop != 0 {
		isNull := true
		for _, s := range f.search {
			if s.value == nil || fmt.Sprintf("%v", s.value) == "<nil>" || len(s.target) == 0 {
				continue
			} else {
				isNull = false
			}

			if s.operator == "ILIKE" || s.operator == "LIKE" {
				newValue, ok := s.value.(*string)
				if !ok {
					continue
				}
				s.value = ("%" + fmt.Sprintf("%v", *newValue) + "%")
			}

			tLoop := len(s.target)
			if tLoop != 0 && s.operator != "" {
				for i := 0; i < tLoop; i++ {
					if i == 0 {
						search += fmt.Sprintf(` %s %s ?`, s.target[i], s.operator)
					} else {
						search += fmt.Sprintf(` OR %s %s ?`, s.target[i], s.operator)
					}
					f.args = append(f.args, s.value)
				}
			}
		}
		if !isNull {
			counter2 += 1
		}
	}

	if isStandalone {
		result += where
	} else {
		result += "AND"
	}
	if counter1+counter2 == 0 {
		return ""
	}
	if counter1 == 1 {
		result += filter
	}
	if counter1+counter2 == 2 {
		result += " AND"
	}
	if counter2 == 1 {
		result += fmt.Sprintf(" (%s)", search)
	}

	return result
}

func (f *filterData) ExportArgs() []interface{} {
	return f.args
}

func (p *pageData) Build() string {

	var result string

	if p.page.OrderBy == "" {
		p.page.OrderBy = p.defaultValue.OrderBy
	}
	if p.page.OrderType == "" {
		p.page.OrderType = p.defaultValue.OrderType
	}

	re := regexp.MustCompile(`[@#$%^&=~*()[\]:;'\"?/><,{}|\\ ]+`)
	if !re.MatchString(p.page.OrderBy) && !re.MatchString(p.page.OrderType) && len(p.page.OrderBy) <= 63 && len(p.page.OrderType) <= 4 {
		if p.page.OrderBy != "" && p.page.OrderType != "" {
			result += fmt.Sprintf(" ORDER BY %s %s", p.page.OrderBy, p.page.OrderType)
		}
	} else {
		p.isContainDangerChar = true
	}

	if p.page.Page == 0 {
		p.page.Page = p.defaultValue.Page
	}
	if p.page.Size == 0 {
		p.page.Size = p.defaultValue.Size
	}

	limit := p.page.Size
	offset := (p.page.Page - 1) * p.page.Size

	if p.page.Page != 0 && p.page.Size != 0 {
		result += " LIMIT ? OFFSET ?"
		p.args = append(p.args, limit, offset)
	}

	return result
}

func (p *pageData) ExportArgs() []interface{} {
	return p.args
}

func (p *pageData) ExportResponse(totalRow int) *PaginationResponse {

	var totalPage int
	if p.page.Size > 0 {
		totalPage = int(math.Ceil(float64(totalRow) / float64(p.page.Size)))
	} else {
		totalPage = 1
	}
	if totalPage == 0 {
		totalPage = 1
	}

	var hasMore bool
	if totalPage-p.page.Page < 1 {
		hasMore = false
	} else {
		hasMore = true
	}

	if p.isContainDangerChar {
		p.page.OrderBy = ""
		p.page.OrderType = ""
	}

	return &PaginationResponse{
		TotalRow:  totalRow,
		TotalPage: totalPage,
		Page:      p.page.Page,
		Size:      p.page.Size,
		HasMore:   hasMore,
		OrderBy:   p.page.OrderBy,
		OrderType: p.page.OrderType,
	}
}

func (b *builderData) JoinArgs(args ...[]interface{}) []interface{} {
	newArgs := make([]interface{}, 0)
	for _, a := range args {
		newArgs = append(newArgs, a...)
	}
	return newArgs
}
