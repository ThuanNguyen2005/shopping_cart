package utils

import (
	"reflect"
	"regexp"
	"strings"
)

var ExtValidate = map[string]bool{
	".png":  true,
	".jpeg": true,
	".jpg":  true,
}
var MimeValidate = map[string]bool{
	"image/jpeg": true,
	"image/png":  true,
}

var (
	matchFirstCap = regexp.MustCompile("(.)[A-Z][a-z]+")
	matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")
)

func CamelToSnake(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(str, "${1}_${2}")
	return strings.ToLower(snake)
}

func NormalizeString(text string) string {
	return strings.ToLower(strings.TrimSpace(text))
}

func InputEmpty(input any, current any) {
	val := reflect.ValueOf(input)
	if !val.IsZero() {
		reflect.ValueOf(current).Elem().Set(val)
	}
}

func ConvertToInt32Pointer(a int32) *int32 {
	if a == 0 {
		return nil
	}
	v := int32(a)
	return &v
}
