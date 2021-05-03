package api

type GetApexTrackerStatsInput struct {
	Platform string
	Id       string
}

type GetApexTrackerStatsOutput struct {
	Data string
}
