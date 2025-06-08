package user

import (
	"net/http"

	"go-echo-v2/internal/services/user"
	utilContext "go-echo-v2/util/context"

	"github.com/labstack/echo/v4"
)

// 型定義
type GetUserResponse struct {
	ID        int64  `json:"id"`
	UID       string `json:"uid"`
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

// インターフェースの定義
type GetUserUsecase interface {
	Exec(c echo.Context) error
}

// 構造体の定義
type getUserUsecase struct {
	userService user.UserService
}

// インスタンス生成用関数の定義
func NewGetUserUsecase(
	userService user.UserService,
) GetUserUsecase {
	return &getUserUsecase{
		userService: userService,
	}
}

// uidから対象のユーザーを取得
func (u *getUserUsecase) Exec(c echo.Context) error {
	ctx := utilContext.CreateContext(c)

	// パスパラメータからuidを取得
	uid := c.Param("uid")

	// TODO: 認可チェックを入れる

	user, err := u.userService.GetUserByUID(ctx, uid)
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
	res := GetUserResponse{
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
