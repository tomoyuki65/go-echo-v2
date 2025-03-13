package healthcheck

import (
	"context"
	"fmt"

	"go-echo-v2/database"
)

// インターフェースの定義
type HealthcheckRepository interface {
	Healthcheck(ctx context.Context) error
}

// 構造体の定義
type healthcheckRepository struct{}

// インスタンス生成用関数の定義
func NewHealthcheckRepository() HealthcheckRepository {
	return &healthcheckRepository{}
}

// Healthcheckメソッドの実装
func (h *healthcheckRepository) Healthcheck(ctx context.Context) error {
	db, err := database.SetupDatabaseWithGorm(ctx)
	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	defer sqlDB.Close()

	// DB接続確認
	if err := sqlDB.Ping(); err != nil {
		return fmt.Errorf("database check failed: %w", err)
	}

	return nil
}
