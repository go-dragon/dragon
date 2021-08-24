package validator

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

func NewValidator() *validator.Validate {
	validate := validator.New()
	validate.RegisterValidation("regex", regex)
	return validate
}

func regex(fl validator.FieldLevel) bool {
	dataAny := fl.Field().Interface()
	switch dataAny.(type) {
	case string:
		// 只有string进行正则匹配
	default:
		return false
	}
	data := dataAny.(string)
	if data == "" {
		return false
	}
	tagVal := fl.Param() // regex=[a-z]+中的[a-z]+
	rx := regexp.MustCompile(tagVal)
	return rx.MatchString(data)
}
