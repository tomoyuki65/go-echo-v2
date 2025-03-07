package index

import (
	"net/http"

	"go-echo-v2/internal/repositories/index"

	"github.com/labstack/echo/v4"
)

// インターフェース定義
type IndexService interface {
	Index(c echo.Context) error
}

// 構造体定義
type indexService struct {
	indexRepository index.IndexRepository
}

// インスタンス生成用関数
func NewIndexService(
	indexRepository index.IndexRepository,
) IndexService {
	return &indexService{
		indexRepository: indexRepository,
	}
}

func (s *indexService) Index(c echo.Context) error {
	text := s.indexRepository.Hello()
	return c.String(http.StatusOK, text)
}
