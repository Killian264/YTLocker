package parsers

import "github.com/Killian264/YTLocker/golocker/models"

// returns the parsed user and an error string if an error occurred
// the error is user safe
func ParseAndValidateUser(user models.User) (models.User, string) {

	if !StringArrayIsValid([]string{user.Username, user.Email, user.Password}) {
		return models.User{}, "user information contained invalid characters"
	}

	parsed := models.User{
		Username: SanitizeString(user.Username),
		Email:    SanitizeString(user.Email),
		Password: SanitizeString(user.Password),
	}

	if !IsEmailValid(parsed.Email) {
		return models.User{}, "user email is invalid"
	}

	if !IsPasswordValid(parsed.Password) {
		return models.User{}, "password does not meet requirements"
	}

	if len(parsed.Username) < 3 || 21 < len(parsed.Username) {
		return models.User{}, "username must be between 3 and 21 characters"
	}

	return parsed, ""

}

func ParsePlaylist(playlist models.Playlist) models.Playlist {

	return models.Playlist{
		Title:       SanitizeString(playlist.Title),
		Description: SanitizeString(playlist.Description),
	}

}
