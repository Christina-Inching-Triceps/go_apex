package controller

import (
	"fmt"
	"gopex/domain/application"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

// apex.tracker.ggへの問い合わせを行うController
type ApexController struct {
	ApexTrackerInteractor application.ApexTrackerInteractor
}

// FIXME: ここでDB(Storage)ioを渡していくはず
func NewApexController() *ApexController {
	return &ApexController{}
}

func (controller *ApexController) GetStats(c echo.Context) (err error) {
	// TODO: contextから値を取得してInputPortを生成

	// URLの組み立て、この辺もApexAPIのクラスを作ったり一部をconfに移す
	url := fmt.Sprintf("https://public-api.tracker.gg/v2/apex/standard/profile/%s/%s", "psn", "raru_ex")

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	// FIXME: 外部API callはrepositoryに隠蔽した方がいいかも
	// repoに隠蔽できるとconfをDB(persistence)レイヤーに依存する形で丸く収まるかも？
	// req.Header.Add(config.API.APEX.Header, config.API.APEX.Token)
	controller.ApexTrackerInteractor.GetStats("psn", "userName")

	// TODO: GetStatsから取得したデータを元にOutputPortを作成してcontextにセット

	// FIXME: この辺の処理をinfra層に隠蔽する
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
	c.String(http.StatusOK, string(body))

	return err

}
