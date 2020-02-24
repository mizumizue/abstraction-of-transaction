package main

import (
	"context"
	"log"

	"github.com/trewanek/abstraction-of-transaction/app/application/command"
	"github.com/trewanek/abstraction-of-transaction/app/application/unit_of_work"
	"github.com/trewanek/abstraction-of-transaction/app/application/usecase"
	"github.com/trewanek/abstraction-of-transaction/app/infrastructure/persistence/rdb"
	"github.com/trewanek/abstraction-of-transaction/app/interface/presenter"
)

func main() {
	ctx := context.Background()
	dbConn, err := rdb.NewDBConn()
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	uow, err := unit_of_work.NewUserUnitOfWork(ctx, dbConn)
	if err != nil {
		log.Fatal(err)
	}
	userRepo := rdb.NewUserMySqlRepository(dbConn)
	pre := presenter.NewStdoutPresenter()
	use := usecase.NewUserUseCase(uow, userRepo, pre)

	cmd := new(command.RegisterUserCommand)
	cmd.UserName = "new user"
	cmd.Email = "hoge@gmail.com"
	cmd.Telephone = "xxx-xxxx-xxxx"

	err = use.Register(ctx, cmd)
	if err != nil {
		log.Fatal(err)
	}
}
