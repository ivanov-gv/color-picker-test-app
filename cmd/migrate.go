//go:build migrate

package main

import (
	"errors"
	"github.com/ivanov-gv/color-picker-test-app/pkg/config"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func init() {
	confPath := os.Getenv("CONFIG_PATH")
	if confPath == "" {
		log.Fatalf("CONFIG_PATH not set")
		return
	}
	cfg, err := config.Parse(confPath)
	if err != nil {
		return
	}

	log.Println(cfg)

	var m *migrate.Migrate
	m, err = migrate.New("file://migrations", cfg.Pg.DbConnString+"?sslmode=disable")
	if err != nil {
		log.Fatalf("Migrate: postgres connect error: %s", err)
	}

	err = m.Up()
	defer m.Close()

	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("Migrate: up error: %s", err)
	}

	if errors.Is(err, migrate.ErrNoChange) {
		log.Printf("Migrate: no change")
		return
	}

	log.Printf("Migrate: up success")
}
