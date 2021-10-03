package parsers

import "github.com/Killian264/YTLocker/golocker/models"

// returns the parsed playlist and an error string if an error occurred
// the error is user safe
func ParseAndValidatePlaylist(playlist models.Playlist) (models.Playlist, string) {
	validColors := map[string]bool{
		"red-1":    true,
		"yellow-1": true,
		"green-1":  true,
		"blue-1":   true,
		"purple-1": true,
		"pink-1":   true,
	}

	_, found := validColors[playlist.Color]

	if playlist.Color == "" || !found {
		return models.Playlist{}, "a valid color must be provided"
	}

	return models.Playlist{
		Title:            SanitizeString(playlist.Title),
		Description:      SanitizeString(playlist.Description),
		Color:            playlist.Color,
		YoutubeAccountID: playlist.YoutubeAccountID,
		Active:           playlist.Active,
	}, ""
}
