package usecase

import (
	"context"
	"fmt"
	"net/http"
	"warabiz/api/config"
	WaraCareer "warabiz/api/internal/models/wara_career"
	"warabiz/api/internal/wara_career/repository"

	// waralabarepo "warabiz/api/internal/warabiz/repository"
	"warabiz/api/pkg/http/exception"

	"warabiz/api/pkg/infra/db"
	"warabiz/api/pkg/infra/logger"
	"warabiz/api/pkg/utils/converter"
	"warabiz/api/pkg/validator"
)

type Usecase interface {
	CreateWaraCareer(ctx context.Context, exc exception.Exception, req *WaraCareer.CreateWaraCareerRequest) (*WaraCareer.CreateWaraCareerResponse, error)
	GetWaraCareerById(ctx context.Context, exc exception.Exception, id int64) (*WaraCareer.WaraCareer, error)
	GetAllWaraCareer(ctx context.Context, exc exception.Exception, req *WaraCareer.GetAllWaraCareerRequest) (*WaraCareer.GetAllWaraCareerResponse, error)
	UpdateWaraCareer(ctx context.Context, exc exception.Exception, req *WaraCareer.UpdateWaraCareerRequest) error
	DeleteWaraCareer(ctx context.Context, exc exception.Exception, id int64) error
	// GetWaraCareerDetailById(ctx context.Context, exc exception.Exception, id int64) (*WaraCareer.GetWaraCareerDetail, error)
}

type WaraCareerUsecase struct {
	repo   repository.Repository
	cfg    *config.Config
	dbList []db.DatabaseAccount
	logger logger.Logger
}

func NewWaraCareerUsecase(repo repository.Repository, cfg *config.Config, dbList []db.DatabaseAccount, logger logger.Logger) Usecase {
	return WaraCareerUsecase{
		repo:   repo,
		cfg:    cfg,
		dbList: dbList,
		logger: logger,
	}
}

func (u WaraCareerUsecase) CreateWaraCareer(ctx context.Context, exc exception.Exception, req *WaraCareer.CreateWaraCareerRequest) (*WaraCareer.CreateWaraCareerResponse, error) {

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
		return nil, exc.NewRestError(http.StatusInternalServerError, "unknown database address", err.Error())
	}

	//* Check duplicate
	isDuplicate, err := u.repo.CheckDuplicateWaraCareer(ctx, selectedDB, req.CareerTitle)
	if err != nil {
		return nil, exc.NewRestError(http.StatusInternalServerError, "failed check duplicate WaraCareer data", err.Error())
	}
	if isDuplicate {
		return nil, exc.NewRestError(http.StatusBadRequest, "data sudah ada", nil)
	}

	// //* Get Creator
	// creator := locals.GetCreator(ctx)

	//* Repo Create WaraCareer
	params := make([]interface{}, 0)
	params = append(params, req.CareerTitle, req.Description, req.Address, req.ImageUrl, "admin")
	WaraCareerId, err := u.repo.CreateWaraCareer(ctx, selectedDB, params...)
	if err != nil {
		return nil, exc.NewRestError(http.StatusInternalServerError, "failed to store WaraCareer data", err.Error())
	}
	if WaraCareerId == 0 {
		return nil, exc.NewRestError(http.StatusInternalServerError, "failed to store WaraCareer data", nil)
	}

	//* Return success response
	return &WaraCareer.CreateWaraCareerResponse{
		Id: WaraCareerId,
	}, nil
}

func (u WaraCareerUsecase) GetWaraCareerById(ctx context.Context, exc exception.Exception, id int64) (*WaraCareer.WaraCareer, error) {

	var err error

	//* Validate Request
	if id == 0 {
		return nil, exc.NewRestError(http.StatusBadRequest, "id wajib diisi", nil)
	}

	//* DB Selector
	selectedDB, err := db.DBSelector(u.dbList, u.cfg.Connection.Warabiz.DriverSource)
	if err != nil {
		return nil, exc.NewRestError(http.StatusInternalServerError, "unknown database address", err.Error())
	}

	//* Repo get WaraCareer by id
	waraCareerDetail, err := u.repo.GetWaraCareerByID(ctx, selectedDB, id)
	if err != nil {
		return nil, exc.NewRestError(http.StatusInternalServerError, "failed to find WaraCareer data", err.Error())
	}
	if waraCareerDetail != nil {
		if waraCareerDetail.Id == 0 {
			return nil, exc.NewRestError(http.StatusNotFound, "WaraCareer tidak ditemukan", nil)
		}
	}

	waraCareerDetail.CreatedAt = converter.ConvertTimeToLocal(waraCareerDetail.CreatedAt, converter.DefaultLoc)
	if waraCareerDetail.UpdatedAt != nil {
		*waraCareerDetail.UpdatedAt = converter.ConvertTimeToLocal(*waraCareerDetail.UpdatedAt, converter.DefaultLoc)
	}

	//* Return success response
	return waraCareerDetail, nil
}

