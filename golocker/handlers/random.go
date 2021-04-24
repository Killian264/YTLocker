package handlers

// HandleSubscriptionNoError handles a new subscription request wrap in a middleware that handles errors
// func CreatePlaylist(w http.ResponseWriter, r *http.Request, s *services.Services) error {

// 	user, err := s.User.GetUserFromRequest(r)
// 	if err != nil {
// 		return err
// 	}

// 	playlist := &models.Playlist{}

// 	err = json.NewDecoder(r.Body).Decode(playlist)
// 	if err != nil {
// 		return err
// 	}

// 	_, err = s.Playlist.Create(playlist, user)

// 	return err

// }
