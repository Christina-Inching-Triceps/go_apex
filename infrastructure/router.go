package infrastructure

import (
	"gopex/adapter/controller"
	"gopex/infrastructure/config"

	"github.com/labstack/echo/v4"
)

// EchoのRouting情報をセットアップします
func SetUpRouting(e *echo.Echo, config *config.Config) {

	apexController := controller.NewApexController()
	e.GET("api/get", apexController.GetStats)

}
