package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/sonymimic1/go-transfer/pkg/util"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {

	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		return util.IsSupportCurrency(currency)
	}

	return true
}
