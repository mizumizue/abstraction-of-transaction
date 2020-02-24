package repository

import (
	"context"

	"github.com/trewanek/abstraction-of-transaction/app/domain/entity"
)

type IUserRepository interface {
	FindAll(ctx context.Context) ([]*entity.User, error)
	Find(ctx context.Context, userID int) (*entity.User, error)
	Create(ctx context.Context, user *entity.User) error
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, userID int) error
}
