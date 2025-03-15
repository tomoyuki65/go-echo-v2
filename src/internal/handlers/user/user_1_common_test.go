package user

import (
	"context"
	"testing"

	"go-echo-v2/database"
	"go-echo-v2/ent"
	entUser "go-echo-v2/ent/user"
	"go-echo-v2/middleware"
	"go-echo-v2/util/validator"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

/*
 * ユーザーAPIのテスト用共通処理
 */

// テスト用echoの初期化処理
func initTestEcho() *echo.Echo {
	e := echo.New()

	// ミドルウェアの設定
	e.Use(middleware.RequestMiddleware)
	e.Use(middleware.LoggerMiddleware())
	e.Use(middleware.CorsMiddleware())
	e.Use(echoMiddleware.Recover())

	// バリデーター設定
	e.Validator = validator.NewCustomValidator()

	return e
}

// テストデータ削除処理
func clearTestDB(t *testing.T) {
	ctx := context.Background()

	db, err := database.SetupDatabase(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	// ユーザーデータ削除
	_, err = db.User.
		Delete().
		Exec(ctx)
	if err != nil {
		t.Fatal(err)
	}
}

// テスト用ユーザーデータ登録Seeder
func userSeeder(t *testing.T) {
	ctx := context.Background()

	db, err := database.SetupDatabase(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	// テスト用ユーザーデータ作成
	_, err = db.User.Create().
		SetUID("test-1").
		SetLastName("姓１").
		SetFirstName("名１").
		SetEmail("test-user1@test.com").
		Save(ctx)
	if err != nil {
		t.Fatal(err)
	}

	_, err = db.User.Create().
		SetUID("test-2").
		SetLastName("姓２").
		SetFirstName("名２").
		SetEmail("test-user2@test.com").
		Save(ctx)
	if err != nil {
		t.Fatal(err)
	}
}

// ユーザー取得
func getUserByUID(t *testing.T, uid string) *ent.User {
	ctx := context.Background()
	db, err := database.SetupDatabase(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	user, err := db.User.Query().
		Where(entUser.UIDEQ(uid)).
		Only(ctx)
	if err != nil {
		t.Fatal(err)
	}

	return user
}
