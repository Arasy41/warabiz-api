package db

import (
	"context"
	"database/sql"
	"errors"
	"reflect"
	"strings"
	"time"

	"github.com/jmoiron/sqlx/reflectx"
)

func (dbAcc DatabaseAccount) txExec(ctx context.Context, query string, args ...interface{}) error {

	var err error
	tx := dbAcc.SqlTx

	//* Translate ? in query
	q := tx.Rebind(query)
	_, err = tx.ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}
	return nil
}

func (dbAcc DatabaseAccount) txQueryRow(ctx context.Context, response interface{}, query string, args ...interface{}) error {

	var err error
	tx := dbAcc.SqlTx

	//* Translate ? in query
	q := tx.Rebind(query)
	//* Set json tag
	tx.Mapper = reflectx.NewMapperFunc(tag, strings.ToLower)

	//* Type validation
	v := reflect.ValueOf(response)
	t := reflect.TypeOf(response)
	_, isTime := reflect.ValueOf(response).Interface().(*time.Time)

	if v.Kind() != reflect.Ptr {
		return errors.New("must pass a pointer, not a value, to StructScan destination")
	}
	if t.Elem().Kind() != reflect.Struct || isTime {
		err = tx.QueryRowxContext(ctx, q, args...).Scan(response)
	} else {
		err = tx.QueryRowxContext(ctx, q, args...).StructScan(response)
	}
	if err == sql.ErrNoRows {
		err = nil
	}
	return err
}

func (dbAcc DatabaseAccount) txQuery(ctx context.Context, response interface{}, query string, args ...interface{}) error {

	var err error
	tx := dbAcc.SqlTx

	//* Translate ? in query
	q := tx.Rebind(query)
	//* Set json tag
	tx.Mapper = reflectx.NewMapperFunc(tag, strings.ToLower)

	//* Type validation
	v := reflect.ValueOf(response)
	t := reflect.TypeOf(response)

	if v.Kind() != reflect.Ptr {
		return errors.New("must pass a pointer, not a value, to StructScan destination")
	}
	if t.Elem().Kind() != reflect.Slice {
		return errors.New("must pass a pointer of slice")
	}

	resType := t.Elem().Elem()
	sliceRes := reflect.New(t.Elem()).Elem()

	//* Get rows data from db
	rows, err := tx.QueryxContext(ctx, q, args...)
	if rows != nil {
		defer rows.Close()
	}
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	for rows.Next() {

		res := reflect.New(resType)
		_, isTime := v.Interface().(*[]time.Time)

		if resType.Kind() != reflect.Struct || isTime {
			err = rows.Scan(res.Elem().Addr().Interface())
		} else {
			err = rows.StructScan(res.Elem().Addr().Interface())
		}
		if err != nil && err != sql.ErrNoRows {
			return err
		}
		sliceRes = reflect.Append(sliceRes, res.Elem())
	}
	reflect.ValueOf(response).Elem().Set(sliceRes)
	return err
}
