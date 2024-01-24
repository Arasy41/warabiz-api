package repository

import (
	"context"
	"warabiz/api/internal/models/warabiz"
	"warabiz/api/pkg/infra/db"
	"warabiz/api/pkg/infra/logger"
)

type Repository interface {
	CreateWarabiz(ctx context.Context, db *db.DatabaseAccount, params ...interface{}) (int64, error)
	CheckDuplicateWarabiz(ctx context.Context, db *db.DatabaseAccount, title string) (bool, error)
	GetWarabizByID(ctx context.Context, db *db.DatabaseAccount, id int64) (*warabiz.Warabiz, error)
	GetAllWarabiz(ctx context.Context, db *db.DatabaseAccount, fWaraName *string, search *string, pageData *db.PageData) (*[]warabiz.WarabizList, *db.PaginationResponse, error)
	UpdateWarabiz(ctx context.Context, db *db.DatabaseAccount, params ...interface{}) error
	DeleteWarabiz(ctx context.Context, db *db.DatabaseAccount, id int64) error
}

type WarabizRepo struct {
	dbList []db.DatabaseAccount
	logger logger.Logger
}

func NewWarabizRepo(dbList []db.DatabaseAccount, logger logger.Logger) WarabizRepo {
	return WarabizRepo{
		dbList: dbList,
		logger: logger,
	}
}

func (r WarabizRepo) CreateWarabiz(ctx context.Context, db *db.DatabaseAccount, params ...interface{}) (int64, error) {
	var response int64
	return response, db.QueryRow(ctx, &response, qCreateWarabiz, params...)
}

func (r WarabizRepo) CheckDuplicateWarabiz(ctx context.Context, db *db.DatabaseAccount, title string) (bool, error) {
	var response bool
	return response, db.QueryRow(ctx, &response, qCheckDuplicateWarabiz, title)
}

func (r WarabizRepo) GetWarabizByID(ctx context.Context, db *db.DatabaseAccount, id int64) (*warabiz.Warabiz, error) {
	var response warabiz.Warabiz
	return &response, db.QueryRow(ctx, &response, qGetWarabizByID, id)
}

func (r WarabizRepo) GetAllWarabiz(ctx context.Context, db *db.DatabaseAccount, fWaraName *string, search *string, pageData *db.PageData) (*[]warabiz.WarabizList, *db.PaginationResponse, error) {

	var response []warabiz.WarabizList
	var total int
	var err error

	builder := db.NewQueryBuilder()

	//* Set Filter
	builder.Filter().
		AddFilter("waralaba_name", "=", fWaraName).
		AddSearch("ILIKE", search, "waralaba_name")

	filter := builder.Filter().Build(false)
	filterArgs := builder.Filter().ExportArgs()

	//* Get count
	err = db.QueryRow(ctx, &total, qCountWarabiz+filter, filterArgs...)
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
	err = db.Query(ctx, &response, qGetAllWarabiz+filter+page, builder.JoinArgs(filterArgs, pageArgs)...)

	return &response, builder.Paging().ExportResponse(total), err
}

func (r WarabizRepo) UpdateWarabiz(ctx context.Context, db *db.DatabaseAccount, params ...interface{}) error {
	return db.Exec(ctx, qUpdateWarabiz, params...)
}

func (r WarabizRepo) DeleteWarabiz(ctx context.Context, db *db.DatabaseAccount, id int64) error {
	return db.Exec(ctx, qDeleteWarabiz, id)
}
