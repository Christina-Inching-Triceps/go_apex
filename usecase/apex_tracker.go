package usecase

// FIXME: 引数はInputPortのデータクラスを別途用意する
type ApexTrackerUseCase interface {
	GetStats(platform string, userId string) string
}
