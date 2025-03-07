package main

import (
	"fmt"
	"log/slog"
	"os"

	"go-echo-v2/middleware"
	"go-echo-v2/router"
	_ "go-echo-v2/docs"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title go-echo-v2 API
// @version 1.0
// @description Go言語（Golang）のフレームワーク「Echo」によるAPI開発サンプルのバージョン２
// @host localhost:8080
// @BasePath /api/v1
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

	// ルーティング設定
	router.SetupRouter(e)

	// API仕様書の設定
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// ポート番号の設定
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	startPort := fmt.Sprintf(":%s", port)

	// サーバー起動
	e.Logger.Fatal(e.Start(startPort))
}
