package repository

import "gopex/domain/entity"

type ApexTrackerRepository interface {
	Find(id uint) (row *entity.ApexTracker, err error)
	Create(row *entity.ApexTracker) (err error)
}
