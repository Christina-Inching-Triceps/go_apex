package usecase

import "gopex/domain/entity"

// TODO: 命名が長すぎる。処理が明確になり次第フォルダを切った方がよい
type UserReadApexProfileSummaryUseCase interface {
	// 本当はuid
	// FIXME: 引数はInputPortのデータクラスを別途用意する
	GetStats(id uint) (string, error)
	StoreStats(content string) (entity.ApexTracker, error)
}
