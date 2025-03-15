package user

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

/*
 * ユーザー作成APIのテスト
 */

// 正常系
func TestCreateUserOK(t *testing.T) {
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
	v1.POST("/user", CreateUser)

	// テスト用リクエストの作成
	reqBody := map[string]interface{}{
		"last_name":  "姓",
		"first_name": "名",
		"email":      "mei.sei@test.com",
	}
	jsonReqBody, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal(err)
	}
	req := httptest.NewRequest(http.MethodPost, "/api/v1/user", bytes.NewBuffer(jsonReqBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	// テスト実行
	e.ServeHTTP(rec, req)

	// レスポンス結果をJSON形式で取得
	var resbody map[string]interface{}
	if err := json.Unmarshal(rec.Body.Bytes(), &resbody); err != nil {
		t.Fatal(err)
	}

	// 検証
	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.Equal(t, "姓", resbody["last_name"])
	assert.Equal(t, "名", resbody["first_name"])
	assert.Equal(t, "mei.sei@test.com", resbody["email"])

	// テストデータ削除処理
	clearTestDB(t)
}
