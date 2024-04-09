package utfunc

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
	"runtime"
	"strings"
)

// GetCurrentCaller returns the name of the current caller function.
func GetCurrentCaller(caller interface{}, level int) string {
	if caller != nil {
		callerType := fmt.Sprintf("%T", caller)
		parts := strings.Split(callerType, ".")
		if len(parts) > 0 {
			return fmt.Sprintf("%s.%s", parts[len(parts)-1], getCurrentFunctionName(level+1))
		}
	}
	return getCurrentFunctionName(level + 1)
}

func getCurrentFunctionName(level int) string {
	pc, _, _, _ := runtime.Caller(level + 1)
	name := runtime.FuncForPC(pc).Name()
	parts := strings.Split(name, ".")
	if len(parts) == 0 {
		return ""
	}
	return parts[len(parts)-1]
}

// CheckValidateStruct validates the fields of the given struct
func CheckValidateStruct(obj interface{}) error {
	v := validator.New()
	return validateStruct(obj, v)
}

func validateStruct(obj interface{}, v *validator.Validate) error {
	err := v.Struct(obj)
	if err != nil {
		var errorMessages []string
		for _, e := range err.(validator.ValidationErrors) {
			fieldName := getFieldJSONTag(obj, e.Field())
			if fieldName == "" {
				fieldName = e.Field()
			}
			errorMessages = append(errorMessages, fmt.Sprintf("%s: %s", fieldName, e.Tag()))
		}
		return fmt.Errorf(strings.Join(errorMessages, ", "))
	}
	return nil
}

func getFieldJSONTag(obj interface{}, fieldName string) string {
	t := reflect.TypeOf(obj)
	if t.Kind() == reflect.Struct {
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			if field.Name == fieldName {
				return field.Tag.Get("json")
			}
		}
	}
	return fieldName
}
