package utils

import (
	"fmt"
	_ "reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateStruct(data interface{}) error {
	if err := validate.Struct(data); err != nil {
		fieldErrors := err.(validator.ValidationErrors)
		var errors []string
		for _, err := range fieldErrors {
			errors = append(errors, err.Field()+" is invalid")
		}
		str := strings.Join(errors, ", ")
		return fmt.Errorf("%s", str)
	}
	return nil
}
