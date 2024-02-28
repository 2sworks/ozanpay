package service

import (
	"context"
	"gorm.io/gorm"
	"ozanpay/model"
)

type IUserService interface {
	Create(ctx context.Context, user *model.User) error
}

type UserService struct {
	DB *gorm.DB
}

func NewUserService(db *gorm.DB) IUserService {
	return &UserService{
		DB: db,
	}
}

func (s *UserService) Create(ctx context.Context, user *model.User) error {
	return s.DB.WithContext(ctx).Create(user).Error
}
