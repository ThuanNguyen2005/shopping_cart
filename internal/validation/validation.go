package validation

import (
	"fmt"
	"strings"

	"23.usersapi/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func InitialValidation() error {
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return fmt.Errorf("failed")
	}
	CustomValidation(v)
	return nil
}

func HandlerValidation(err error) gin.H {
	if validationErr, ok := err.(validator.ValidationErrors); ok {
		Errors := make(map[string]string)
		for _, e := range validationErr {
			root := strings.Split(e.Namespace(), ".")[0]
			paramName := strings.TrimPrefix(e.Namespace(), root+".")
			parts := strings.Split(paramName, ".")
			for i, part := range parts {
				if strings.Contains(part, "[") {
					partName := part[:strings.Index(part, "[")]
					partIdx := part[strings.Index(part, "["):]
					parts[i] = utils.CamelToSnake(partName) + partIdx
				} else {
					parts[i] = utils.CamelToSnake(part)
				}
			}
			fieldPath := strings.Join(parts, ".")
			switch e.Tag() {
			case "gt":
				Errors[fieldPath] = fmt.Sprintf("%s phải lớn hơn %s.", fieldPath, e.Param())
			case "lt":
				Errors[fieldPath] = fmt.Sprintf("%s phải nhỏ hơn %s.", fieldPath, e.Param())
			case "uuid":
				Errors[fieldPath] = fmt.Sprintf("%s phải đúng định dạng uuid.", fieldPath)
			case "slug":
				Errors[fieldPath] = fmt.Sprintf("%s phải đúng định dạng slug.", fieldPath)
			case "oneof":
				Errors[fieldPath] = fmt.Sprintf("%s phải chứa một trong: %+v", fieldPath, e.Param())
			case "search":
				Errors[fieldPath] = fmt.Sprintf("%s phải đúng định dạng search.", fieldPath)
			case "email":
				Errors[fieldPath] = fmt.Sprintf("%s phải đúng định dạng Email.", fieldPath)
			case "max":
				Errors[fieldPath] = fmt.Sprintf("%s phải có độ dài ngắn hơn %s.", fieldPath, e.Param())
			case "min":
				Errors[fieldPath] = fmt.Sprintf("%s phải có độ dài dài hơn %s.", fieldPath, e.Param())
			case "gte":
				Errors[fieldPath] = fmt.Sprintf("%s phải lớn hơn hoặc bằng: %s ký tự.", fieldPath, e.Param())
			case "lte":
				Errors[fieldPath] = fmt.Sprintf("%s phải nhỏ hơn hoặc bằng: %s ký tự.", fieldPath, e.Param())
			case "required":
				Errors[fieldPath] = fmt.Sprintf("%s phải được nhập.", fieldPath)
			case "datetime":
				Errors[fieldPath] = fmt.Sprintf("%s phải có định dạng yyyy-mm-dd", fieldPath)
			case "gte_int":
				Errors[fieldPath] = fmt.Sprintf("%s phải có giá trị lớn hơn hoặc bằng %s.", fieldPath, e.Param())
			case "lte_int":
				Errors[fieldPath] = fmt.Sprintf("%s phải có giá trị nhỏ hơn hoặc bằng %s.", fieldPath, e.Param())
			case "file_ext":
				allowedExt := strings.Join(strings.Split(e.Param(), " "), ", ")
				Errors[fieldPath] = fmt.Sprintf("%s chỉ cho phép các đuôi %s.", fieldPath, allowedExt)
			case "email_address":
				Errors[fieldPath] = fmt.Sprintf("%s chứa domain bị cấm.", fieldPath)
			case "password_regex":
				Errors[fieldPath] = fmt.Sprintf("%s chứa 8 ký tự, ký tự in hoa, ký tự in thường, chữ số, ký tự đặc biệt.", fieldPath)
			}
		}
		return gin.H{
			"err": Errors,
		}
	}
	return gin.H{
		"err": err,
	}
}
