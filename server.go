package main

import (
	"gopex/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	// ===== Init Echo Server
	e := echo.New()

	// ===== Setup Router
	routes.Router(e)

	e.Logger.Info(e.Start(":1323"))
}
