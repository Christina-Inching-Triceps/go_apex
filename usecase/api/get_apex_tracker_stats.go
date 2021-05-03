package api

type GetApexTrackerStatsUseCase interface {
	GetStats(input GetApexTrackerStatsInput) (GetApexTrackerStatsOutput, error)
}
