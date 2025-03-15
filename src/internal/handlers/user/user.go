package user

import (
	repoUser "go-echo-v2/internal/repositories/user"
	"go-echo-v2/internal/services/user"

	"github.com/labstack/echo/v4"
)

// OpenAPI仕様書用の型定義
type CreateUserRequestBody struct {
	LastName  string `json:"last_name" validate:"required" example:"山田"`
	FirstName string `json:"first_name" validate:"required" example:"太郎"`
	Email     string `json:"email" validate:"required,email" example:"t.yamada@example.com"`
}

type CreateUserResponse struct {
	UID       string `json:"uid" example:"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"`
	LastName  string `json:"last_name" example:"山田"`
	FirstName string `json:"first_name" example:"太郎"`
	Email     string `json:"email" example:"t.yamada@example.com"`
}

type UserResponse struct {
	ID        int64  `json:"id" example:"1"`
	UID       string `json:"uid" example:"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx`
	LastName  string `json:"last_name" example:"山田"`
	FirstName string `json:"first_name" example:"太郎"`
	Email     string `json:"email" example:"t.yamada@example.com"`
	CreatedAt string `json:"created_at" example:"2025-03-15 18:08:00"`
	UpdatedAt string `json:"updated_at" example:"2025-03-15 18:08:00"`
	DeletedAt string `json:"deleted_at" example:""`
}

type UpdateUserRequestBody struct {
	LastName  string `json:"last_name" example:"佐藤"`
	FirstName string `json:"first_name" example:"太郎"`
	Email     string `json:"email" validate:"email" example:"t.sato@example.com"`
}

type OKResponse struct {
	Message string `json:"message" example:"OK"`
}

type BadRequestResponse struct {
	Message string `json:"message" example:"リクエストボディが不正です。: error message"`
}

type UnauthorizedResponse struct {
	Message string `json:"message" example:"Unauthorized"`
}

type UnprocessableEntityResponse struct {
	Message string `json:"message" example:"バリデーションエラー: error message"`
}

type InternalServerErrorResponse struct {
	Message string `json:"message" example:"Internal Server Error: error message"`
}

// @Description ユーザー作成API
// @Tags user
// @Param CreateUserRequestBody body CreateUserRequestBody true "作成するユーザー情報"
// @Success 201 {object} CreateUserResponse
// @Failure 400 {object} BadRequestResponse
// @Failure 422 {object} UnprocessableEntityResponse
// @Failure 500 {object} InternalServerErrorResponse
// @Router /api/v1/user [post]
func CreateUser(c echo.Context) error {
	// インスタンス生成
	userRepository := repoUser.NewUserRepository()
	userService := user.NewUserService(userRepository)

	// サービス実行
	return userService.CreateUser(c)
}

// @Description 全てのユーザー取得API <br/> ※削除済みユーザー含む
// @Tags user
// @Security Bearer
// @Success 200 {object} []UserResponse "対象データが存在しない場合は空の配列「[]」を返す。"
// @Failure 401 {object} UnauthorizedResponse
// @Failure 500 {object} InternalServerErrorResponse
// @Router /api/v1/users [get]
func GetAllUsers(c echo.Context) error {
	// インスタンス生成
	userRepository := repoUser.NewUserRepository()
	userService := user.NewUserService(userRepository)

	// サービス実行
	return userService.GetAllUsers(c)
}

// @Description 有効な対象ユーザー取得API
// @Tags user
// @Security Bearer
// @Param uid path string true "uid"
// @Success 200 {object} UserResponse "対象データが存在しない場合は空のオブジェクト「{}」を返す。"
// @Failure 401 {object} UnauthorizedResponse
// @Failure 500 {object} InternalServerErrorResponse
// @Router /api/v1/user/:uid [get]
func GetUserByUID(c echo.Context) error {
	// インスタンス生成
	userRepository := repoUser.NewUserRepository()
	userService := user.NewUserService(userRepository)

	// サービス実行
	return userService.GetUserByUID(c)
}

// @Description 対象ユーザー更新API
// @Tags user
// @Security Bearer
// @Param uid path string true "uid"
// @Param UpdateUserRequestBody body UpdateUserRequestBody true "更新するユーザー情報"
// @Success 200 {object} UserResponse
// @Failure 400 {object} BadRequestResponse
// @Failure 401 {object} UnauthorizedResponse
// @Failure 422 {object} UnprocessableEntityResponse
// @Failure 500 {object} InternalServerErrorResponse
// @Router /api/v1/user/:uid [put]
func UpdateUserByUID(c echo.Context) error {
	// インスタンス生成
	userRepository := repoUser.NewUserRepository()
	userService := user.NewUserService(userRepository)

	// サービス実行
	return userService.UpdateUserByUID(c)
}

// @Description 対象ユーザー削除API
// @Tags user
// @Security Bearer
// @Param uid path string true "uid"
// @Success 200 {object} OKResponse
// @Failure 401 {object} UnauthorizedResponse
// @Failure 500 {object} InternalServerErrorResponse
// @Router /api/v1/user/:uid [delete]
func DeleteUserByUID(c echo.Context) error {
	// インスタンス生成
	userRepository := repoUser.NewUserRepository()
	userService := user.NewUserService(userRepository)

	// サービス実行
	return userService.DeleteUserByUID(c)
}
