package db

import (
	"errors"
	"net/http"
	"warabiz/api/pkg/http/exception"

	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

type DatabaseAccount struct {
	Source     string
	SqlDB      *sqlx.DB
	SqlTx      *sqlx.Tx
	GormDB     *gorm.DB
	isCommited bool
}

func DBSelector(dbList []DatabaseAccount, source string) (*DatabaseAccount, error) {
	for _, dbAccount := range dbList {
		if dbAccount.Source == source {
			return &dbAccount, nil
		}
	}
	return nil, errors.New("failed to select db")
}

func (db *DatabaseAccount) IsTx() bool {
	return db.SqlTx != nil
}

func (db *DatabaseAccount) BeginTx() error {
	tx, err := db.SqlDB.Beginx()
	if err != nil {
		return err
	}
	db.SqlTx = tx
	return nil
}

func (db *DatabaseAccount) RecoverTx(exc exception.Exception, r interface{}) {
	if !db.IsTx() {
		return
	}
	if !db.isCommited {
		if exc != nil {
			if exc.IsError() {
				db.SqlTx.Rollback()
			}
		}
		if r != nil {
			db.SqlTx.Rollback()
		}
	}
}

func (db *DatabaseAccount) TxCommit(exc exception.Exception) error {
	if db.IsTx() {
		err := db.SqlTx.Commit()
		if err != nil {
			return exc.NewRestError(http.StatusInternalServerError, "failed to commit transaction", err.Error())
		}
		db.isCommited = true
	} else {
		return exc.NewRestError(http.StatusInternalServerError, "no transactions found", nil)
	}
	return nil
}
