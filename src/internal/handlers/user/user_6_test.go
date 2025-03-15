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
 * 対象ユーザー削除APIのテスト
 */

// 正常系
func TestDeleteUserByUID(t *testing.T) {
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
	v1.DELETE("/user/:uid", DeleteUserByUID, middleware.AuthMiddleware)

	// テスト用データ登録
	userSeeder(t)

	// テスト用リクエストの作成
	req := httptest.NewRequest(http.MethodDelete, "/api/v1/user/test-1", nil)
	token := "zMdtq_glzI7oqq8yXjMgEOW6XfrSUMFGqw"
	bearerToken := "Bearer " + token
	req.Header.Set("Authorization", bearerToken)
	rec := httptest.NewRecorder()

	// テスト実行
	e.ServeHTTP(rec, req)

	// レスポンス結果をJSON形式で取得
	var resbody map[string]interface{}
	if err := json.Unmarshal(rec.Body.Bytes(), &resbody); err != nil {
		t.Fatal(err)
	}

	// DBから対象ユーザーを取得
	user := getUserByUID(t, "test-1")

	// 検証
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "OK", resbody["message"])
	assert.NotEmpty(t, user.DeletedAt)

	// テストデータ削除処理
	clearTestDB(t)
}
