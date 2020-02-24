package service

import (
	"context"
	"database/sql"

	"github.com/trewanek/abstraction-of-transaction/app/domain/entity"
	"github.com/trewanek/abstraction-of-transaction/app/domain/repository"
)

type IUserService interface {
	Exists(ctx context.Context, user *entity.User) (bool, error)
}

type UserService struct {
	userRepo repository.IUserRepository
}

func NewUserService(userRepo repository.IUserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (service *UserService) Exists(ctx context.Context, user *entity.User) (bool, error) {
	found, err := service.userRepo.Find(ctx, user.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return found != nil, nil
}
