package service

import (
	"context"
	"fmt"
	"github.com/ivanov-gv/color-picker-test-app/internal/model"
)

type ColorService struct {
	colorDao    ColorDao
	userService UserInterface
}

var _ ColorInterface = &ColorService{}

func NewColorService(colorDao ColorDao, userService UserInterface) ColorInterface {
	return &ColorService{
		colorDao:    colorDao,
		userService: userService}
}

func (s *ColorService) GetAll(ctx context.Context, userId int) ([]model.Color, error) {
	return s.colorDao.GetAll(ctx, userId)
}

func (s *ColorService) Get(ctx context.Context, userId int, colorId int) (model.Color, error) {
	return s.colorDao.Get(ctx, userId, colorId)
}

func (s *ColorService) Add(ctx context.Context, userId int, color model.Color) (model.Color, error) {
	ok, err := s.userService.Exist(ctx, userId)

	if err != nil {
		return model.Color{}, fmt.Errorf("ColorService - Add - error while check user existence: %w", err)
	}

	if !ok {
		return model.Color{}, fmt.Errorf("ColorService - Add - user does not exist: id=%d", userId)
	}

	return s.colorDao.Add(ctx, userId, color)
}

func (s *ColorService) Delete(ctx context.Context, userId int, colorId int) (err error) {
	return s.colorDao.Delete(ctx, userId, colorId)
}
