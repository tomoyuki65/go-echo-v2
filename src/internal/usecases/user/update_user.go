package user

import (
	"fmt"
	"net/http"

	"go-echo-v2/internal/services/user"
	utilContext "go-echo-v2/util/context"
	"go-echo-v2/util/logger"

	"github.com/labstack/echo/v4"
)

// 型定義
type UpdateUserResponse struct {
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
type UpdateUserUsecase interface {
	Exec(c echo.Context) error
}

// 構造体の定義
type updateUserUsecase struct {
	userService user.UserService
}

// インスタンス生成用関数の定義
func NewUpdateUserUsecase(
	userService user.UserService,
) UpdateUserUsecase {
	return &updateUserUsecase{
		userService: userService,
	}
}

// uidから対象のユーザーを更新
func (u *updateUserUsecase) Exec(c echo.Context) error {
	ctx := utilContext.CreateContext(c)

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

	user, err := u.userService.UpdateUserByUID(
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
	res := UpdateUserResponse{
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
