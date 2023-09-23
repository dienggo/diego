package helper

import (
	"log"
	"reflect"
	"regexp"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
		panic(err)
		return
	}
}

// ReplaceSpecialCharacters : to replace special character
func ReplaceSpecialCharacters(str string, replace string) string {
	str = regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(str, replace)
	return str
}

// StructToMap is method to convert struct to map[string]any
func StructToMap(obj interface{}) map[string]any {
	result := make(map[string]any)
	val := reflect.ValueOf(obj)
	typ := val.Type()

	if val.Kind() != reflect.Struct {
		return nil // Only structs can be converted to maps
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := typ.Field(i).Name
		result[fieldName] = field.Interface()
	}

	return result
}
