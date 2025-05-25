package index

import (
	repoIndex "go-echo-v2/internal/repositories/index"
	serviceIndex "go-echo-v2/internal/services/index"
	usecaseIndex "go-echo-v2/internal/usecases/index"

	"github.com/labstack/echo/v4"
)

// @Description テキスト「Hello World !!」を出力する。
// @Tags index
// @Success 200
// @Router /api/v1/ [get]
func Index(c echo.Context) error {
	// インスタンス生成
	indexRepository := repoIndex.NewIndexRepository()
	indexService := serviceIndex.NewIndexService(indexRepository)
	indexUsecase := usecaseIndex.NewIndexUsecase(indexService)

	// ユースケースの実行
	return indexUsecase.Exec(c)
}
