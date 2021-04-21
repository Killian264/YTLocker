package parsers

import (
	"regexp"
	"strings"
)

// StringIsValid checks if an array is valid
func StringArrayIsValid(strings []string) bool {

	for _, str := range strings {
		if !StringIsValid(str) {
			return false
		}
	}
	return true

}

// StringIsValid checks if a string is non empty and has no illegal characters
func StringIsValid(str string) bool {

	if str == "" {
		return false
	}

	sanitized := SanitizeString(str)

	return str == sanitized
}

// Sanitize string removes html tags and removes these characters ` " ' > < . ? \ * & ( ) ; : } {
func SanitizeString(str string) string {

	result := stripTags(str)
	result = stripIllegal(result)
	return result
}

func stripTags(str string) string {
	re := regexp.MustCompile(`<(.|\n)*?>`)

	result := re.ReplaceAll([]byte(str), []byte(""))

	return string(result)
}

func stripIllegal(str string) string {

	remove := []rune{'`', '"', '\'', '>', '<', '.', '?', '\\', '*', '&', '(', ')', ';', ':', '}', '{'}

	for _, char := range remove {
		str = strings.ReplaceAll(str, string(char), "")
	}

	return str
}