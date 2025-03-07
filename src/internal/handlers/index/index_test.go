package index

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	mockRepoIndex "go-echo-v2/internal/repositories/index/mock_index"
	"go-echo-v2/internal/services/index"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestIndex(t *testing.T) {
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
