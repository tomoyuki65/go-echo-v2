package user

import (
	"bytes"
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
 * 対象ユーザー更新APIのテスト
 */

// 正常系
func TestUpdateUserByUIDOK(t *testing.T) {
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
	v1.PUT("/user/:uid", UpdateUserByUID, middleware.AuthMiddleware)

	// テスト用データ登録
	userSeeder(t)

	// テスト用リクエストの作成
	reqBody := map[string]interface{}{
		"last_name":  "更新姓",
		"first_name": "更新名",
		"email":      "update-user@test.com",
	}
	jsonReqBody, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal(err)
	}
	req := httptest.NewRequest(http.MethodPut, "/api/v1/user/test-1", bytes.NewBuffer(jsonReqBody))
	req.Header.Set("Content-Type", "application/json")
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

	// 検証
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "test-1", resbody["uid"])
	assert.Equal(t, "更新姓", resbody["last_name"])
	assert.Equal(t, "更新名", resbody["first_name"])
	assert.Equal(t, "update-user@test.com", resbody["email"])
	assert.NotEmpty(t, resbody["created_at"])
	assert.NotEmpty(t, resbody["updated_at"])
	assert.Empty(t, resbody["deleted_at"])

	// テストデータ削除処理
	clearTestDB(t)
}
