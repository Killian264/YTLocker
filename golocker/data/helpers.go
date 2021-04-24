package data

import (
	"strings"
)

func notFound(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), "record not found")
}

func removeNotFound(err error) error {
	if notFound(err) {
		return nil
	}

	return err
}
