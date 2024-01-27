package usecase

import (
	"context"
	"fmt"
	"net/http"
	"warabiz/api/config"
	"warabiz/api/internal/WaralabaCategory/repository"
	category "warabiz/api/internal/models/category"
	// waralabarepo "warabiz/api/internal/warabiz/repository"
	"warabiz/api/pkg/http/exception"

	"warabiz/api/pkg/infra/db"
	"warabiz/api/pkg/infra/logger"
	"warabiz/api/pkg/utils/converter"
	"warabiz/api/pkg/validator"
)

type Usecase interface {
	CreateCategory(ctx context.Context, exc exception.Exception, req *category.CreateCategoryRequest) (*category.CreateCategoryResponse, error)
	GetCategoryById(ctx context.Context, exc exception.Exception, id int64) (*category.Category, error)
	GetAllCategory(ctx context.Context, exc exception.Exception, req *category.GetAllCategoryRequest) (*category.GetAllCategoryResponse, error)
	UpdateCategory(ctx context.Context, exc exception.Exception, req *category.UpdateCategoryRequest) error
	DeleteCategory(ctx context.Context, exc exception.Exception, id int64) error
	// GetCategoryDetailById(ctx context.Context, exc exception.Exception, id int64) (*category.GetCategoryDetail, error)
}

type CategoryUsecase struct {
	repo   repository.Repository
	cfg    *config.Config
	dbList []db.DatabaseAccount
	logger logger.Logger
}

func NewCategoryUsecase(repo repository.Repository, cfg *config.Config, dbList []db.DatabaseAccount, logger logger.Logger) Usecase {
	return CategoryUsecase{
		repo:   repo,
		cfg:    cfg,
		dbList: dbList,
		logger: logger,
	}
}

func (u CategoryUsecase) CreateCategory(ctx context.Context, exc exception.Exception, req *category.CreateCategoryRequest) (*category.CreateCategoryResponse, error) {

	var err error
	errMap := exception.NewMapErr()

	//* Validate Request
	errMap.JoinErrors(validator.ValidateStruct(req))
	if errMap.IsNotEmpty() {
		return nil, exc.NewRestError(http.StatusBadRequest, "validasi error", errMap)
	}

	//* DB Selector
	selectedDB, err := db.DBSelector(u.dbList, u.cfg.Connection.Warabiz.DriverSource)
	if err != nil {
		return nil, exc.NewRestError(http.StatusInternalServerError, "unkown database address", err.Error())
	}

	//* Check duplicate
	isDuplicate, err := u.repo.CheckDuplicateCategory(ctx, selectedDB, req.CategoryName)
	if err != nil {
		return nil, exc.NewRestError(http.StatusInternalServerError, "failed check duplicate category data", err.Error())
	}
	if isDuplicate {
		return nil, exc.NewRestError(http.StatusBadRequest, "data sudah ada", nil)
	}

	// //* Get Creator
	// creator := locals.GetCreator(ctx)

	//* Repo Create Category
	params := make([]interface{}, 0)
	params = append(params, req.CategoryName, "admin")
	categoryId, err := u.repo.CreateCategory(ctx, selectedDB, params...)
	if err != nil {
		return nil, exc.NewRestError(http.StatusInternalServerError, "failed to store category data", err.Error())
	}
	if categoryId == 0 {
		return nil, exc.NewRestError(http.StatusInternalServerError, "failed to store category data", nil)
	}

	//* Return success response
	return &category.CreateCategoryResponse{
		Id: categoryId,
	}, nil
}

func (u CategoryUsecase) GetCategoryById(ctx context.Context, exc exception.Exception, id int64) (*category.Category, error) {

	var err error

	//* Validate Request
	if id == 0 {
		return nil, exc.NewRestError(http.StatusBadRequest, "id wajib diisi", nil)
	}

	//* DB Selector
	selectedDB, err := db.DBSelector(u.dbList, u.cfg.Connection.Warabiz.DriverSource)
	if err != nil {
		return nil, exc.NewRestError(http.StatusInternalServerError, "unkown database address", err.Error())
	}

	//* Repo get category by id
	category, err := u.repo.GetCategoryByID(ctx, selectedDB, id)
	if err != nil {
		return nil, exc.NewRestError(http.StatusInternalServerError, "failed to find category data", err.Error())
	}
	if category != nil {
		if category.Id == 0 {
			return nil, exc.NewRestError(http.StatusNotFound, "category tidak ditemukan", nil)
		}
	}

	category.CreatedAt = converter.ConvertTimeToLocal(category.CreatedAt, converter.DefaultLoc)
	if category.UpdatedAt != nil {
		*category.UpdatedAt = converter.ConvertTimeToLocal(*category.UpdatedAt, converter.DefaultLoc)
	}

	//* Return success response
	return category, nil
}

