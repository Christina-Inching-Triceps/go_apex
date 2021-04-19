package persistence

type ApexTrackerRepository struct {
}

// FIXME: DBをコンストラクタで注入
func NewApexTrackerRepository() *ApexTrackerRepository {
	return &ApexTrackerRepository{}
}
