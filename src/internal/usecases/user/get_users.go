package user

import (
	"net/http"

	"go-echo-v2/internal/services/user"
	utilContext "go-echo-v2/util/context"

	"github.com/labstack/echo/v4"
)

// 型定義
type GetUsersResponse struct {
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
type GetUsersUsecase interface {
	Exec(c echo.Context) error
}

// 構造体の定義
type getUsersUsecase struct {
	userService user.UserService
}

// インスタンス生成用関数の定義
func NewGetUsersUsecase(
	userService user.UserService,
) GetUsersUsecase {
	return &getUsersUsecase{
		userService: userService,
	}
}

// 有効な全てのユーザーを取得
func (u *getUsersUsecase) Exec(c echo.Context) error {
	ctx := utilContext.CreateContext(c)

	users, err := u.userService.GetAllUsers(ctx)
	if len(users) == 0 || err != nil {
		// データが０件またはエラーの場合は空の配列を返す
		res := []map[string]interface{}{}
		return c.JSON(http.StatusOK, res)
	}

	// レスポンス結果の整形
	var res []GetUsersResponse
	for _, user := range users {
		// 日付項目のフォーマット変換
		createdAt := user.CreatedAt.Format("2006-01-02 15:04:05")
		updatedAt := user.UpdatedAt.Format("2006-01-02 15:04:05")
		deletedAt := ""
		if user.DeletedAt != nil {
			deletedAt = user.DeletedAt.Format("2006-01-02 15:04:05")
		}

		res = append(res, GetUsersResponse{
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
