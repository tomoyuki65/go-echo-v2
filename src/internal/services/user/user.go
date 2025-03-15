package user

import (
	"fmt"
	"net/http"

	"go-echo-v2/internal/repositories/user"
	"go-echo-v2/util/logger"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// 型定義
type CreateUserRequestBody struct {
	LastName  string `json:"last_name" validate:"required"`
	FirstName string `json:"first_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
}

type CreateUserResponse struct {
	UID       string `json:"uid"`
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
	Email     string `json:"email"`
}

type UserResponse struct {
	ID        int64  `json:"id"`
	UID       string `json:"uid"`
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

type UpdateUserRequestBody struct {
	LastName  string `json:"last_name" validate:"omitempty"`
	FirstName string `json:"first_name" validate:"omitempty"`
	Email     string `json:"email" validate:"omitempty,email"`
}

// インターフェースの定義
type UserService interface {
	CreateUser(c echo.Context) error
	GetAllUsers(c echo.Context) error
	GetUserByUID(c echo.Context) error
	UpdateUserByUID(c echo.Context) error
	DeleteUserByUID(c echo.Context) error
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
func (s *userService) CreateUser(c echo.Context) error {
	ctx := c.Request().Context()

	var r CreateUserRequestBody
	if err := c.Bind(&r); err != nil {
		msg := fmt.Sprintf("リクエストボディが不正です。: %v", err)
		logger.Warn(ctx, msg)
		return echo.NewHTTPError(http.StatusBadRequest, msg)
	}

	// バリデーションチェック
	if err := c.Validate(&r); err != nil {
		msg := fmt.Sprintf("バリデーションエラー: %v", err)
		logger.Warn(ctx, msg)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, msg)
	}

	// UIDの設定
	uid := uuid.New().String()

	// ユーザー作成
	user, err := s.userRepository.CreateUser(ctx, uid, r.LastName, r.FirstName, r.Email)
	if err != nil {
		msg := fmt.Sprintf("ユーザーを作成できませんでした。: %v", err)
		logger.Error(ctx, msg)
		return echo.NewHTTPError(http.StatusInternalServerError, msg)
	}

	res := CreateUserResponse{
		UID:       user.UID,
		LastName:  user.LastName,
		FirstName: user.FirstName,
		Email:     user.Email,
	}

	return c.JSON(http.StatusCreated, res)
}

// 有効な全てのユーザーを取得
func (s *userService) GetAllUsers(c echo.Context) error {
	ctx := c.Request().Context()

	users, err := s.userRepository.GetAllUsers(ctx)
	if len(users) == 0 || err != nil {
		// データが０件またはエラーの場合は空の配列を返す
		res := []map[string]interface{}{}
		return c.JSON(http.StatusOK, res)
	}

	// レスポンス結果の整形
	var res []UserResponse
	for _, user := range users {
		// 日付項目のフォーマット変換
		createdAt := user.CreatedAt.Format("2006-01-02 15:04:05")
		updatedAt := user.UpdatedAt.Format("2006-01-02 15:04:05")
		deletedAt := ""
		if user.DeletedAt != nil {
			deletedAt = user.DeletedAt.Format("2006-01-02 15:04:05")
		}

		res = append(res, UserResponse{
			ID:        user.ID,
			UID:       user.UID,
			LastName:  user.LastName,
			FirstName: user.FirstName,
			Email:     user.Email,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
			DeletedAt: deletedAt,
		})
	}

	return c.JSON(http.StatusOK, res)
}

// uidから対象のユーザーを取得
func (s *userService) GetUserByUID(c echo.Context) error {
	ctx := c.Request().Context()

	// パスパラメータからuidを取得
	uid := c.Param("uid")

	// TODO: 認可チェックを入れる

	user, err := s.userRepository.GetUserByUID(ctx, uid)
	if user == nil || err != nil {
		// データが存在しない、またはエラーの場合は空のオブジェクトを返す
		res := map[string]interface{}{}
		return c.JSON(http.StatusOK, res)
	}

	// レスポンス結果の整形
	createdAt := user.CreatedAt.Format("2006-01-02 15:04:05")
	updatedAt := user.UpdatedAt.Format("2006-01-02 15:04:05")
	deletedAt := ""
	if user.DeletedAt != nil {
		deletedAt = user.DeletedAt.Format("2006-01-02 15:04:05")
	}
	res := UserResponse{
		ID:        user.ID,
		UID:       user.UID,
		LastName:  user.LastName,
		FirstName: user.FirstName,
		Email:     user.Email,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		DeletedAt: deletedAt,
	}

	return c.JSON(http.StatusOK, res)
}

// uidから対象のユーザーを更新
func (s *userService) UpdateUserByUID(c echo.Context) error {
	ctx := c.Request().Context()

	// パスパラメータからuidを取得
	uid := c.Param("uid")

	// TODO: 認可チェックを入れる

	var r UpdateUserRequestBody
	if err := c.Bind(&r); err != nil {
		msg := fmt.Sprintf("リクエストボディが不正です。: %v", err)
		logger.Warn(ctx, msg)
		return echo.NewHTTPError(http.StatusBadRequest, msg)
	}

	// バリデーションチェック
	if r.LastName == "" && r.FirstName == "" && r.Email == "" {
		msg := "リクエスボディが空です。"
		logger.Warn(ctx, msg)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, msg)
	}

	if err := c.Validate(&r); err != nil {
		msg := fmt.Sprintf("バリデーションエラー: %v", err)
		logger.Warn(ctx, msg)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, msg)
	}

	user, err := s.userRepository.UpdateUserByUID(
		ctx,
		uid,
		r.LastName,
		r.FirstName,
		r.Email,
	)
	if err != nil {
		msg := fmt.Sprintf("ユーザーを更新できませんでした。: %v", err)
		logger.Error(ctx, msg)
		return echo.NewHTTPError(http.StatusInternalServerError, msg)
	}

	// レスポンス結果の整形
	createdAt := user.CreatedAt.Format("2006-01-02 15:04:05")
	updatedAt := user.UpdatedAt.Format("2006-01-02 15:04:05")
	deletedAt := ""
	if user.DeletedAt != nil {
		deletedAt = user.DeletedAt.Format("2006-01-02 15:04:05")
	}
	res := UserResponse{
		ID:        user.ID,
		UID:       user.UID,
		LastName:  user.LastName,
		FirstName: user.FirstName,
		Email:     user.Email,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		DeletedAt: deletedAt,
	}

	return c.JSON(http.StatusOK, res)
}

// uidから対象ユーザーを論理削除
func (s *userService) DeleteUserByUID(c echo.Context) error {
	ctx := c.Request().Context()

	// パスパラメータからuidを取得
	uid := c.Param("uid")

	// TODO: 認可チェックを入れる

	err := s.userRepository.DeleteUserByUID(ctx, uid)
	if err != nil {
		msg := fmt.Sprintf("ユーザーを削除できませんでした。: %v", err)
		logger.Error(ctx, msg)
		return echo.NewHTTPError(http.StatusInternalServerError, msg)
	}

	res := map[string]interface{}{
		"message": "OK",
	}

	return c.JSON(http.StatusOK, res)
}
