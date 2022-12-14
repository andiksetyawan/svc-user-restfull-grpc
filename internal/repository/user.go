package repository

import (
	"context"

	"gorm.io/gorm"
	"svc-user/internal/entity"
)

type IUserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
}

type userRepository struct {
	Db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{
		Db: db,
	}
}

func (u userRepository) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	if err := u.Db.WithContext(ctx).Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
