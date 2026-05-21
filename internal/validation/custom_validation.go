package validation

import (
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

func CustomValidation(v *validator.Validate) error {
	blockedDomain := map[string]bool{
		"@blacklist.com": true,
		"@edu.vn":        true,
	}
	v.RegisterValidation("email_address", func(fl validator.FieldLevel) bool {
		part := strings.Split(fl.Field().String(), "@")
		if len(part) != 2 {
			return false
		}
		return !blockedDomain["@"+part[1]]
	})

	v.RegisterValidation("password_regex", func(fl validator.FieldLevel) bool {
		password := fl.Field().String()
		if len(password) < 8 {
			return false
		}
		hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
		hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
		hasDigit := regexp.MustCompile(`[0-9]`).MatchString(password)
		hasSpecial := regexp.MustCompile(`[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?~]`).MatchString(password)
		return hasLower && hasUpper && hasDigit && hasSpecial
	})

	regexSlug := regexp.MustCompile(`^[a-z0-9]+(?:[-.][a-z0-9]+)*$`)
	v.RegisterValidation("slug", func(fl validator.FieldLevel) bool {
		return regexSlug.MatchString(fl.Field().String())
	})
	regexSearch := regexp.MustCompile(`^[a-zA-Z0-9\\s]+$`)
	v.RegisterValidation("search", func(fl validator.FieldLevel) bool {
		return regexSearch.MatchString(fl.Field().String())
	})
	v.RegisterValidation("gte_int", func(fl validator.FieldLevel) bool {
		param, err := strconv.Atoi(fl.Param())
		if err != nil {
			return false
		}
		return fl.Field().Int() >= int64(param)
	})
	v.RegisterValidation("lte_int", func(fl validator.FieldLevel) bool {
		param, err := strconv.Atoi(fl.Param())
		if err != nil {
			return false
		}
		return fl.Field().Int() >= int64(param)
	})
	v.RegisterValidation("file_ext", func(fl validator.FieldLevel) bool {
		params := fl.Param()
		if params == "" {
			return false
		}
		allowedParam := strings.Fields(params)
		postfixField := strings.TrimPrefix(strings.ToLower(filepath.Ext(fl.Field().String())), ".")
		for _, allowed := range allowedParam {
			if postfixField == allowed {
				return true
			}
		}
		return false
	})
	return nil
}
