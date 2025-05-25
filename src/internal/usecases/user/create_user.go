package user

import (
	"fmt"
	"net/http"

	"go-echo-v2/internal/services/user"
	utilContext "go-echo-v2/util/context"
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

// インターフェースの定義
type CreateUserUsecases interface {
	Exec(c echo.Context) error
}

// 構造体の定義
type createUserUsecases struct {
	userService user.UserService
}

// インスタンス生成用関数の定義
func NewCreateUserUsecases(
	userService user.UserService,
) CreateUserUsecases {
	return &createUserUsecases{
		userService: userService,
	}
}

// ユーザー作成
func (s *createUserUsecases) Exec(c echo.Context) error {
	ctx := utilContext.CreateContext(c)

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
	user, err := s.userService.CreateUser(ctx, uid, r.LastName, r.FirstName, r.Email)
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
