package healthcheck

import (
	repoHealthcheck "go-echo-v2/internal/repositories/healthcheck"
	"go-echo-v2/internal/services/healthcheck"

	"github.com/labstack/echo/v4"
)

// OpenAPI仕様書用の型定義
type OKResponse struct {
	Message string `json:"message" example:"Health Check OK !!"`
}

type InternalServerErrorResponse struct {
	Message string `json:"message" example:"Failed to health check: error message"`
}

// @Description APIとDBの接続確認をするためのヘルスチェックAPI
// @Tags healthcheck
// @Security Bearer
// @Success 200 {object} OKResponse
// @Failure 500 {object} InternalServerErrorResponse
// @Router /api/v1/healthcheck [get]
func Healthcheck(c echo.Context) error {
	// インスタンス生成
	healthcheckRepository := repoHealthcheck.NewHealthcheckRepository()
	healthcheckService := healthcheck.NewHealthcheckService(healthcheckRepository)

	// サービス実行
	return healthcheckService.Healthcheck(c)
}
