package api

import (
	"context"
	"fmt"
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

func (interactor *GetApexTrackerStatsInteractor) GetStatsWithContext(ctx context.Context, input usecase.GetApexTrackerStatsInput) (usecase.GetApexTrackerStatsOutput, error) {
	stats, err := interactor.client.GetStatsWithContext(
		ctx,
		input.Platform,
		input.ID,
	)
	if err != nil {
		err = fmt.Errorf("[Fatal] Failed to request to apex.tracker.gg stats endpoint: %w \n", err)

		return usecase.GetApexTrackerStatsOutput{}, err
	}

	output := usecase.GetApexTrackerStatsOutput{
		Data: stats,
	}

	return output, err
}
