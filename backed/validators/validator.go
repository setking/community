package validators

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// ValidateEmail 验证邮箱
func ValidateEmail(fl validator.FieldLevel) bool {
	email := fl.Field().String()
	// 更严格的邮箱正则表达式
	pattern := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	matched, _ := regexp.MatchString(pattern, email)
	return matched
}

// 验证Direction 验证方向
func ValidateDirection(fl validator.FieldLevel) bool {
	value, ok := fl.Field().Interface().(float64)
	if !ok {
		return false
	}
	// 允许的值：1.0, 0.0, -1.0
	return value == 1.0 || value == 0.0 || value == -1.0
}
