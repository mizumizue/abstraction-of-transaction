package usecase

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/trewanek/abstraction-of-transaction/app/application/command"
	"github.com/trewanek/abstraction-of-transaction/app/application/unit_of_work"
	"github.com/trewanek/abstraction-of-transaction/app/domain/entity"
	"github.com/trewanek/abstraction-of-transaction/app/domain/repository"
	"github.com/trewanek/abstraction-of-transaction/app/domain/service"
	"github.com/trewanek/abstraction-of-transaction/app/interface/presenter"
)

type IUserUseCase interface {
	FetchAll(ctx context.Context, user *entity.User) error
	Find(ctx context.Context, user *entity.User) error
	Register(ctx context.Context, user *entity.User) error
	Update(ctx context.Context, user *entity.User) error
	Remove(ctx context.Context, user *entity.User) error
}

type UserUseCase struct {
	uow      *unit_of_work.UserUnitOfWork
	userSer  service.IUserService
	userRepo repository.IUserRepository
	pre      presenter.IPresenter
}

func NewUserUseCase(
	uow *unit_of_work.UserUnitOfWork,
	userRepo repository.IUserRepository,
	pre presenter.IPresenter,
) *UserUseCase {
	return &UserUseCase{
		uow:      uow,
		userSer:  service.NewUserService(userRepo),
		userRepo: userRepo,
		pre:      pre,
	}
}

func (u *UserUseCase) FetchAll(ctx context.Context, user *entity.User) error {
	panic("implement me")
}

func (u *UserUseCase) Find(ctx context.Context, user *entity.User) error {
	panic("implement me")
}

func (u *UserUseCase) Register(
	ctx context.Context,
	command *command.RegisterUserCommand,
) error {
	insertedUserID, err := func(uow *unit_of_work.UserUnitOfWork) (int, error) {
		src := rand.NewSource(time.Now().UnixNano())
		newUser := &entity.User{
			UserID:    rand.New(src).Int(),
			UserName:  command.UserName,
			Email:     command.Email,
			Telephone: command.Telephone,
		}

		exists, err := u.userSer.Exists(ctx, newUser)
		if err != nil {
			return 0, fmt.Errorf("database err: %w", err)
		}
		if exists {
			return 0, fmt.Errorf("user is duplicate")
		}

		if err = uow.UserRepository.Create(ctx, newUser); err != nil {
			return 0, err
		}

		return newUser.UserID, nil
	}(u.uow)
	if err != nil {
		u.uow.Rollback()
		log.Fatalf("rollback detail: %v", err)
	}
	u.uow.Commit()

	user, err := u.userRepo.Find(ctx, insertedUserID)
	if err != nil {
		return err
	}

	if err := u.pre.Present(user); err != nil {
		return err
	}

	return nil
}

func (u *UserUseCase) Update(ctx context.Context, user *entity.User) error {
	panic("implement me")
}

func (u *UserUseCase) Remove(ctx context.Context, user *entity.User) error {
	panic("implement me")
}
