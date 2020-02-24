package rdb

import (
	"context"
	"database/sql"
	"reflect"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DBConn struct {
	db *sqlx.DB
	tx *sqlx.Tx
}

func NewDBConn() (*DBConn, error) {
	dbConn, err := sqlx.Connect("mysql", "mysql:password@/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	}
	return &DBConn{
		db: dbConn,
	}, nil
}

func (conn *DBConn) Begin(ctx context.Context) error {
	if conn.inTransaction() {
		return nil
	}
	tx, err := conn.db.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}
	conn.tx = tx
	return nil
}

func (conn *DBConn) Rollback() {
	if conn.inTransaction() {
		conn.tx.Rollback()
		conn.tx = nil
	}
}

func (conn *DBConn) Commit() {
	if conn.inTransaction() {
		conn.tx.Commit()
		conn.tx = nil
	}
}

func (conn *DBConn) inTransaction() bool {
	return conn.tx != nil
}

func (conn *DBConn) Close() error {
	return conn.db.Close()
}

func (conn *DBConn) get(ctx context.Context, dst interface{}, query string, args ...interface{}) error {
	if conn.inTransaction() {
		return conn.tx.GetContext(ctx, dst, query, args...)
	}
	return conn.db.GetContext(ctx, dst, query, args...)
}

func (conn *DBConn) query(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error) {
	if conn.inTransaction() {
		return conn.tx.QueryxContext(ctx, query, args...)
	}
	return conn.db.QueryxContext(ctx, query, args...)
}

func (conn *DBConn) exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	if conn.inTransaction() {
		return conn.tx.ExecContext(ctx, query, args...)
	}
	return conn.db.ExecContext(ctx, query, args...)
}

func (conn *DBConn) namedQuery(ctx context.Context, query string, namedParam map[string]interface{}) (*sqlx.Rows, error) {
	panic("implement me")
}

func (conn *DBConn) namedExec(ctx context.Context, query string, namedParam map[string]interface{}) (sql.Result, error) {
	panic("implement me")
}

func (conn *DBConn) rowsScan(ctx context.Context, rows *sqlx.Rows, typeI interface{}) (interface{}, error) {
	defer func() {
		if rows != nil {
			return
		}
		_ = rows.Close()
	}()

	sliceType := reflect.SliceOf(reflect.PtrTo(reflect.TypeOf(typeI)))
	list := reflect.New(sliceType).Elem()

	for rows.Next() {
		dst := reflect.New(reflect.TypeOf(typeI))
		dstI := dst.Interface()
		if err := rows.StructScan(dstI); err != nil {
			return nil, err
		}
		list = reflect.Append(list, dst)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return list.Interface(), nil
}
