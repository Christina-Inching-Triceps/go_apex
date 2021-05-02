package controller

import (
	"fmt"
	"gopex/infrastructure/config"
	"gopex/usecase"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

// apex.tracker.ggへの問い合わせを行うController
type ApexController struct {
	ApexTrackerUseCase usecase.ApexTrackerUseCase
	config             *config.Config
}

func NewApexController(usecase *usecase.ApexTrackerUseCase, config *config.Config) *ApexController {
	return &ApexController{
		ApexTrackerUseCase: *usecase,
		config:             config,
	}
}

func (controller *ApexController) GetStats(c echo.Context) (err error) {
	// TODO: contextから値を取得してInputPortを生成

	// URLの組み立て、この辺もApexAPIのクラスを作ったり一部をconfに移す
	url := fmt.Sprintf("https://public-api.tracker.gg/v2/apex/standard/profile/%s/%s", "psn", "raru_ex")

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	// TODO: GetStatsから取得したデータを元にOutputPortを作成してcontextにセット

	// FIXME: この辺の処理をinfra層に隠蔽する
	req.Header.Add(controller.config.API.APEX.Header, controller.config.API.APEX.Token)

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	created, err := controller.ApexTrackerUseCase.StoreStats(string(body))
	stats, err := controller.ApexTrackerUseCase.GetStats(created.ID)

	c.String(http.StatusOK, stats)

	return err

}
