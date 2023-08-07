package helper

import (
	"log"
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
