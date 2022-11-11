package main

import (
	"context"
	"fmt"
	"github.com/ivanov-gv/color-picker-test-app/internal/controller"
	"github.com/ivanov-gv/color-picker-test-app/internal/service"
	"github.com/ivanov-gv/color-picker-test-app/internal/service/dao"
	"github.com/ivanov-gv/color-picker-test-app/pkg/config"
	"github.com/ivanov-gv/color-picker-test-app/pkg/logger"
	"github.com/ivanov-gv/color-picker-test-app/pkg/postgres"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func main() {
	// Get logger interface.
	log := logger.New()

	if err := mainNoExit(log); err != nil {
		log.Fatalf("fatal err: %s", err.Error())
	}
}

func mainNoExit(log logrus.FieldLogger) error {
	confPath := os.Getenv("CONFIG_PATH")
	if confPath == "" {
		return fmt.Errorf("CONFIG_PATH not set")
	}
	cfg, err := config.Parse(confPath)
	if err != nil {
		return err
	}

	log.Println(cfg)
	log.Println("Starting the service...")

	// Postgresql
	pool, err := postgres.New(cfg.Pg)
	if err != nil {
		return fmt.Errorf("can't create pg pool: %s", err.Error())
	}

	// Use cases
	userDao := dao.NewUserPostgres(pool)
	userService := service.NewUserService(userDao)
	colorDao := dao.NewColorPostgres(pool)
	colorService := service.NewColorService(colorDao, userService)

	// HTTP router
	ctx := context.Background()

	router, err := controller.Router(ctx, log, colorService, userService)
	if err != nil {
		return fmt.Errorf("can't init router: %s", err.Error())
	}

	log.Print("The service is ready to listen and serve.")
	return http.ListenAndServe(
		cfg.Http.AppPort,
		router,
	)
}
