package csv

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"unicode/utf8"

	utilContext "go-echo-v2/util/context"
	"go-echo-v2/util/logger"

	"github.com/jszwec/csvutil"
	"github.com/labstack/echo/v4"
)

// レスポンス結果の型定義
type ImportCsvData struct {
	No        string `csv:"no" json:"no" validate:"required"`
	LastName  string `csv:"last_name" json:"last_name" validate:"required"`
	FirstName string `csv:"first_name" json:"first_name" validate:"required"`
	Email     string `csv:"email" json:"email" validate:"required"`
}

// インターフェースの定義
type CsvUsecase interface {
	Exec(c echo.Context) error
}

// 構造体の定義
type csvUsecase struct{}

// インスタンス生成用関数の定義
func NewCsvUsecase() CsvUsecase {
	return &csvUsecase{}
}

// CSVファイルがutf8か判定する関数
func checkUTF8(fh *multipart.FileHeader) (bool, error) {
	f, err := fh.Open()
	if err != nil {
		return false, err
	}
	defer f.Close()

	buf := bytes.NewBuffer(nil)
	io.Copy(buf, f)
	b := buf.Bytes()

	return utf8.Valid(b), nil
}

// メソッド定義
func (u *csvUsecase) Exec(c echo.Context) error {
	ctx := utilContext.CreateContext(c)

	// CSVファイルを取得（CSVファイルにはヘッダー行が必要）
	fh, err := c.FormFile("csv")
	if err != nil {
		msg := fmt.Sprintf("CSVファイルが指定されていません。: %v", err)
		logger.Warn(ctx, msg)
		return echo.NewHTTPError(http.StatusBadRequest, msg)
	}

	// CSVファイルのutf8チェック
	isUTF8, err := checkUTF8(fh)
	if err != nil {
		msg := fmt.Sprintf("CSVファイルの文字コードチェックに失敗しました。: %v", err)
		logger.Error(ctx, msg)
		return echo.NewHTTPError(http.StatusInternalServerError, msg)
	}
	if !isUTF8 {
		msg := "CSVファイルがutf8形式ではありません。"
		logger.Warn(ctx, msg)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, msg)
	}

	// CSVファイルをオープン
	f, err := fh.Open()
	if err != nil {
		msg := fmt.Sprintf("CSVファイルのオープンに失敗しました: %v", err)
		logger.Error(ctx, msg)
		return echo.NewHTTPError(http.StatusInternalServerError, msg)
	}
	defer f.Close()

	// CSVリーダーの作成
	reader := csv.NewReader(f)

	// CSVデコーダーの作成
	decoder, err := csvutil.NewDecoder(reader)
	if err != nil {
		msg := fmt.Sprintf("CSVデコーダーの作成に失敗しました。: %v", err)
		logger.Error(ctx, msg)
		return echo.NewHTTPError(http.StatusInternalServerError, msg)
	}

	//　データマッピング
	var csvData []ImportCsvData
	if err := decoder.Decode(&csvData); err != nil {
		msg := fmt.Sprintf("データマッピングに失敗しました。: %v", err)
		logger.Error(ctx, msg)
		return echo.NewHTTPError(http.StatusInternalServerError, msg)
	}

	// バリデーションチェック
	var errMsgs []string
	for i, data := range csvData {
		// CSVデータをログ出力
		logger.Info(ctx, fmt.Sprintf("No: %s, LastName: %s, FirstName: %s, Email: %s", data.No, data.LastName, data.FirstName, data.Email))

		if err := c.Validate(data); err != nil {
			msg := fmt.Sprintf("%d件目[%v]", i+1, err)
			errMsgs = append(errMsgs, msg)
		}
	}

	var errMsgText string
	if len(errMsgs) > 0 {
		for _, msg := range errMsgs {
			errMsgText += msg + " "
		}
		msg := fmt.Sprintf("バリデーションエラー: %v", errMsgText)
		logger.Warn(ctx, msg)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, msg)
	}

	return c.JSON(http.StatusOK, csvData)
}
