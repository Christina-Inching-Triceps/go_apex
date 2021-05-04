package api

import "context"

type ApexTrackerClient interface {
	GetStatsWithContext(ctx context.Context, platform string, id string) (string, error)
}
