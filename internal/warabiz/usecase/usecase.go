package usecase

import (
	"context"
	"fmt"
	"net/http"

	// "path/filepath"

	"warabiz/api/config"
	"warabiz/api/internal/models/warabiz"
	"warabiz/api/internal/warabiz/repository"
	"warabiz/api/pkg/http/exception"

	// "warabiz/api/pkg/http/locals"
	"warabiz/api/pkg/infra/db"
	"warabiz/api/pkg/infra/logger"
	"warabiz/api/pkg/utils/converter"
	"warabiz/api/pkg/validator"
	// "golang.org/x/exp/slices"
)

// const (
// 	maxFileMB   = 2
// 	maxFileSize = 1024 * 1024 * maxFileMB
// )

type Usecase interface {
	CreateWarabiz(ctx context.Context, exc exception.Exception, req *warabiz.CreateWarabizRequest) (*warabiz.CreateWarabizResponse, error)
	GetWarabizById(ctx context.Context, exc exception.Exception, id int64) (*warabiz.WarabizDetailResponse, error)
	GetAllWarabiz(ctx context.Context, exc exception.Exception, req *warabiz.GetAllWarabizRequest) (*warabiz.GetAllWarabizResponse, error)
	UpdateWarabiz(ctx context.Context, exc exception.Exception, req *warabiz.UpdateWarabizRequest) error
	DeleteWarabiz(ctx context.Context, exc exception.Exception, id int64) error
}

type WarabizUsecase struct {
	repo   repository.Repository
	cfg    *config.Config
	dbList []db.DatabaseAccount
	logger logger.Logger
}

func NewWarabizUsecase(repo repository.Repository, cfg *config.Config, dbList []db.DatabaseAccount, logger logger.Logger) Usecase {
	return WarabizUsecase{
		repo:   repo,
		cfg:    cfg,
		dbList: dbList,
		logger: logger,
	}
}

func (u WarabizUsecase) CreateWarabiz(ctx context.Context, exc exception.Exception, req *warabiz.CreateWarabizRequest) (*warabiz.CreateWarabizResponse, error) {

	var err error
	errMap := exception.NewMapErr()

	//* Validate Request
	errMap.JoinErrors(validator.ValidateStruct(req))
	if errMap.IsNotEmpty() {
		return nil, exc.NewRestError(http.StatusBadRequest, "validasi error", errMap)
	}

	// //* Validate file size
	// if req.Thumbnail.Size > maxFileSize {
	// 	return nil, exc.NewRestError(http.StatusBadRequest, fmt.Sprintf("ukuran file %s tidak boleh melebihi %v MB", req.Thumbnail.Filename, maxFileMB), nil)
	// }

	// //* Check Allowed File Extension
	// allowedExtensionFile := []string{".jpg", ".jpeg", ".png", ".heic", ".webp", ".svg"}

	// thumbnailExt := filepath.Ext(req.Thumbnail.Filename)
	// if !slices.Contains(allowedExtensionFile, thumbnailExt) {
	// 	return nil, exc.NewRestError(http.StatusBadRequest, fmt.Sprintf("jenis file '%s' tidak diizinkan untuk diunggah", thumbnailExt), nil)
	// }

	//* Sanitize Html
	req.Sanitize()

	//* DB Selector
	selectedDB, err := db.DBSelector(u.dbList, u.cfg.Connection.Warabiz.DriverSource)
	if err != nil {
		return nil, exc.NewRestError(http.StatusInternalServerError, "unkown database address", err.Error())
	}

	//* Check duplicate
	isDuplicate, err := u.repo.CheckDuplicateWarabiz(ctx, selectedDB, req.WaralabaName)
	if err != nil {
		return nil, exc.NewRestError(http.StatusInternalServerError, "failed check duplicate warabiz data", err.Error())
	}
	if isDuplicate {
		return nil, exc.NewRestError(http.StatusBadRequest, "data sudah ada", nil)
	}

	// //* Format Warabiz Title
	// formattedTitleWarabiz := strings.ReplaceAll(strings.ToLower(req.Title), " ", "-")

	// //* Get Creator
	// creator := locals.GetCreator(ctx)

	// //* GCP Repo Upload Mitra file
	// thumbnailUrl, err := u.gcpRepo.UploadFile(ctx, req.Thumbnail, getThumbnailName(formattedTitleWarabiz, thumbnailExt))
	// if err != nil {
	// 	return nil, exc.NewRestError(http.StatusInternalServerError, "failed to upload file", err.Error())
	// }

	//* Repo Create Warabiz
	params := make([]interface{}, 0)
	params = append(params, req.Id, req.CategoryId, req.WaralabaName, req.Prize, req.Contact, req.BrochureLink, req.Since, req.OutletTotal, req.LicenseDuration, "admin")
	warabizId, err := u.repo.CreateWarabiz(ctx, selectedDB, params...)
	if err != nil {
		return nil, exc.NewRestError(http.StatusInternalServerError, "failed to store warabiz data", err.Error())
	}
	if warabizId == 0 {
		return nil, exc.NewRestError(http.StatusInternalServerError, "failed to store warabiz data", nil)
	}

	//* Return success response
	return &warabiz.CreateWarabizResponse{
		Id: warabizId,
	}, nil
}

