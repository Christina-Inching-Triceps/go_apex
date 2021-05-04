package external

import (
	"context"
	"fmt"
	"gopex/adapter/api"
	"gopex/infrastructure/config"
	"io/ioutil"
	"net/http"
)

type ApexTrackerClient struct {
	config *config.Config
}

func NewApexTrackerClient(config *config.Config) api.ApexTrackerClient {
	return &ApexTrackerClient{
		config: config,
	}
}

func (api *ApexTrackerClient) GetStatsWithContext(ctx context.Context, platform string, id string) (string, error) {
	url := fmt.Sprintf("https://public-api.tracker.gg/v2/apex/standard/profile/%s/%s", "psn", "raru_ex")

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", fmt.Errorf("[Fatal] Failed to create http request instance: %w \n", err)
	}

	req.Header.Add(api.config.API.APEX.Header, api.config.API.APEX.Token)

	client := new(http.Client)

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("[Fatal] Failed to request execution: %w \n", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("[Fatal] Failed to read response body: %w \n", err)

		return "", err
	}

	return string(body), err
}
