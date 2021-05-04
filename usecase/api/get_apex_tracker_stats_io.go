package api

type GetApexTrackerStatsInput struct {
	Platform string
	ID       string
}

type GetApexTrackerStatsOutput struct {
	Data string
}
