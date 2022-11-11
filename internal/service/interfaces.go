package service

import (
	"context"
	"github.com/ivanov-gv/color-picker-test-app/internal/model"
)

type ColorInterface interface {
	GetAll(ctx context.Context, userId int) ([]model.Color, error)
	Get(ctx context.Context, userId int, colorId int) (model.Color, error)
	Add(ctx context.Context, userId int, color model.Color) (model.Color, error)
	Delete(ctx context.Context, userId int, colorId int) (err error)
}
type ColorDao interface {
	GetAll(ctx context.Context, userId int) ([]model.Color, error)
	Get(ctx context.Context, userId int, colorId int) (model.Color, error)
	Add(ctx context.Context, userId int, color model.Color) (model.Color, error)
	Delete(ctx context.Context, userId int, colorId int) error
}

type UserInterface interface {
	Create(ctx context.Context) (model.User, error)
	Exist(ctx context.Context, userId int) (bool, error)
}
type UserDao interface {
	Create(ctx context.Context) (model.User, error)
	Exist(ctx context.Context, userId int) (bool, error)
}
