package utils

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

const URL_REGEX = `^(http|https):\/\/[^ "]+$`

func validateRegex(regex, value string) bool {
	reg := regexp.MustCompile(regex)
	return reg.Match([]byte(value))
}

func CheckUrlRegexFunc(fl validator.FieldLevel) bool {
	if value, ok := fl.Field().Interface().(string); ok {
		return validateRegex(URL_REGEX, value)
	}
	return false
}
