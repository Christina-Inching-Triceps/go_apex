package application

import (
	"gopex/adapter/repository"
	"gopex/domain/entity"
	"gopex/usecase"
)

type ApexTrackerInteractor struct {
	apexTrackerRepository repository.ApexTrackerRepository
}

func NewApexTrackerInteractor(apexTrackerRepository *repository.ApexTrackerRepository) usecase.ApexTrackerUseCase {
	return &ApexTrackerInteractor{
		apexTrackerRepository: *apexTrackerRepository,
	}
}

func (interactor *ApexTrackerInteractor) GetStats(id uint) (string, error) {
	row, err := interactor.apexTrackerRepository.Find(id)
	if err != nil {
		return "", err
	}
	return row.Content, err
}

func (interactor *ApexTrackerInteractor) StoreStats(content string) (entity.ApexTracker, error) {
	created := entity.ApexTracker{Content: content}
	result := interactor.apexTrackerRepository.Create(&created)
	return created, result
}
