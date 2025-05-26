package csv

import (
	usecaseCsv "go-echo-v2/internal/usecases/csv"

	"github.com/labstack/echo/v4"
)

// OpenAPI仕様書用の型定義
type ImportCsvResponse struct {
	No        string `json:"no" example:"1"`
	LastName  string `json:"last_name" example:"田中"`
	FirstName string `json:"first_name" example:"太郎"`
	Email     string `json:"email" example:"t.tanaka@example.com"`
}

type BadRequestResponse struct {
	Message string `json:"message" example:"Bad Request"`
}

type UnauthorizedResponse struct {
	Message string `json:"message" example:"Unauthorized"`
}

type UnprocessableEntityResponse struct {
	Message string `json:"message" example:"Unprocessable Entity"`
}

type InternalServerErrorResponse struct {
	Message string `json:"message" example:"Internal Server Error"`
}

// @Description CSVファイルのインポート用API
// @Tags csv
// @Accept multipart/form-data
// @Param file formData file true "CSV file to upload"
// @Security Bearer
// @Success 200 {object} ImportCsvResponse
// @Failure 400 {object} BadRequestResponse
// @Failure 401 {object} UnauthorizedResponse
// @Failure 422 {object} UnprocessableEntityResponse
// @Failure 500 {object} InternalServerErrorResponse
// @Router /api/v1/csv/import [post]
func ImportCsv(c echo.Context) error {
	// インスタンス生成
	csvUsecase := usecaseCsv.NewCsvUsecase()

	// ユースケースの実行
	return csvUsecase.Exec(c)
}
