package database

import (
	"context"
	"fmt"
	"os"

	"go-echo-v2/ent"

	"entgo.io/ent/dialect"
	_ "github.com/lib/pq" // DB driver
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateDsnForEnt() string {
	env := os.Getenv("ENV")

	// DBの接続情報設定
	var dsn string
	if env == "testing" {
		dsn = fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s password=%s search_path=public sslmode=disable",
			"pg-db",
			"5432",
			"testuser",
			"testdb",
			"test-password",
		)
	} else {
		host := os.Getenv("POSTGRES_HOST")
		port := os.Getenv("POSTGRES_PORT")
		user := os.Getenv("POSTGRES_USER")
		db := os.Getenv("POSTGRES_DB")
		pass := os.Getenv("POSTGRES_PASSWORD")

		if env == "production" {
			dsn = fmt.Sprintf(
				"host=%s port=%s user=%s dbname=%s password=%s search_path=public",
				host,
				port,
				user,
				db,
				pass,
			)
		} else {
			dsn = fmt.Sprintf(
				"host=%s port=%s user=%s dbname=%s password=%s search_path=public sslmode=disable",
				host,
				port,
				user,
				db,
				pass,
			)
		}
	}

	return dsn
}

func CreateDsnForAtlas() string {
	env := os.Getenv("ENV")

	// DBの接続情報設定
	var dsn string
	if env == "testing" {
		dsn = fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s?search_path=public&sslmode=disable",
			"testuser",
			"test-password",
			"pg-db",
			"5432",
			"testdb",
		)
	} else {
		host := os.Getenv("POSTGRES_HOST")
		port := os.Getenv("POSTGRES_PORT")
		user := os.Getenv("POSTGRES_USER")
		db := os.Getenv("POSTGRES_DB")
		pass := os.Getenv("POSTGRES_PASSWORD")

		if env == "production" {
			dsn = fmt.Sprintf(
				"postgres://%s:%s@%s:%s/%s?search_path=public",
				user,
				pass,
				host,
				port,
				db,
			)
		} else {
			dsn = fmt.Sprintf(
				"postgres://%s:%s@%s:%s/%s?search_path=public&sslmode=disable",
				user,
				pass,
				host,
				port,
				db,
			)
		}
	}

	return dsn
}

func SetupDatabase(ctx context.Context) (*ent.Client, error) {
	// DBの接続情報の取得
	dsn := CreateDsnForEnt()

	// DBクライアントの取得
	var client *ent.Client
	client, err := ent.Open(dialect.Postgres, dsn)
	if err != nil {
		return nil, err
	}

	return client, nil
}

// GORMで接続したい場合
func SetupDatabaseWithGorm(ctx context.Context) (*gorm.DB, error) {
	// DBの接続情報の取得
	dsn := CreateDsnForAtlas()

	// DBクライアントの取得
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
