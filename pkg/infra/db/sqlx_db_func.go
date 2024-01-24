package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/jmoiron/sqlx/reflectx"
)

const (
	tag = "json"
)

const (
	showQueryArgs = false
)

func show(query string, args ...interface{}) {
	if showQueryArgs {
		fmt.Println(query)
		fmt.Println(args...)
	}
}

func (dbAcc DatabaseAccount) Exec(ctx context.Context, query string, args ...interface{}) error {

	var err error
	show(query, args...)

	db := dbAcc.SqlDB
	if db == nil {
		return errors.New("sql client not found")
	}

	//* Tx version
	if dbAcc.IsTx() {
		return dbAcc.txExec(ctx, query, args...)
	}

	//* Translate ? in query
	q := db.Rebind(query)
	_, err = db.ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}
	return nil
}

func (dbAcc DatabaseAccount) QueryRow(ctx context.Context, response interface{}, query string, args ...interface{}) error {

	var err error
	show(query, args...)

	db := dbAcc.SqlDB
	if db == nil {
		return errors.New("sql client not found")
	}

	//* Tx version
	if dbAcc.IsTx() {
		return dbAcc.txQueryRow(ctx, response, query, args...)
	}

	//* Translate ? in query
	q := db.Rebind(query)
	//* Set json tag
	db.Mapper = reflectx.NewMapperFunc(tag, strings.ToLower)

	//* Type validation
	v := reflect.ValueOf(response)
	t := reflect.TypeOf(response)
	_, isTime := reflect.ValueOf(response).Interface().(*time.Time)

	if v.Kind() != reflect.Ptr {
		return errors.New("must pass a pointer, not a value, to StructScan destination")
	}
	if t.Elem().Kind() != reflect.Struct || isTime {
		err = db.QueryRowxContext(ctx, q, args...).Scan(response)
	} else {
		err = db.QueryRowxContext(ctx, q, args...).StructScan(response)
	}
	if err == sql.ErrNoRows {
		err = nil
	}
	return err
}

func (dbAcc DatabaseAccount) Query(ctx context.Context, response interface{}, query string, args ...interface{}) error {

	var err error
	show(query, args...)

	db := dbAcc.SqlDB
	if db == nil {
		return errors.New("sql client not found")
	}

	//* Tx version
	if dbAcc.IsTx() {
		return dbAcc.txQuery(ctx, response, query, args...)
	}

	//* Translate ? in query
	q := db.Rebind(query)
	//* Set json tag
	db.Mapper = reflectx.NewMapperFunc(tag, strings.ToLower)

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
	rows, err := db.QueryxContext(ctx, q, args...)
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
