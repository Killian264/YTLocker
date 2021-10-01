package parsers

import (
	"testing"

	"github.com/Killian264/YTLocker/golocker/models"
	"github.com/stretchr/testify/assert"
)

func Test_ParseAndValidatePlaylist(t *testing.T) {
	playlist := models.Playlist{
		Title:       "test",
		Description: "",
		Color:       "red-1",
	}

	got, err := ParseAndValidatePlaylist(playlist)
	assert.Empty(t, err)
	assert.Equal(t, playlist, got)

	playlist.Color = "red-123"

	got, err = ParseAndValidatePlaylist(playlist)
	assert.NotEmpty(t, err)
}
