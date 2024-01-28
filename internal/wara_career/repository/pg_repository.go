package repository

import (
	"context"
	WaraCareer "warabiz/api/internal/models/wara_career"
	"warabiz/api/pkg/infra/db"
	"warabiz/api/pkg/infra/logger"
)

type Repository interface {
	CreateWaraCareer(ctx context.Context, db *db.DatabaseAccount, params ...interface{}) (int64, error)
	CheckDuplicateWaraCareer(ctx context.Context, db *db.DatabaseAccount, waracareer string) (bool, error)
	GetWaraCareerByID(ctx context.Context, db *db.DatabaseAccount, id int64) (*WaraCareer.WaraCareer, error)
	GetAllWaraCareer(ctx context.Context, db *db.DatabaseAccount, search *string, pageData *db.PageData) (*[]WaraCareer.WaraCareerList, *db.PaginationResponse, error)
	UpdateWaraCareer(ctx context.Context, db *db.DatabaseAccount, params ...interface{}) error
	DeleteWaraCareer(ctx context.Context, db *db.DatabaseAccount, id int64) error
	// GetWaraCareerDetailByID(ctx context.Context, db *db.DatabaseAccount, id int64) (*waracareer.GetWaraCareerDetail, error)
}

type WaraCareerRepo struct {
	dbList []db.DatabaseAccount
	logger logger.Logger
}

func NewWaraCareerRepo(dbList []db.DatabaseAccount, logger logger.Logger) WaraCareerRepo {
	return WaraCareerRepo{
		dbList: dbList,
		logger: logger,
	}
}

func (r WaraCareerRepo) CreateWaraCareer(ctx context.Context, db *db.DatabaseAccount, params ...interface{}) (int64, error) {
	var response int64
	return response, db.QueryRow(ctx, &response, qCreateWaraCareer, params...)
}

func (r WaraCareerRepo) CheckDuplicateWaraCareer(ctx context.Context, db *db.DatabaseAccount, waracareer string) (bool, error) {
	var response bool
	return response, db.QueryRow(ctx, &response, qCheckDuplicateWaraCareer, waracareer)
}

func (r WaraCareerRepo) GetWaraCareerByID(ctx context.Context, db *db.DatabaseAccount, id int64) (*WaraCareer.WaraCareer, error) {
	var response WaraCareer.WaraCareer
	return &response, db.QueryRow(ctx, &response, qGetWaraCareerByID, id)
}

func (r WaraCareerRepo) GetAllWaraCareer(ctx context.Context, db *db.DatabaseAccount, search *string, pageData *db.PageData) (*[]WaraCareer.WaraCareerList, *db.PaginationResponse, error) {

	var response []WaraCareer.WaraCareerList
	var total int
	var err error

	builder := db.NewQueryBuilder()

	//* Set Filter
	builder.Filter().
		AddSearch("ILIKE", search, "career_titles")

	filter := builder.Filter().Build(false)
	filterArgs := builder.Filter().ExportArgs()

	//* Get count
	err = db.QueryRow(ctx, &total, qCountWaraCareer+filter, filterArgs...)
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
	err = db.Query(ctx, &response, qGetAllWaraCareer+filter+page, builder.JoinArgs(filterArgs, pageArgs)...)

	return &response, builder.Paging().ExportResponse(total), err
}

func (r WaraCareerRepo) UpdateWaraCareer(ctx context.Context, db *db.DatabaseAccount, params ...interface{}) error {
	return db.Exec(ctx, qUpdateWaraCareer, params...)
}

func (r WaraCareerRepo) DeleteWaraCareer(ctx context.Context, db *db.DatabaseAccount, id int64) error {
	return db.Exec(ctx, qDeleteWaraCareer, id)
}

// func (r WaraCareerRepo) GetWaraCareerDetailByID(ctx context.Context, db *db.DatabaseAccount, id int64) (*waracareer.GetWaraCareerDetail, error) {
// 	var response waracareer.GetWaraCareerDetail
// 	return &response, db.QueryRow(ctx, &response, qGetWaraCareerDetailByID, id)
// }
