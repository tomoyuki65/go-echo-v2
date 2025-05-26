package csv

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"go-echo-v2/middleware"
	"go-echo-v2/util/validator"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

type ImportCsvData struct {
	No        string `csv:"no" validate:"required"`
	LastName  string `csv:"last_name" validate:"required"`
	FirstName string `csv:"first_name" validate:"required"`
	Email     string `csv:"email" validate:"required"`
}

func TestImportCsvOK(t *testing.T) {
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
	e.Use(middleware.RequestMiddleware)
	e.Use(middleware.LoggerMiddleware())
	e.Use(middleware.CorsMiddleware())
	e.Use(echoMiddleware.Recover())
	// バリデーター設定
	e.Validator = validator.NewCustomValidator()
	v1 := e.Group("/api/v1")
	v1.POST("/csv/import", ImportCsv, middleware.AuthMiddleware)

	// CSVファイルのパス
	csvFilePath := "data/test-data-ok.csv"

	// リクエストボディの作成
	reqBody := &bytes.Buffer{}
	writer := multipart.NewWriter(reqBody)
	csvData, err := os.ReadFile(csvFilePath)
	if err != nil {
		t.Fatalf("failed to read CSV file: %v", err)
	}

	part, err := writer.CreateFormFile("csv", filepath.Base(csvFilePath))
	if err != nil {
		t.Fatalf("failed to create form file: %v", err)
	}

	_, err = part.Write(csvData)
	if err != nil {
		t.Fatalf("failed to write CSV file to form file: %v", err)
	}
	writer.Close()

	// テスト用リクエストの作成
	req := httptest.NewRequest(http.MethodPost, "/api/v1/csv/import", reqBody)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	token := "zMdtq_glzI7oqq8yXjMgEOW6XfrSUMFGqw"
	bearerToken := "Bearer " + token
	req.Header.Set("Authorization", bearerToken)
	rec := httptest.NewRecorder()

	// テスト実行
	e.ServeHTTP(rec, req)

	// レスポンス結果のJSONを取得
	var resbody []map[string]interface{}
	if err := json.Unmarshal(rec.Body.Bytes(), &resbody); err != nil {
		t.Fatal(err)
	}

	// 検証
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, 2, len(resbody))
	assert.Equal(t, "1", resbody[0]["no"])
	assert.Equal(t, "田中", resbody[0]["last_name"])
	assert.Equal(t, "太郎", resbody[0]["first_name"])
	assert.Equal(t, "t.tanaka@example.com", resbody[0]["email"])
	assert.Equal(t, "2", resbody[1]["no"])
	assert.Equal(t, "佐々木", resbody[1]["last_name"])
	assert.Equal(t, "一郎", resbody[1]["first_name"])
	assert.Equal(t, "ichirou.sasaki@example.com", resbody[1]["email"])
}

func TestImportCsvValidErr(t *testing.T) {
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
	e.Use(middleware.RequestMiddleware)
	e.Use(middleware.LoggerMiddleware())
	e.Use(middleware.CorsMiddleware())
	e.Use(echoMiddleware.Recover())
	// バリデーター設定
	e.Validator = validator.NewCustomValidator()
	v1 := e.Group("/api/v1")
	v1.POST("/csv/import", ImportCsv, middleware.AuthMiddleware)

	// CSVファイルのパス
	csvFilePath := "data/test-data-err.csv"

	// リクエストボディの作成
	reqBody := &bytes.Buffer{}
	writer := multipart.NewWriter(reqBody)
	csvData, err := os.ReadFile(csvFilePath)
	if err != nil {
		t.Fatalf("failed to read CSV file: %v", err)
	}

	part, err := writer.CreateFormFile("csv", filepath.Base(csvFilePath))
	if err != nil {
		t.Fatalf("failed to create form file: %v", err)
	}

	_, err = part.Write(csvData)
	if err != nil {
		t.Fatalf("failed to write CSV file to form file: %v", err)
	}
	writer.Close()

	// テスト用リクエストの作成
	req := httptest.NewRequest(http.MethodPost, "/api/v1/csv/import", reqBody)
	req.Header.Set("Content-Type", writer.FormDataContentType())
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

	// 検証
	assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
	assert.Contains(t, resbody["message"], "バリデーションエラー")
	assert.Contains(t, resbody["message"], "1件目[Key: 'ImportCsvData.FirstName' Error:Field validation for 'FirstName' failed on the 'required' tag]")
	assert.Contains(t, resbody["message"], "2件目[Key: 'ImportCsvData.LastName' Error:Field validation for 'LastName' failed on the 'required' tagKey: 'ImportCsvData.FirstName' Error:Field validation for 'FirstName' failed on the 'required' tag]")
}
