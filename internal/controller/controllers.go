package controller

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/ivanov-gv/color-picker-test-app/internal/controller/user"

	"github.com/ivanov-gv/color-picker-test-app/internal/controller/color"
	"github.com/ivanov-gv/color-picker-test-app/internal/service"

	"github.com/sirupsen/logrus"
)

const (
	userRoute  = "/user"
	colorRoute = "/{user_id:[0-9]+}/color"
)

// Router register necessary routes and returns an instance of a router.
func Router(ctx context.Context, log logrus.FieldLogger,
	colorService service.ColorInterface,
	userService service.UserInterface) (*mux.Router, error) {
	r := mux.NewRouter()
	r.Use(ContentTypeJson)

	v := validator.New()

	// k8s
	NewK8s(r)

	// user
	userRouter := r.PathPrefix(userRoute).Subrouter()
	user.NewUserController(userRouter, userService, v, ctx, log)

	// user/color
	colorRouter := userRouter.PathPrefix(colorRoute).Subrouter()
	color.NewColorController(colorRouter, colorService, v, ctx, log)

	return r, nil
}
