package index

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	mockRepoIndex "go-echo-v2/internal/repositories/index/mock_index"
	"go-echo-v2/internal/services/index"
	"go-echo-v2/middleware"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestIndex(t *testing.T) {
	// .env ファイルの読み込み
	if err := godotenv.Load("../../../.env"); err != nil {
		slog.Error(".envファイルの読み込みに失敗しました。")
	}

	// テスト用のENV設定
	env := os.Getenv("ENV")
	os.Setenv("ENV", "testing")
	defer os.Setenv("ENV", env)

	e := echo.New()

	// テスト用リクエストのecho.Context作成
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// リポジトリのモック化
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockIndexRepository := mockRepoIndex.NewMockIndexRepository(ctrl)
	mockIndexRepository.EXPECT().Hello().Return("Hello World !!")
	indexService := index.NewIndexService(mockIndexRepository)

	// テスト実行
	err := indexService.Index(c)

	// 検証
	assert.Equal(t, nil, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "Hello World !!", rec.Body.String())
}

// 認証用ミドルウェアのテスト
func TestAuthMiddleware(t *testing.T) {
	// .env ファイルの読み込み
	if err := godotenv.Load("../../../.env"); err != nil {
		slog.Error(".envファイルの読み込みに失敗しました。")
	}

	// テスト用のENV設定
	env := os.Getenv("ENV")
	os.Setenv("ENV", "testing")
	defer os.Setenv("ENV", env)

	e := echo.New()

	// ミドルウェアの適用
	v1 := e.Group("/api/v1")
	v1.GET("/", Index, middleware.AuthMiddleware)

	// テスト用リクエストのecho.Context作成
	req := httptest.NewRequest(http.MethodGet, "/api/v1/", nil)
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