func (u CategoryUsecase) GetAllCategory(ctx context.Context, exc exception.Exception, req *category.GetAllCategoryRequest) (*category.GetAllCategoryResponse, error) {

	var err error

	//* DB Selector
	selectedDB, err := db.DBSelector(u.dbList, u.cfg.Connection.Warabiz.DriverSource)
	if err != nil {
		return nil, exc.NewRestError(http.StatusInternalServerError, "unkown database address", err.Error())
	}

	//* Repo get all category
	newcategories, pageInfo, err := u.repo.GetAllCategory(ctx, selectedDB, req.Search, &db.PageData{
		Page:      req.Page,
		Size:      req.PageSize,
		OrderBy:   req.OrderBy,
		OrderType: req.OrderType,
	})
	if err != nil {
		return nil, exc.NewRestError(http.StatusInternalServerError, "failed to find category data", err.Error())
	}
	if newcategories != nil {
		if len(*newcategories) == 0 {
			exc.SetLog("warning_msg", "Search completed successfully. No matching data found based on the search criteria provided")
		}
	}

	//* Create success response
	return &category.GetAllCategoryResponse{
		Category:   *newcategories,
		Pagination: pageInfo,
	}, nil
}

func (u CategoryUsecase) UpdateCategory(ctx context.Context, exc exception.Exception, req *category.UpdateCategoryRequest) error {

	var err error
	errMap := exception.NewMapErr()

	//* Validate Request
	errMap.JoinErrors(validator.ValidateStruct(req))
	if errMap.IsNotEmpty() {
		return exc.NewRestError(http.StatusBadRequest, "validasi error", errMap)
	}

	//* DB Selector
	selectedDB, err := db.DBSelector(u.dbList, u.cfg.Connection.Warabiz.DriverSource)
	if err != nil {
		return exc.NewRestError(http.StatusInternalServerError, "unknown database address", err.Error())
	}

	//* Validate data
	category, err := u.repo.GetCategoryByID(ctx, selectedDB, req.Id)
	if err != nil {
		return exc.NewRestError(http.StatusInternalServerError, "failed to find category data", err.Error())
	}
	if category != nil {
		if category.Id == 0 {
			return exc.NewRestError(http.StatusBadRequest, fmt.Sprintf("category dengan id %v tidak ditemukan", req.Id), nil)
		}
	}

	//* Check duplicate
	if req.CategoryName != category.CategoryName {
		isDuplicate, err := u.repo.CheckDuplicateCategory(ctx, selectedDB, req.CategoryName)
		if err != nil {
			return exc.NewRestError(http.StatusInternalServerError, "failed check duplicate category data", err.Error())
		}
		if isDuplicate {
			return exc.NewRestError(http.StatusBadRequest, "duplikasi data terdeteksi", nil)
		}
	}

	// //* Get editor
	// editor := locals.GetCreator(ctx)

	//* Repo update category
	params := make([]interface{}, 0)
	params = append(params, req.CategoryName, "Admin", req.Id)
	err = u.repo.UpdateCategory(ctx, selectedDB, params...)
	if err != nil {
		return exc.NewRestError(http.StatusInternalServerError, "failed to update category data", err.Error())
	}

	//* Create success response
	return nil
}

func (u CategoryUsecase) DeleteCategory(ctx context.Context, exc exception.Exception, id int64) error {

	var err error

	//* Validate Request
	if id == 0 {
		return exc.NewRestError(http.StatusBadRequest, "id wajib diisi", nil)
	}

	//* DB Selector
	selectedDB, err := db.DBSelector(u.dbList, u.cfg.Connection.Warabiz.DriverSource)
	if err != nil {
		return exc.NewRestError(http.StatusInternalServerError, "unkown database address", err.Error())
	}

	//* Validate data
	category, err := u.repo.GetCategoryByID(ctx, selectedDB, id)
	if err != nil {
		return exc.NewRestError(http.StatusInternalServerError, "failed to find category data", err.Error())
	}
	if category != nil {
		if category.Id == 0 {
			return exc.NewRestError(http.StatusBadRequest, fmt.Sprintf("category dengan id %v tidak ditemukan", id), nil)
		}
	}

	//* Repo delete category
	err = u.repo.DeleteCategory(ctx, selectedDB, id)
	if err != nil {
		return exc.NewRestError(http.StatusInternalServerError, "failed to delete category data", err.Error())
	}

	//* Create success response
	return nil
}

// func (u CategoryUsecase) GetCategoryDetailById(ctx context.Context, exc exception.Exception, id int64) (*category.GetCategoryDetail, error) {

// 	var err error

// 	//* Validate Request
// 	if id == 0 {
// 		return nil, exc.NewRestError(http.StatusBadRequest, "id wajib diisi", nil)
// 	}

// 	//* DB Selector
// 	selectedDB, err := db.DBSelector(u.dbList, u.cfg.Connection.Warabiz.DriverSource)
// 	if err != nil {
// 		return nil, exc.NewRestError(http.StatusInternalServerError, "unkown database address", err.Error())
// 	}

// 	//* Repo get categorydetail by id
// 	category, err := u.repo.GetCategoryDetailByID(ctx, selectedDB, id)
// 	if err != nil {
// 		return nil, exc.NewRestError(http.StatusInternalServerError, "failed to find category data", err.Error())
// 	}

// 	//* Get
// 	news, err := u.repo.GetByID(ctx, selectedDB, category.Id)
// 	if err != nil {
// 		return nil, exc.NewRestError(http.StatusInternalServerError, "failed to get news data", err.Error())
// 	}

// 	//* Get Category
// 	categories, err := u.repo.GetCategory(ctx, selectedDB, category.CategoryId)
// 	if err != nil {
// 		return nil, exc.NewRestError(http.StatusInternalServerError, "failed to get news category", err.Error())
// 	}

// 	//* Return success response
// 	response := category
// 	response.Category = categories

// 	return response, nil
// }
