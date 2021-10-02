package handlers

import (
	"net/http"

	"github.com/Killian264/YTLocker/golocker/models"
	"github.com/Killian264/YTLocker/golocker/services"
)

// AccountList gets a list of accounts
func AccountList(w http.ResponseWriter, r *http.Request, s *services.Services) Response {
	user := GetUserFromRequest(r)

	accounts, err := s.OauthManager.GetUserAccountList(user)
	if err != nil {
		return BlankResponse(err)
	}

	return NewResponse(http.StatusOK, accounts, "")
}

// AccountGet gets an account
func AccountGet(w http.ResponseWriter, r *http.Request, s *services.Services) Response {
	account := GetAccountFromRequest(r)

	account.YoutubeToken = models.YoutubeToken{}

	return NewResponse(http.StatusOK, account, "")
}
