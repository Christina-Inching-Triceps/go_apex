package controller

import (
	"gopex/infrastructure/config"
	"gopex/usecase"
	apiUsecase "gopex/usecase/api"
	"net/http"

	"github.com/labstack/echo/v4"
)

// apex.tracker.ggへの問い合わせを行うController
type ApexController struct {
	UserReadApexProfileSummaryUseCase usecase.UserReadApexProfileSummaryUseCase
	GetApexTrackerStatsUseCase        apiUsecase.GetApexTrackerStatsUseCase
	config                            *config.Config
}

func NewApexController(usecase *usecase.UserReadApexProfileSummaryUseCase, api *apiUsecase.GetApexTrackerStatsUseCase, config *config.Config) *ApexController {
	return &ApexController{
		UserReadApexProfileSummaryUseCase: *usecase,
		GetApexTrackerStatsUseCase:        *api,
		config:                            config,
	}
}

func (controller *ApexController) GetStats(c echo.Context) (err error) {
	input := apiUsecase.GetApexTrackerStatsInput{
		Platform: "psn",
		Id:       "raru_ex",
	}

	body, err := controller.GetApexTrackerStatsUseCase.GetStats(input)
	if err != nil {
		return err
	}

	created, err := controller.UserReadApexProfileSummaryUseCase.StoreStats(body.Data)
	stats, err := controller.UserReadApexProfileSummaryUseCase.GetStats(created.ID)

	c.String(http.StatusOK, stats)

	return err

}
