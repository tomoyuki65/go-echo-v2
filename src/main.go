package main

import (
	"fmt"
	"log/slog"
	"os"

	_ "go-echo-v2/docs"
	"go-echo-v2/middleware"
	"go-echo-v2/router"
	"go-echo-v2/util/validator"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title go-echo-v2 API
// @version 1.0
// @description Go言語（Golang）のフレームワーク「Echo」によるAPI開発サンプルのバージョン２
// @host localhost:8080
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and token.
func main() {
	// .env ファイルの読み込み
	err := godotenv.Load()
	if err != nil {
		slog.Error(".envファイルの読み込みに失敗しました。")
	}

	e := echo.New()

	// ミドルウェアの設定
	e.Use(middleware.RequestMiddleware)
	e.Use(middleware.LoggerMiddleware())
	e.Use(middleware.CorsMiddleware())
	e.Use(echoMiddleware.Recover()) // panic発生時にサーバー停止を防ぐ

	// バリデーション設定
	e.Validator = validator.NewCustomValidator()

	// ルーティング設定
	router.SetupRouter(e)

	// API仕様書の設定
	env := os.Getenv("ENV")
	if env != "production" {
		e.GET("/swagger/*", echoSwagger.WrapHandler)
	}

	// ポート番号の設定
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	startPort := fmt.Sprintf(":%s", port)

	// サーバー起動
	e.Logger.Fatal(e.Start(startPort))
}
