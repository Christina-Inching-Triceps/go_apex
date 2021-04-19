package application

import "gopex/adapter/repository"

type ApexTrackerInteractor struct {
	ApexTrackerRepository repository.ApexTrackerRepository
}

func (interactor *ApexTrackerInteractor) GetStats(platform string, userId string) string {
	return "get stats"
}
