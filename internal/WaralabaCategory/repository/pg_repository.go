package repository

import (
	"context"
	category "warabiz/api/internal/models/category"
	"warabiz/api/pkg/infra/db"
	"warabiz/api/pkg/infra/logger"
)

type Repository interface {
	CreateCategory(ctx context.Context, db *db.DatabaseAccount, params ...interface{}) (int64, error)
	CheckDuplicateCategory(ctx context.Context, db *db.DatabaseAccount, category string) (bool, error)
	GetCategoryByID(ctx context.Context, db *db.DatabaseAccount, id int64) (*category.Category, error)
	GetAllCategory(ctx context.Context, db *db.DatabaseAccount, search *string, pageData *db.PageData) (*[]category.CategoryList, *db.PaginationResponse, error)
	UpdateCategory(ctx context.Context, db *db.DatabaseAccount, params ...interface{}) error
	DeleteCategory(ctx context.Context, db *db.DatabaseAccount, id int64) error
	GetCategoryDetailByID(ctx context.Context, db *db.DatabaseAccount, id int64) (*category.GetCategoryDetail, error)
}

type CategoryRepo struct {
	dbList []db.DatabaseAccount
	logger logger.Logger
}

func NewCategoryRepo(dbList []db.DatabaseAccount, logger logger.Logger) CategoryRepo {
	return CategoryRepo{
		dbList: dbList,
		logger: logger,
	}
}

func (r CategoryRepo) CreateCategory(ctx context.Context, db *db.DatabaseAccount, params ...interface{}) (int64, error) {
	var response int64
	return response, db.QueryRow(ctx, &response, qCreateCategory, params...)
}

func (r CategoryRepo) CheckDuplicateCategory(ctx context.Context, db *db.DatabaseAccount, category string) (bool, error) {
	var response bool
	return response, db.QueryRow(ctx, &response, qCheckDuplicateCategory, category)
}

func (r CategoryRepo) GetCategoryByID(ctx context.Context, db *db.DatabaseAccount, id int64) (*category.Category, error) {
	var response category.Category
	return &response, db.QueryRow(ctx, &response, qGetCategoryByID, id)
}

func (r CategoryRepo) GetAllCategory(ctx context.Context, db *db.DatabaseAccount, search *string, pageData *db.PageData) (*[]category.CategoryList, *db.PaginationResponse, error) {

	var response []category.CategoryList
	var total int
	var err error

	builder := db.NewQueryBuilder()

	//* Set Filter
	builder.Filter().
		AddSearch("ILIKE", search, "category_name")

	filter := builder.Filter().Build(false)
	filterArgs := builder.Filter().ExportArgs()

	//* Get count
	err = db.QueryRow(ctx, &total, qCountCategory+filter, filterArgs...)
	if err != nil {
		return nil, nil, err
	}

	//* Set Pagination
	builder.Paging().
		SetLimit(pageData.Page, pageData.Size).
		SetOrder(pageData.OrderBy, pageData.OrderType).
		SetDefaultLimit(1, 0).
		SetDefaultOrder("created_at", "DESC")

	page := builder.Paging().Build()
	pageArgs := builder.Paging().ExportArgs()

	//* Get response
	err = db.Query(ctx, &response, qGetAllCategory+filter+page, builder.JoinArgs(filterArgs, pageArgs)...)

	return &response, builder.Paging().ExportResponse(total), err
}

func (r CategoryRepo) UpdateCategory(ctx context.Context, db *db.DatabaseAccount, params ...interface{}) error {
	return db.Exec(ctx, qUpdateCategory, params...)
}

func (r CategoryRepo) DeleteCategory(ctx context.Context, db *db.DatabaseAccount, id int64) error {
	return db.Exec(ctx, qDeleteCategory, id)
}

func (r CategoryRepo) GetCategoryDetailByID(ctx context.Context, db *db.DatabaseAccount, id int64) (*category.GetCategoryDetail, error) {
	var response category.GetCategoryDetail
	return &response, db.QueryRow(ctx, &response, qGetCategoryDetailByID, id)
}