func (u WarabizUsecase) GetWarabizById(ctx context.Context, exc exception.Exception, id int64) (*warabiz.WarabizDetailResponse, error) {

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

	//* Repo get warabiz by id
	data, err := u.repo.GetWarabizByID(ctx, selectedDB, id)
	if err != nil {
		return nil, exc.NewRestError(http.StatusInternalServerError, "failed to find data", err.Error())
	}
	if data != nil {
		if data.Id == 0 {
			return nil, exc.NewRestError(http.StatusNotFound, "data tidak ditemukan", nil)
		}
	}

	data.CreatedAt = converter.ConvertTimeToLocal(data.CreatedAt, converter.DefaultLoc)
	if data.UpdatedAt != nil {
		*data.UpdatedAt = converter.ConvertTimeToLocal(*data.UpdatedAt, converter.DefaultLoc)
	}

	//* Return success response

	warabizDetail := &warabiz.WarabizDetailResponse{}

	return warabizDetail, nil
}

func (u WarabizUsecase) GetAllWarabiz(ctx context.Context, exc exception.Exception, req *warabiz.GetAllWarabizRequest) (*warabiz.GetAllWarabizResponse, error) {

	var err error

	//* DB Selector
	selectedDB, err := db.DBSelector(u.dbList, u.cfg.Connection.Warabiz.DriverSource)
	if err != nil {
		return nil, exc.NewRestError(http.StatusInternalServerError, "unkown database address", err.Error())
	}

	//* Repo get all warabiz
	warabizs, pageInfo, err := u.repo.GetAllWarabiz(ctx, selectedDB, req.WaralabaName, req.Search, &db.PageData{
		Page:      req.Page,
		Size:      req.PageSize,
		OrderBy:   req.OrderBy,
		OrderType: req.OrderType,
	})
	if err != nil {
		return nil, exc.NewRestError(http.StatusInternalServerError, "failed to find warabiz data", err.Error())
	}
	if warabizs != nil {
		if len(*warabizs) == 0 {
			exc.SetLog("warning_msg", "Search completed successfully. No matching data found based on the search criteria provided")
		}
	}

	//* Create success response
	return &warabiz.GetAllWarabizResponse{
		Warabiz:    *warabizs,
		Pagination: pageInfo,
	}, nil
}

func (u WarabizUsecase) UpdateWarabiz(ctx context.Context, exc exception.Exception, req *warabiz.UpdateWarabizRequest) error {

	var err error
	errMap := exception.NewMapErr()

	//* Validate Request
	errMap.JoinErrors(validator.ValidateStruct(req))
	if errMap.IsNotEmpty() {
		return exc.NewRestError(http.StatusBadRequest, "validasi error", errMap)
	}

	//* Sanitize Html
	req.Sanitize()

	// allowedExtensionFile := []string{".jpg", ".jpeg", ".png", ".heic", ".webp", ".svg"}
	// var thumbnailExt string

	// if req.IsUpdateThumbnail {
	// 	if req.Thumbnail == nil {
	// 		return exc.NewRestError(http.StatusBadRequest, "missing icon", nil)
	// 	}
	// 	//* Validate file size
	// 	if req.Thumbnail.Size > maxFileSize {
	// 		return exc.NewRestError(http.StatusBadRequest, fmt.Sprintf("ukuran file %s tidak boleh melebihi %v MB", req.Thumbnail.Filename, maxFileMB), nil)
	// 	}
	// 	//* Check Allowed File Extension
	// 	thumbnailExt = filepath.Ext(req.Thumbnail.Filename)
	// 	if !slices.Contains(allowedExtensionFile, thumbnailExt) {
	// 		return exc.NewRestError(http.StatusBadRequest, fmt.Sprintf("jenis file '%s' tidak diizinkan untuk diunggah", thumbnailExt), nil)
	// 	}
	// }

	//* DB Selector
	selectedDB, err := db.DBSelector(u.dbList, u.cfg.Connection.Warabiz.DriverSource)
	if err != nil {
		return exc.NewRestError(http.StatusInternalServerError, "unknown database address", err.Error())
	}

	//* Validate data
	warabiz, err := u.repo.GetWarabizByID(ctx, selectedDB, req.Id)
	if err != nil {
		return exc.NewRestError(http.StatusInternalServerError, "failed to find warabiz data", err.Error())
	}
	if warabiz != nil {
		if warabiz.Id == 0 {
			return exc.NewRestError(http.StatusBadRequest, fmt.Sprintf("warabiz dengan id %v tidak ditemukan", req.Id), nil)
		}
	}

	//* Check duplicate
	if req.WaralabaName != warabiz.WaralabaName {
		isDuplicate, err := u.repo.CheckDuplicateWarabiz(ctx, selectedDB, req.WaralabaName)
		if err != nil {
			return exc.NewRestError(http.StatusInternalServerError, "failed check duplicate warabiz data", err.Error())
		}
		if isDuplicate {
			return exc.NewRestError(http.StatusBadRequest, "duplikasi data terdeteksi", nil)
		}
	}

	// //* Get editor
	// editor := locals.GetCreator(ctx)

	//* Create success response
	return nil
}

func (u WarabizUsecase) DeleteWarabiz(ctx context.Context, exc exception.Exception, id int64) error {

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
	warabiz, err := u.repo.GetWarabizByID(ctx, selectedDB, id)
	if err != nil {
		return exc.NewRestError(http.StatusInternalServerError, "failed to find warabiz data", err.Error())
	}
	if warabiz != nil {
		if warabiz.Id == 0 {
			return exc.NewRestError(http.StatusBadRequest, fmt.Sprintf("warabiz dengan id %v tidak ditemukan", id), nil)
		}
	}

	//* Repo delete warabiz
	err = u.repo.DeleteWarabiz(ctx, selectedDB, id)
	if err != nil {
		return exc.NewRestError(http.StatusInternalServerError, "failed to delete warabiz data", err.Error())
	}

	//* Create success response
	return nil
}

// func getThumbnailName(formattedTitleWarabiz, thumbnailExt string) string {
// 	return fmt.Sprintf("%s/thumbnail%s", formattedTitleWarabiz, thumbnailExt)
// }
