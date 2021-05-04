package infrastructure

import (
	"gopex/adapter/controller"
	"gopex/infrastructure/config"

	"github.com/labstack/echo/v4"
)

// EchoのRouting情報をセットアップします
func InitApexTrackerRouting(e *echo.Echo, config *config.Config, controller *controller.ApexController) {
	e.GET("api/get", controller.GetStats)
}
