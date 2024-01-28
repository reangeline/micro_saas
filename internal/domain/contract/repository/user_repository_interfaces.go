package repository

import (
	"context"

	"github.com/reangeline/micro_saas/internal/infra/database/model"
)

type UserRepositoryInterface interface {
	CreateUser(ctx context.Context, user *model.UserModel) error
	FindAll() ([]*model.UserModel, error)
	FindByUserEmail(email string) (*model.UserModel, error)
	UpdateByEmail(input *model.UserModel) (*model.UserModel, error)
}
