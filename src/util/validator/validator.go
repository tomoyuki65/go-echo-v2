package validator

import (
	"github.com/go-playground/validator/v10"
)

type customValidator struct {
	validator *validator.Validate
}

func (cv *customValidator) Validate(i interface{}) error {
	// 必要であれば、ここでカスタムバリデーションルールを追加
	return cv.validator.Struct(i)
}

func NewCustomValidator() *customValidator {
	return &customValidator{
		validator: validator.New(),
	}
}
