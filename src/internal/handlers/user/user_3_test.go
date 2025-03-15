package user

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"go-echo-v2/middleware"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

/*
 * 全てのユーザー取得APIのテスト
 */

// 正常系
func TestGetAllUsersOK(t *testing.T) {
	// .env ファイルの読み込み
	if err := godotenv.Load("../../../.env"); err != nil {
		slog.Error(".envファイルの読み込みに失敗しました。")
	}

	// テスト用のENV設定
	env := os.Getenv("ENV")
	os.Setenv("ENV", "testing")
	defer os.Setenv("ENV", env)

	// ルーティング設定
	e := initTestEcho()
	v1 := e.Group("/api/v1")
	v1.GET("/users", GetAllUsers, middleware.AuthMiddleware)

	// テスト用データ登録
	userSeeder(t)

	// テスト用リクエストの作成
	req := httptest.NewRequest(http.MethodGet, "/api/v1/users", nil)
	token := "zMdtq_glzI7oqq8yXjMgEOW6XfrSUMFGqw"
	bearerToken := "Bearer " + token
	req.Header.Set("Authorization", bearerToken)
	rec := httptest.NewRecorder()

	// テスト実行
	e.ServeHTTP(rec, req)

	// レスポンス結果をJSON形式で取得
	var resbody []map[string]interface{}
	if err := json.Unmarshal(rec.Body.Bytes(), &resbody); err != nil {
		t.Fatal(err)
	}

	// 検証
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, 2, len(resbody))
	assert.Equal(t, "test-1", resbody[0]["uid"])
	assert.Equal(t, "姓１", resbody[0]["last_name"])
	assert.Equal(t, "名１", resbody[0]["first_name"])
	assert.Equal(t, "test-user1@test.com", resbody[0]["email"])
	assert.NotEmpty(t, resbody[0]["created_at"])
	assert.NotEmpty(t, resbody[0]["updated_at"])
	assert.Empty(t, resbody[0]["deleted_at"])
	assert.Equal(t, "test-2", resbody[1]["uid"])
	assert.Equal(t, "姓２", resbody[1]["last_name"])
	assert.Equal(t, "名２", resbody[1]["first_name"])
	assert.Equal(t, "test-user2@test.com", resbody[1]["email"])
	assert.NotEmpty(t, resbody[1]["created_at"])
	assert.NotEmpty(t, resbody[1]["updated_at"])
	assert.Empty(t, resbody[1]["deleted_at"])

	// テストデータ削除処理
	clearTestDB(t)
}
