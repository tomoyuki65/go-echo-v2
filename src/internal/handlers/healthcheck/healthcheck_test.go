package healthcheck

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	mockRepoHealthcheck "go-echo-v2/internal/repositories/healthcheck/mock_healthcheck"
	"go-echo-v2/internal/services/healthcheck"
	"go-echo-v2/middleware"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestHealthCheckOK(t *testing.T) {
	// .env ファイルの読み込み
	if err := godotenv.Load("../../../.env"); err != nil {
		slog.Error(".envファイルの読み込みに失敗しました。")
	}

	// テスト用のENV設定
	env := os.Getenv("ENV")
	os.Setenv("ENV", "testing")
	defer os.Setenv("ENV", env)

	// ミドルウェアの適用
	e := echo.New()
	v1 := e.Group("/api/v1")
	v1.GET("/healthcheck", Healthcheck, middleware.ApiKeyAuthMiddleware)

	// テスト用リクエストの作成
	req := httptest.NewRequest(http.MethodGet, "/api/v1/healthcheck", nil)
	token := "zMdtq_glzI7oqq8yXjMgEOW6XfrSUMFGqw"
	bearerToken := "Bearer " + token
	req.Header.Set("Authorization", bearerToken)
	rec := httptest.NewRecorder()

	// テスト実行
	e.ServeHTTP(rec, req)

	// レスポンス結果のJSONを取得
	var resbody map[string]interface{}
	if err := json.Unmarshal(rec.Body.Bytes(), &resbody); err != nil {
		t.Fatal(err)
	}

	expected := map[string]interface{}{
		"message": "Health Check OK !!",
	}

	// 検証
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, expected, resbody)
}

func TestHealthCheckNotToken(t *testing.T) {
	// .env ファイルの読み込み
	if err := godotenv.Load("../../../.env"); err != nil {
		slog.Error(".envファイルの読み込みに失敗しました。")
	}

	// テスト用のENV設定
	env := os.Getenv("ENV")
	os.Setenv("ENV", "testing")
	defer os.Setenv("ENV", env)

	// ミドルウェアの適用
	e := echo.New()
	v1 := e.Group("/api/v1")
	v1.GET("/healthcheck", Healthcheck, middleware.ApiKeyAuthMiddleware)

	// テスト用リクエストの作成
	req := httptest.NewRequest(http.MethodGet, "/api/v1/healthcheck", nil)
	rec := httptest.NewRecorder()

	// テスト実行
	e.ServeHTTP(rec, req)

	// レスポンス結果のJSONを取得
	var resbody map[string]interface{}
	if err := json.Unmarshal(rec.Body.Bytes(), &resbody); err != nil {
		t.Fatal(err)
	}

	expected := map[string]interface{}{
		"message": "Unauthorized",
	}

	// 検証
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	assert.Equal(t, expected, resbody)
}

func TestHealthCheckAuthError(t *testing.T) {
	// .env ファイルの読み込み
	if err := godotenv.Load("../../../.env"); err != nil {
		slog.Error(".envファイルの読み込みに失敗しました。")
	}

	// テスト用のENV設定
	env := os.Getenv("ENV")
	os.Setenv("ENV", "testing")
	defer os.Setenv("ENV", env)

	// ミドルウェアの適用
	e := echo.New()
	v1 := e.Group("/api/v1")
	v1.GET("/healthcheck", Healthcheck, middleware.ApiKeyAuthMiddleware)

	// テスト用リクエストの作成
	req := httptest.NewRequest(http.MethodGet, "/api/v1/healthcheck", nil)
	token := "xxxxxxxxxx"
	bearerToken := "Bearer " + token
	req.Header.Set("Authorization", bearerToken)
	rec := httptest.NewRecorder()

	// テスト実行
	e.ServeHTTP(rec, req)

	// レスポンス結果のJSONを取得
	var resbody map[string]interface{}
	if err := json.Unmarshal(rec.Body.Bytes(), &resbody); err != nil {
		t.Fatal(err)
	}

	expected := map[string]interface{}{
		"message": "Unauthorized",
	}

	// 検証
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	assert.Equal(t, expected, resbody)
}

func TestHealthCheckDBError(t *testing.T) {
	// .env ファイルの読み込み
	if err := godotenv.Load("../../../.env"); err != nil {
		slog.Error(".envファイルの読み込みに失敗しました。")
	}

	// テスト用のENV設定
	env := os.Getenv("ENV")
	os.Setenv("ENV", "testing")
	defer os.Setenv("ENV", env)

	e := echo.New()

	// テスト用リクエストの作成
	req := httptest.NewRequest(http.MethodGet, "/api/v1/healthcheck", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ctx := c.Request().Context()

	// リポジトリのモック化
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockHealthcheckRepository := mockRepoHealthcheck.NewMockHealthcheckRepository(ctrl)
	mockHealthcheckRepository.EXPECT().Healthcheck(ctx).Return(echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("database check failed: err")))
	healthcheckService := healthcheck.NewHealthcheckService(mockHealthcheckRepository)

	// テスト実行
	err := healthcheckService.Healthcheck(c)

	// 検証
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "database check failed: err")
}
