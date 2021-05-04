package application

import (
	"fmt"
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
		err = fmt.Errorf("[Fatal] Failed to find tracker data from table (id=%d): %w \n", id, err)

		return "", err
	}

	return row.Content, err
}

func (interactor *UserReadApexProfileSummaryInteractor) StoreStats(content string) (entity.ApexTracker, error) {
	created := entity.ApexTracker{Content: content}

	err := interactor.apexTrackerRepository.Create(&created)
	if err != nil {
		err = fmt.Errorf("[Fatal] Failed to create tracker data to table: %w \n", err)
	}

	return created, err
}
