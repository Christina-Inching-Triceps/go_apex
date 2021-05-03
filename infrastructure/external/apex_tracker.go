package external

import (
	"fmt"
	"gopex/adapter/api"
	"gopex/infrastructure/config"
	"io/ioutil"
	"net/http"
)

type ApexTrackerClient struct {
	config *config.Config
}

func NewApexTrackerApi(config *config.Config) api.ApexTrackerClient {
	return &ApexTrackerClient{
		config: config,
	}
}

func (api *ApexTrackerClient) GetStats(platform string, id string) (string, error) {
	url := fmt.Sprintf("https://public-api.tracker.gg/v2/apex/standard/profile/%s/%s", "psn", "raru_ex")

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	// FIXME: この辺の処理をinfra層に隠蔽する
	req.Header.Add(api.config.API.APEX.Header, api.config.API.APEX.Token)

	client := new(http.Client)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	return string(body), err
}
