package main

import (
	"context"
	"fmt"

	"go-echo-v2/database"
	"go-echo-v2/util/logger"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()

	// .env ファイルの読み込み
	err := godotenv.Load()
	if err != nil {
		logger.Error(ctx, ".envファイルの読み込みに失敗しました。")
	}

	// DBに接続してクライアント取得
	client, err := database.SetupDatabase(ctx)
	if err != nil {
		msg := fmt.Sprintf("DB接続に失敗しました。: %v", err)
		logger.Error(ctx, msg)
	}
	defer client.Close()

	// ユーザー登録
	uid := uuid.New().String()
	lastName := gofakeit.LastName()
	firstName := gofakeit.FirstName()
	email := gofakeit.Email()

	user, err := client.User.Create().
		SetUID(uid).
		SetLastName(lastName).
		SetFirstName(firstName).
		SetEmail(email).
		Save(ctx)
	if err != nil {
		msg := fmt.Sprintf("ユーザー登録に失敗しました。: %v", err)
		logger.Error(ctx, msg)
	}

	logger.Info(ctx, user.String())
}
