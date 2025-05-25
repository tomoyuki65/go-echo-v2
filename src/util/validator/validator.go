package validator

import (
	"fmt"
	"strconv"

	"github.com/go-playground/validator/v10"
)

type customValidator struct {
	validator *validator.Validate
}

// カラムのバイト数チェック
func validateColumnByteSize(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	param := fl.Param()

	checkByte, err := strconv.Atoi(param)
	if err != nil {
		return false
	}

	return len(value) <= checkByte
}

func (cv *customValidator) Validate(i interface{}) error {
	err := cv.validator.Struct(i)
	if err != nil {
		// カスタムエラーメッセージの設定
		var errMsg string
		for _, err := range err.(validator.ValidationErrors) {
			switch err.ActualTag() {
			case "byte-size":
				errMsg += fmt.Sprintf("%sは%sバイト以下で入力して下さい。", err.Field(), err.Param())
			default:
				errMsg += err.Error()
			}
		}

		return fmt.Errorf("%s", errMsg)
	}

	return nil
}

func NewCustomValidator() *customValidator {
	v := validator.New()
	// カスタムバリデーション設定
	v.RegisterValidation("byte-size", validateColumnByteSize)

	return &customValidator{
		validator: v,
	}
}
