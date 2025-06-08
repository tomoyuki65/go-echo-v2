package user

import (
	"fmt"
	"net/http"

	"go-echo-v2/internal/services/user"
	utilContext "go-echo-v2/util/context"
	"go-echo-v2/util/logger"

	"github.com/labstack/echo/v4"
)

// インターフェースの定義
type DeleteUserUsecase interface {
	Exec(c echo.Context) error
}

// 構造体の定義
type deleteUserUsecase struct {
	userService user.UserService
}

// インスタンス生成用関数の定義
func NewDeleteUserUsecase(
	userService user.UserService,
) DeleteUserUsecase {
	return &deleteUserUsecase{
		userService: userService,
	}
}

// uidから対象ユーザーを論理削除
func (u *deleteUserUsecase) Exec(c echo.Context) error {
	ctx := utilContext.CreateContext(c)

	// パスパラメータからuidを取得
	uid := c.Param("uid")

	// TODO: 認可チェックを入れる

	err := u.userService.DeleteUserByUID(ctx, uid)
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
