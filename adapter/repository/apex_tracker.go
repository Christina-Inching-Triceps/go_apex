package repository

import "gopex/domain/entity"

// TODO: rename 実際にはもう少し細かく分かれるはず
type ApexTrackerRepository interface {
	Find(id uint) (row *entity.ApexTracker, err error)
	Create(row *entity.ApexTracker) (err error)
}
