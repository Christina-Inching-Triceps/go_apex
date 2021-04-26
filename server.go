package main

import (
	"gopex/infrastructure"
	"gopex/infrastructure/config"

	"github.com/labstack/echo/v4"
)

func main() {
	// ===== Init Echo Server
	e := echo.New()

	// ===== Load Config File
	config, err := config.Load()
	if err != nil {
		e.Logger.Fatal(err)
	}

	// ===== Init DB Instance
	// See: https://please-sleep.cou929.nu/go-sql-db-connection-pool.html
	db, err := infrastructure.NewDB(config)
	if err != nil {
		e.Logger.Fatal(err)
	}
	// FIXME: 暫定
	_ = db

	// ===== Setup Router
	infrastructure.SetUpRouting(e, config)

	e.Logger.Info(e.Start(":1323"))
}
