package usecase

import "gopex/domain/entity"

// FIXME: 引数はInputPortのデータクラスを別途用意する
// TODO: Repository実装を仮に作成。API実装を今後追加。その際にrename and フォルダ構造変更
type ApexTrackerUseCase interface {
	// 本当はuid
	GetStats(id uint) (string, error)
	StoreStats(content string) (entity.ApexTracker, error)
}
