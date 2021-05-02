package infrastructure

import (
	"gopex/adapter/controller"
	"gopex/infrastructure/config"
	"gopex/usecase"

	"github.com/labstack/echo/v4"
)

// EchoのRouting情報をセットアップします
func InitApexTrackerRouting(e *echo.Echo, config *config.Config, interactor *usecase.ApexTrackerUseCase) {

	apexController := controller.NewApexController(interactor, config)
	e.GET("api/get", apexController.GetStats)

}
