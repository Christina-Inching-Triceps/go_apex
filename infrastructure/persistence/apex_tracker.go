package persistence

import "gopex/adapter/repository"

type ApexTrackerRepository struct {
}

// FIXME: DBをコンストラクタで注入
func NewApexTrackerRepository() repository.ApexTrackerRepository {
	return &ApexTrackerRepository{}
}
