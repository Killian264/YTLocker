package parsers

import (
	"testing"

	"github.com/Killian264/YTLocker/golocker/models"
	"github.com/stretchr/testify/assert"
)

func Test_ParseAndValidateUser(t *testing.T) {
	user := models.User{
		Username: "test-one",
		Email:    "cool@gmail.com",
		Password: "password123",
	}

	got, err := ParseAndValidateUser(user)
	assert.Empty(t, err)
	assert.Equal(t, user, got)

	badChars := user
	badChars.Username = "&%_*@)!($_("

	_, err = ParseAndValidateUser(badChars)
	assert.NotEmpty(t, err)

	badName := user
	badName.Username = "qq"

	_, err = ParseAndValidateUser(badName)
	assert.NotEmpty(t, err)

	badPass := user
	badPass.Password = "lll"

	_, err = ParseAndValidateUser(badPass)
	assert.NotEmpty(t, err)

	badEmail := user
	badEmail.Email = "wowee"

	_, err = ParseAndValidateUser(badEmail)
	assert.NotEmpty(t, err)
}

func Test_ParseAndValidatePlaylist(t *testing.T) {
	playlist := models.Playlist{
		Title: "test",
		Description: "",
		Color: "red-1",
	}

	got, err := ParseAndValidatePlaylist(playlist)
	assert.Empty(t, err)
	assert.Equal(t, playlist, got)

	playlist.Color = "red-123"

	got, err = ParseAndValidatePlaylist(playlist)
	assert.NotEmpty(t, err)
}
