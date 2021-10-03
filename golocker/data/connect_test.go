package data

import (
	"testing"
)

func Test_Definition_Is_Valid(t *testing.T) {
	data := InMemoryMySQLConnect()

	data.createTables()
}
