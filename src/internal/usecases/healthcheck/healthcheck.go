package healthcheck

import (
	"fmt"
	"net/http"

	"go-echo-v2/internal/services/healthcheck"
	utilContext "go-echo-v2/util/context"
	"go-echo-v2/util/logger"

	"github.com/labstack/echo/v4"
)

// レスポンス結果の型定義
type OKResponse struct {
	Message string `json:"message"`
}

// インターフェースの定義
type HealthcheckUsecase interface {
	Exec(c echo.Context) error
}

// 構造体の定義
type healthcheckUsecase struct {
	healthcheckService healthcheck.HealthcheckService
}

// インスタンス生成用関数の定義
func NewHealthcheckUsecase(
	healthcheckService healthcheck.HealthcheckService,
) HealthcheckUsecase {
	return &healthcheckUsecase{
		healthcheckService: healthcheckService,
	}
}

// Healthcheckメソッドの実装
func (s *healthcheckUsecase) Exec(c echo.Context) error {
	ctx := utilContext.CreateContext(c)

	err := s.healthcheckService.Healthcheck(ctx)
	if err != nil {
		msg := fmt.Sprintf("Failed to health check: %v", err)
		logger.Error(ctx, msg)
		return echo.NewHTTPError(http.StatusInternalServerError, msg)
	}

	res := OKResponse{
		Message: "Health Check OK !!",
	}

	return c.JSON(http.StatusOK, res)
}
