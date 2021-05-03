package api

type ApexTrackerClient interface {
	GetStats(platform string, id string) (string, error)
}
