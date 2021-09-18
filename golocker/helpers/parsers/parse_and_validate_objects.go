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
// returns the parsed playlist and an error string if an error occurred
// the error is user safe
func ParseAndValidatePlaylist(playlist models.Playlist) (models.Playlist, string) {
	validColors := map[string]bool{
		"red-1": true, 
		"yellow-1": true, 
		"green-1": true, 
		"blue-1": true, 
		"purple-1": true,
		"pink-1": true,
	}

	_, found := validColors[playlist.Color];

	if(playlist.Color == "" || !found){
		return models.Playlist{}, "a valid color must be provided"
	}

	return models.Playlist{
		Title:       SanitizeString(playlist.Title),
		Description: SanitizeString(playlist.Description),
		Color: playlist.Color,
	}, ""
}
