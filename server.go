package main

import (
	"fmt"
	"gopex/config"
	"gopex/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	// ===== Init Echo Server
	e := echo.New()

	// ===== Load Config
	 config, err := config.Load()
	 if err != nil {
		 e.Logger.Fatal(err)
	 }

	 // TODO
	 fmt.Printf("[Debug] token: %s \n", config.Apex.Token)
	 _ = config

	// ===== Setup Router
	routes.Router(e)

	e.Logger.Info(e.Start(":1323"))
}
