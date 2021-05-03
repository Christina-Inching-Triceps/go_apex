package application

import (
	"gopex/adapter/repository"
	"gopex/domain/entity"
	"gopex/usecase"
)

type UserReadApexProfileSummaryInteractor struct {
	apexTrackerRepository repository.ApexTrackerRepository
}

func NewUserReadApexProfileSummaryInteractor(apexTrackerRepository *repository.ApexTrackerRepository) usecase.UserReadApexProfileSummaryUseCase {
	return &UserReadApexProfileSummaryInteractor{
		apexTrackerRepository: *apexTrackerRepository,
	}
}

func (interactor *UserReadApexProfileSummaryInteractor) GetStats(id uint) (string, error) {
	row, err := interactor.apexTrackerRepository.Find(id)
	if err != nil {
		return "", err
	}
	return row.Content, err
}

func (interactor *UserReadApexProfileSummaryInteractor) StoreStats(content string) (entity.ApexTracker, error) {
	created := entity.ApexTracker{Content: content}
	result := interactor.apexTrackerRepository.Create(&created)
	return created, result
}
