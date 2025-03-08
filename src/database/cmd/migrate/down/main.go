package main

import (
	"context"
	"fmt"
	"os/exec"

	"go-echo-v2/database"
	"go-echo-v2/util/logger"

	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()

	// .env ファイルの読み込み
	err := godotenv.Load()
	if err != nil {
		logger.Error(ctx, ".envファイルの読み込みに失敗しました。")
	}

	// DBの接続情報取得
	dsn := database.CreateDsnForAtlas()

	// 一時的に使うDB接続情報
	devDsn := "postgres://testuser:test-password@pg-db:5432/tmpdb?search_path=public&sslmode=disable"

	// コマンド生成
	cmd := exec.Command("atlas", "migrate", "down", "--dir", "file://ent/migrate/migrations", "--url", dsn, "--dev-url", devDsn)

	// コマンド実行
	out, err := cmd.CombinedOutput()
	if err != nil {
		msg := fmt.Sprintf("コマンド実行で失敗しました。：%s\n", err.Error())
		logger.Error(ctx, msg)
	}

	// ログ出力
	fmt.Println(string(out))
}
