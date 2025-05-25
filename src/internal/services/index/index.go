package index

import (
	"fmt"

	"go-echo-v2/internal/repositories/index"
)

// インターフェース定義
type IndexService interface {
	Index() (string, error)
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

func (s *indexService) Index() (string, error) {
	text := s.indexRepository.Hello()
	if text == "" {
		return "", fmt.Errorf("textが空です。")
	}

	return text, nil
}
