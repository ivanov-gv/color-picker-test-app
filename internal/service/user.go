package service

import (
	"context"
	"github.com/ivanov-gv/color-picker-test-app/internal/model"
)

type UserService struct {
	userDao UserDao
}

var _ UserInterface = &UserService{}

func NewUserService(dao UserDao) UserInterface {
	return &UserService{userDao: dao}
}

func (s *UserService) Create(ctx context.Context) (model.User, error) {
	return s.userDao.Create(ctx)
}

func (s *UserService) Exist(ctx context.Context, userId int) (bool, error) {
	return s.userDao.Exist(ctx, userId)
}
