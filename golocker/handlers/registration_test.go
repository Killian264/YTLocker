package handlers

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserFromRequest(t *testing.T) {

	req, err := http.NewRequest("GET", "/adsfasdf/asdf", nil)
	assert.Nil(t, err)

	req.Header["Authorization"] = []string{"TEMP_API_BEARER"}

}