func (u WaraCareerUsecase) GetAllWaraCareer(ctx context.Context, exc exception.Exception, req *WaraCareer.GetAllWaraCareerRequest) (*WaraCareer.GetAllWaraCareerResponse, error) {

	var err error

	//* DB Selector
	selectedDB, err := db.DBSelector(u.dbList, u.cfg.Connection.Warabiz.DriverSource)
	if err != nil {
		return nil, exc.NewRestError(http.StatusInternalServerError, "unknown database address", err.Error())
	}

	//* Repo get all WaraCareer
	newcategories, pageInfo, err := u.repo.GetAllWaraCareer(ctx, selectedDB, req.Search, &db.PageData{
		Page:      req.Page,
		Size:      req.PageSize,
		OrderBy:   req.OrderBy,
		OrderType: req.OrderType,
	})
	if err != nil {
		return nil, exc.NewRestError(http.StatusInternalServerError, "failed to find WaraCareer data", err.Error())
	}
	if newcategories != nil {
		if len(*newcategories) == 0 {
			exc.SetLog("warning_msg", "Search completed successfully. No matching data found based on the search criteria provided")
		}
	}

	//* Create success response
	return &WaraCareer.GetAllWaraCareerResponse{
		WaraCareer: *newcategories,
		Pagination: pageInfo,
	}, nil
}

func (u WaraCareerUsecase) UpdateWaraCareer(ctx context.Context, exc exception.Exception, req *WaraCareer.UpdateWaraCareerRequest) error {

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
	WaraCareer, err := u.repo.GetWaraCareerByID(ctx, selectedDB, req.Id)
	if err != nil {
		return exc.NewRestError(http.StatusInternalServerError, "failed to find WaraCareer data", err.Error())
	}
	if WaraCareer != nil {
		if WaraCareer.Id == 0 {
			return exc.NewRestError(http.StatusBadRequest, fmt.Sprintf("WaraCareer dengan id %v tidak ditemukan", req.Id), nil)
		}
	}

	//* Check duplicate
	if req.CareerTitle != WaraCareer.CareerTitle {
		isDuplicate, err := u.repo.CheckDuplicateWaraCareer(ctx, selectedDB, req.CareerTitle)
		if err != nil {
			return exc.NewRestError(http.StatusInternalServerError, "failed check duplicate WaraCareer data", err.Error())
		}
		if isDuplicate {
			return exc.NewRestError(http.StatusBadRequest, "duplikasi data terdeteksi", nil)
		}
	}

	// //* Get editor
	// editor := locals.GetCreator(ctx)

	//* Repo update WaraCareer
	params := make([]interface{}, 0)
	params = append(params, req.CareerTitle, "Admin", req.Id)
	err = u.repo.UpdateWaraCareer(ctx, selectedDB, params...)
	if err != nil {
		return exc.NewRestError(http.StatusInternalServerError, "failed to update WaraCareer data", err.Error())
	}

	//* Create success response
	return nil
}

func (u WaraCareerUsecase) DeleteWaraCareer(ctx context.Context, exc exception.Exception, id int64) error {

	var err error

	//* Validate Request
	if id == 0 {
		return exc.NewRestError(http.StatusBadRequest, "id wajib diisi", nil)
	}

	//* DB Selector
	selectedDB, err := db.DBSelector(u.dbList, u.cfg.Connection.Warabiz.DriverSource)
	if err != nil {
		return exc.NewRestError(http.StatusInternalServerError, "unknown database address", err.Error())
	}

	//* Validate data
	WaraCareer, err := u.repo.GetWaraCareerByID(ctx, selectedDB, id)
	if err != nil {
		return exc.NewRestError(http.StatusInternalServerError, "failed to find WaraCareer data", err.Error())
	}
	if WaraCareer != nil {
		if WaraCareer.Id == 0 {
			return exc.NewRestError(http.StatusBadRequest, fmt.Sprintf("WaraCareer dengan id %v tidak ditemukan", id), nil)
		}
	}

	//* Repo delete WaraCareer
	err = u.repo.DeleteWaraCareer(ctx, selectedDB, id)
	if err != nil {
		return exc.NewRestError(http.StatusInternalServerError, "failed to delete WaraCareer data", err.Error())
	}

	//* Create success response
	return nil
}

// func (u WaraCareerUsecase) GetWaraCareerDetailById(ctx context.Context, exc exception.Exception, id int64) (*WaraCareer.GetWaraCareerDetail, error) {

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

// 	//* Repo get WaraCareerdetail by id
// 	WaraCareer, err := u.repo.GetWaraCareerDetailByID(ctx, selectedDB, id)
// 	if err != nil {
// 		return nil, exc.NewRestError(http.StatusInternalServerError, "failed to find WaraCareer data", err.Error())
// 	}

// 	//* Get
// 	news, err := u.repo.GetByID(ctx, selectedDB, WaraCareer.Id)
// 	if err != nil {
// 		return nil, exc.NewRestError(http.StatusInternalServerError, "failed to get news data", err.Error())
// 	}

// 	//* Get WaraCareer
// 	categories, err := u.repo.GetWaraCareer(ctx, selectedDB, WaraCareer.WaraCareerId)
// 	if err != nil {
// 		return nil, exc.NewRestError(http.StatusInternalServerError, "failed to get news WaraCareer", err.Error())
// 	}

// 	//* Return success response
// 	response := WaraCareer
// 	response.WaraCareer = categories

// 	return response, nil
// }
