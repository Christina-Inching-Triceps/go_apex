package persistence

import (
	"gopex/adapter/repository"
	"gopex/domain/entity"

	"gorm.io/gorm"
)

// TODO: rename 実際にはもう少し細かく分かれるはず
type ApexTrackerRepository struct {
	db *gorm.DB
}

// FIXME: DBをコンストラクタで注入
func NewApexTrackerRepository(db *gorm.DB) repository.ApexTrackerRepository {
	return &ApexTrackerRepository{
		db: db,
	}
}

func (repo *ApexTrackerRepository) Create(row *entity.ApexTracker) (err error) {
	return repo.db.Create(&row).Error
}

func (repo *ApexTrackerRepository) Find(id uint) (row *entity.ApexTracker, err error) {
	result := repo.db.Model(&entity.ApexTracker{}).First(&row, id)
	return row, result.Error
}
