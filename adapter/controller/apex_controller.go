package controller

import (
	"context"
	"fmt"
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

func NewApexController(
	usecase *usecase.UserReadApexProfileSummaryUseCase,
	api *apiUsecase.GetApexTrackerStatsUseCase,
	config *config.Config) *ApexController {
	return &ApexController{
		UserReadApexProfileSummaryUseCase: *usecase,
		GetApexTrackerStatsUseCase:        *api,
		config:                            config,
	}
}

func (controller *ApexController) GetStats(c echo.Context) (err error) {
	input := apiUsecase.GetApexTrackerStatsInput{
		Platform: "psn",
		ID:       "raru_ex",
	}

	body, err := controller.GetApexTrackerStatsUseCase.GetStatsWithContext(context.Background(), input)
	if err != nil {
		return fmt.Errorf("[Fatal] Request to apex.tracker.gg is failed : %w \n", err)
	}

	created, err := controller.UserReadApexProfileSummaryUseCase.StoreStats(body.Data)
	if err != nil {
		return fmt.Errorf("[Fatal] Failed to store tracker stats to db: %w \n", err)
	}

	stats, err := controller.UserReadApexProfileSummaryUseCase.GetStats(created.ID)
	if err != nil {
		return fmt.Errorf("[Fatal] Failed to select stats from db: %w \n", err)
	}

	err = c.String(http.StatusOK, stats)
	if err != nil {
		err = fmt.Errorf("[Fatal] Failed to create response by echo: %w \n", err)

		return err
	}

	return err
}
