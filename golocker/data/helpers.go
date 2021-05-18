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

type OnlyID struct {
	ID uint64
}

func parseOnlyIDArray(ids []OnlyID) []uint64 {
	parsed := []uint64{}

	for _, id := range ids {
		parsed = append(parsed, id.ID)
	}

	return parsed;
}