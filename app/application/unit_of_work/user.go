package unit_of_work

import (
	"context"

	"github.com/trewanek/abstraction-of-transaction/app/domain/repository"
	"github.com/trewanek/abstraction-of-transaction/app/infrastructure/persistence/rdb"
)

type UserUnitOfWork struct {
	dbConn         *rdb.DBConn
	UserRepository repository.IUserRepository
}

func NewUserUnitOfWork(ctx context.Context, conn *rdb.DBConn) (*UserUnitOfWork, error) {
	err := conn.Begin(ctx)
	if err != nil {
		return nil, err
	}
	return &UserUnitOfWork{
		dbConn:         conn,
		UserRepository: rdb.NewUserMySqlRepository(conn),
	}, nil
}

func (uow *UserUnitOfWork) Rollback() {
	uow.dbConn.Rollback()
}

func (uow *UserUnitOfWork) Commit() {
	uow.dbConn.Commit()
}
