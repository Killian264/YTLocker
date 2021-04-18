package parsers

import (
	"fmt"
	"regexp"
)

func ValidateStringArray(strings []string) error {

	for _, str := range strings {
		err := ValidateString(str)
		if err != nil {
			return err
		}
	}
	return nil

}

func ValidateString(str string) error {
	//TODO: may need separate sanatize function later anyway?
	if str == "" {
		return fmt.Errorf("Registration information cannot be empty")
	}

	re := regexp.MustCompile(`<(.|\n)*?>`)

	result := re.Find([]byte(str))

	if result != nil {
		return fmt.Errorf("Registration information cannot contain: " + string(result))
	}

	return nil
}
