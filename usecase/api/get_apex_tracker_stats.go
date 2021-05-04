package api

import "context"

type GetApexTrackerStatsUseCase interface {
	GetStatsWithContext(ctx context.Context, input GetApexTrackerStatsInput) (GetApexTrackerStatsOutput, error)
}
