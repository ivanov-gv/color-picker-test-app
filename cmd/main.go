package main

import (
	"context"
	"fmt"
	"github.com/ivanov-gv/color-picker-test-app/internal/controller"
	"github.com/ivanov-gv/color-picker-test-app/internal/service"
	"github.com/ivanov-gv/color-picker-test-app/internal/service/dao"
	"github.com/ivanov-gv/color-picker-test-app/pkg/config"
	"github.com/ivanov-gv/color-picker-test-app/pkg/httpserver"
	"github.com/ivanov-gv/color-picker-test-app/pkg/logger"
	"github.com/ivanov-gv/color-picker-test-app/pkg/postgres"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log := logger.New()

	if err := app(log); err != nil {
		log.Fatalf("fatal err: %s", err.Error())
	}
}

func app(log logrus.FieldLogger) error {
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
		return fmt.Errorf("can't create pg pool: %w", err)
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
		return fmt.Errorf("can't init router: %w", err)
	}

	httpServer := httpserver.New(router, cfg.Http)
	httpServer.Start()

	// Start app
	log.Print("The service is ready to listen and serve.")

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		log.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		return fmt.Errorf("app - Run - httpServer.Shutdown: %w", err)
	}
	return nil
}
