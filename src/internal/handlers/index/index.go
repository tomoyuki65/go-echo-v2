package index

import (
	repoIndex "go-echo-v2/internal/repositories/index"
	"go-echo-v2/internal/services/index"

	"github.com/labstack/echo/v4"
)

// @Description テキスト「Hello World !!」を出力する。
// @Tags index
// @Success 200
// @Router /api/v1/ [get]
func Index(c echo.Context) error {
	// インスタンス生成
	indexRepository := repoIndex.NewIndexRepository()
	indexService := index.NewIndexService(indexRepository)

	// サービス実行
	return indexService.Index(c)
}
