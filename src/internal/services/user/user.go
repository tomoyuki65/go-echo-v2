package user

import (
	"context"
	"fmt"

	"go-echo-v2/ent"
	"go-echo-v2/internal/repositories/user"
	"go-echo-v2/util/logger"
)

// インターフェースの定義
type UserService interface {
	CreateUser(
		ctx context.Context,
		uid string,
		firstName string,
		lastName string,
		email string,
	) (*ent.User, error)
	GetAllUsers(ctx context.Context) ([]*ent.User, error)
	GetUserByUID(ctx context.Context, uid string) (*ent.User, error)
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
type userService struct {
	userRepository user.UserRepository
}

// インスタンス生成用関数の定義
func NewUserService(
	userRepository user.UserRepository,
) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

// ユーザー作成
func (s *userService) CreateUser(
	ctx context.Context,
	uid string,
	lastName string,
	firstName string,
	email string,
) (*ent.User, error) {
	user, err := s.userRepository.CreateUser(ctx, uid, lastName, firstName, email)

	return user, err
}

// 有効な全てのユーザーを取得
func (s *userService) GetAllUsers(ctx context.Context) ([]*ent.User, error) {
	users, err := s.userRepository.GetAllUsers(ctx)
	if err != nil {
		msg := fmt.Sprintf("Failed to userRepository.GetAllUsers: %v", err)
		logger.Warn(ctx, msg)
	}

	return users, err
}

// uidから対象のユーザーを取得
func (s *userService) GetUserByUID(ctx context.Context, uid string) (*ent.User, error) {
	user, err := s.userRepository.GetUserByUID(ctx, uid)
	if err != nil {
		msg := fmt.Sprintf("Failed to userRepository.GetUserByUID: %v", err)
		logger.Warn(ctx, msg)
	}

	return user, err
}

// uidから対象のユーザーを更新
func (s *userService) UpdateUserByUID(
	ctx context.Context,
	uid string,
	lastName string,
	firstName string,
	email string,
) (*ent.User, error) {
	user, err := s.userRepository.UpdateUserByUID(
		ctx,
		uid,
		lastName,
		firstName,
		email,
	)

	return user, err
}

// uidから対象ユーザーを論理削除
func (s *userService) DeleteUserByUID(ctx context.Context, uid string) error {
	err := s.userRepository.DeleteUserByUID(ctx, uid)

	return err
}
