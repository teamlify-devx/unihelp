package validator

import (
	"context"
	"unicode"

	"github.com/go-playground/validator/v10"
)

// Use a single instance of Validate, it caches struct info
var validate *validator.Validate

func init() {
	validate = validator.New()
	validate.RegisterValidation("strongpass", strongPass)
}

// strongPass checks that the field contains at least one uppercase, one lowercase and one digit.
func strongPass(fl validator.FieldLevel) bool {
	val := fl.Field().String()
	var hasUpper, hasLower, hasDigit bool
	for _, r := range val {
		switch {
		case unicode.IsUpper(r):
			hasUpper = true
		case unicode.IsLower(r):
			hasLower = true
		case unicode.IsDigit(r):
			hasDigit = true
		}
	}
	return hasUpper && hasLower && hasDigit
}

// ValidateStruct validates struct fields.
func ValidateStruct(ctx context.Context, s interface{}) error {
	return validate.StructCtx(ctx, s)
}
