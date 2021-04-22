package data

import "strings"

func NotFound(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), "record not found")
}

func RemoveNotFound(err error) error {
	if NotFound(err) {
		return nil
	}

	return err
}
