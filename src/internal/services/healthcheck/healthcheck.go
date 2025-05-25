package healthcheck

import (
	"context"

	"go-echo-v2/internal/repositories/healthcheck"
)

// インターフェースの定義
type HealthcheckService interface {
	Healthcheck(ctx context.Context) error
}

// 構造体の定義
type healthcheckService struct {
	healthcheckRepository healthcheck.HealthcheckRepository
}

// インスタンス生成用関数の定義
func NewHealthcheckService(
	healthcheckRepository healthcheck.HealthcheckRepository,
) HealthcheckService {
	return &healthcheckService{
		healthcheckRepository: healthcheckRepository,
	}
}

// Healthcheckメソッドの実装
func (s *healthcheckService) Healthcheck(ctx context.Context) error {
	err := s.healthcheckRepository.Healthcheck(ctx)
	if err != nil {
		return err
	}

	return nil
}
