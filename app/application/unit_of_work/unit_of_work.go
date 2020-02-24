package unit_of_work

type IUnitOfWork interface {
	Rollback()
	Commit() // 呼び出すことでこれまでの変更処理を全てデータストアに適用する
}
