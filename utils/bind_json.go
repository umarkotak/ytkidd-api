package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New(validator.WithRequiredStructEnabled())
)

func BindJson(r *http.Request, dest any) error {
	bodyByte, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bodyByte, dest)
	if err != nil {
		return err
	}

	err = validate.Struct(dest)
	if err != nil {
		errStr := []string{}
		for _, err := range err.(validator.ValidationErrors) {
			// fmt.Println(err.Namespace())
			// fmt.Println(err.Field())
			// fmt.Println(err.StructNamespace())
			// fmt.Println(err.StructField())
			// fmt.Println(err.Tag())
			// fmt.Println(err.ActualTag())
			// fmt.Println(err.Kind())
			// fmt.Println(err.Type())
			// fmt.Println(err.Value())
			// fmt.Println(err.Param())
			errStr = append(errStr, fmt.Sprintf("%v: %v", err.Field(), err.Tag()))
		}

		return fmt.Errorf(strings.Join(errStr, ", "))
	}

	return nil
}

func ValidateStruct(obj any) error {
	err := validate.Struct(obj)
	if err != nil {
		errStr := []string{}
		for _, err := range err.(validator.ValidationErrors) {
			errStr = append(errStr, fmt.Sprintf("%v: %v", err.Field(), err.Tag()))
		}

		return fmt.Errorf(strings.Join(errStr, ", "))
	}

	return nil
}
