package index

import (
	"os"
)

// インターフェース定義
type IndexRepository interface {
	Hello() string
}

// 構造体定義
type indexRepository struct{}

// インスタンス生成用関数
func NewIndexRepository() IndexRepository {
	return &indexRepository{}
}

// メソッド定義
func (r *indexRepository) Hello() string {
	env := os.Getenv("ENV")

	res := "Hello World !!"
	if env == "testing" {
		res = "Testing Hello World !!"
	}

	return res
}
