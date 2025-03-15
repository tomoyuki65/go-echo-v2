package user

import (
	"context"
	"time"

	"go-echo-v2/database"
	"go-echo-v2/ent"
	entUser "go-echo-v2/ent/user"
)

// インターフェースの定義
type UserRepository interface {
	CreateUser(
		ctx context.Context,
		uid string,
		firstName string,
		lastName string,
		email string,
	) (*ent.User, error)
	GetAllUsers(ctx context.Context) ([]*ent.User, error)
	GetUserByUID(ctx context.Context, uid string) (*ent.User, error)
	GetUserByEmail(ctx context.Context, email string) (*ent.User, error)
	UpdateUserByUID(
		ctx context.Context,
		uid string,
		lastName string,
		firstName string,
		email string,
	) (*ent.User, error)
	DeleteUserByUID(ctx context.Context, uid string) error
}

// 構造体の定義
type userRepository struct{}

// インスタンス生成用関数の定義
func NewUserRepository() UserRepository {
	return &userRepository{}
}

// メソッドの実装
func (u *userRepository) CreateUser(
	ctx context.Context,
	uid string,
	lastName string,
	firstName string,
	email string,
) (*ent.User, error) {
	db, err := database.SetupDatabase(ctx)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	user, err := db.User.Create().
		SetUID(uid).
		SetLastName(lastName).
		SetFirstName(firstName).
		SetEmail(email).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userRepository) GetAllUsers(ctx context.Context) ([]*ent.User, error) {
	db, err := database.SetupDatabase(ctx)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	users, err := db.User.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *userRepository) GetUserByUID(ctx context.Context, uid string) (*ent.User, error) {
	db, err := database.SetupDatabase(ctx)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	user, err := db.User.Query().
		Where(entUser.UIDEQ(uid)).
		Where(entUser.DeletedAtIsNil()).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userRepository) GetUserByEmail(ctx context.Context, email string) (*ent.User, error) {
	db, err := database.SetupDatabase(ctx)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	user, err := db.User.Query().
		Where(entUser.EmailEQ(email)).
		Where(entUser.DeletedAtIsNil()).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userRepository) UpdateUserByUID(
	ctx context.Context,
	uid string,
	lastName string,
	firstName string,
	email string,
) (*ent.User, error) {
	db, err := database.SetupDatabase(ctx)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	user, err := db.User.Query().
		Where(entUser.UIDEQ(uid)).
		Where(entUser.DeletedAtIsNil()).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	updateUser := user.Update()
	if lastName != "" {
		updateUser = updateUser.SetLastName(lastName)
	}
	if firstName != "" {
		updateUser = updateUser.SetFirstName(firstName)
	}
	if email != "" {
		updateUser = updateUser.SetEmail(email)
	}

	user, err = updateUser.Save(ctx)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userRepository) DeleteUserByUID(ctx context.Context, uid string) error {
	db, err := database.SetupDatabase(ctx)
	if err != nil {
		return err
	}
	defer db.Close()

	user, err := db.User.Query().
		Where(entUser.UIDEQ(uid)).
		Where(entUser.DeletedAtIsNil()).
		Only(ctx)
	if err != nil {
		return err
	}

	// 現在の日時を文字列で取得
	date := time.Now()
	dateString := date.Format("2006-01-02 15:04:05")

	// 更新用のemailの値を設定
	updateEmail := user.Email + dateString

	_, err = user.Update().
		SetEmail(updateEmail).
		SetDeletedAt(date).
		Save(ctx)
	if err != nil {
		return err
	}

	return nil
}
