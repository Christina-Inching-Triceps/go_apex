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

	// ===== Setup Router
	infrastructure.SetUpRouting(e, config)

	e.Logger.Info(e.Start(":1323"))
}
