package routes

import (
	"fmt"
	"gopex/config"
	"io/ioutil"
	"net/http"
	"github.com/labstack/echo/v4"
)

func Router(e* echo.Echo) {

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("api/get", func(c echo.Context) error {
		// 一旦ここでconfigをload。もう少し広い範囲で取得したい or Singletonにする？
		config, err := config.Load()
		if err != nil {
			e.Logger.Fatal(err)
		}

		// URLの組み立て、この辺もApexAPIのクラスを作ったり一部をconfに移す
		url := fmt.Sprintf("https://public-api.tracker.gg/v2/apex/standard/profile/%s/%s", "psn", "raru_ex")

		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			e.Logger.Fatal(err)
		}
		req.Header.Add(config.API.APEX.Header, config.API.APEX.Token)

		client := new(http.Client)
		resp, err := client.Do(req)
		if err != nil {
			e.Logger.Fatal(err)
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			e.Logger.Fatal(err)
		}

		return c.String(http.StatusOK, string(body))
	})

}
