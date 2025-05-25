package index

import (
	"fmt"
	"net/http"

	"go-echo-v2/internal/services/index"

	"github.com/labstack/echo/v4"
)

// インターフェース定義
type IndexUsecase interface {
	Exec(c echo.Context) error
}

// 構造体定義
type indexUsecase struct {
	indexService index.IndexService
}

// インスタンス生成用関数
func NewIndexUsecase(
	indexService index.IndexService,
) IndexUsecase {
	return &indexUsecase{
		indexService: indexService,
	}
}

func (s *indexUsecase) Exec(c echo.Context) error {
	text, err := s.indexService.Index()
	if err != nil {
		msg := fmt.Sprintf("Internal Server Error: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, msg)
	}

	return c.String(http.StatusOK, text)
}
