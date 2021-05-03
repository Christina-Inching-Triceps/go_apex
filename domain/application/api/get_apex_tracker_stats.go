package api

import (
	"gopex/adapter/api"
	usecase "gopex/usecase/api"
)

type GetApexTrackerStatsInteractor struct {
	client api.ApexTrackerClient
}

func NewGetApexTrackerStatsInteractor(client *api.ApexTrackerClient) usecase.GetApexTrackerStatsUseCase {
	return &GetApexTrackerStatsInteractor{
		client: *client,
	}
}

func (interactor *GetApexTrackerStatsInteractor) GetStats(input usecase.GetApexTrackerStatsInput) (usecase.GetApexTrackerStatsOutput, error) {
	stats, err := interactor.client.GetStats(
		input.Platform,
		input.Id,
	)

	output := usecase.GetApexTrackerStatsOutput{
		Data: stats,
	}

	return output, err
}
