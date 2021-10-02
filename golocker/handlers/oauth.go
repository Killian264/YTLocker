package handlers

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/Killian264/YTLocker/golocker/data"
	"github.com/Killian264/YTLocker/golocker/models"
	"github.com/Killian264/YTLocker/golocker/services"
	"github.com/Killian264/YTLocker/golocker/services/ytservice"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type stateDetails struct {
	Bearer    string
	ScopeType string
}

func OAuthAuthenticate(w http.ResponseWriter, r *http.Request, s *services.Services) Response {
	bearer := r.URL.Query().Get("bearer")

	scopes, accessType, _ := getOAuthDetails(bearer)

	config := s.OauthManager.GetBaseConfig()

	parameters := url.Values{}
	parameters.Add("client_id", config.ClientID)
	parameters.Add("scope", strings.Join(scopes, " "))
	parameters.Add("access_type", accessType)
	parameters.Add("redirect_uri", config.RedirectURL)
	parameters.Add("response_type", "code")
	parameters.Add("include_granted_scopes", "true")
	parameters.Add("state", bearer)

	return createOAuthRedirect(google.Endpoint.AuthURL, "", parameters)
}

func OAuthAuthenticateCallback(w http.ResponseWriter, r *http.Request, s *services.Services) Response {
	code := r.FormValue("code")
	bearer := r.FormValue("state")

	youtubeService := &ytservice.YTPlaylist{}
	config := s.OauthManager.GetBaseConfig()

	_, _, scope := getOAuthDetails(bearer)

	token, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		return createOAuthRedirect(s.Config.WebRedirectUrl, "failed to exchange oauth code to token err: "+err.Error(), url.Values{})
	}

	err = youtubeService.Initialize(config, *token)
	if err != nil {
		return createOAuthRedirect(s.Config.WebRedirectUrl, "failed to initialize youtube service : "+err.Error(), url.Values{})
	}

	userDetails, err := youtubeService.GetUser()
	if err != nil {
		return createOAuthRedirect(s.Config.WebRedirectUrl, "failed to get user info: "+err.Error(), url.Values{})
	}

	account, err := s.OauthManager.GetAccountByEmail(userDetails.Email)
	if err != nil && err != data.ErrorNotFound {
		return createOAuthRedirect(s.Config.WebRedirectUrl, "failed to get account from email: "+err.Error(), url.Values{})
	}

	if err == data.ErrorNotFound {
		account, err = s.OauthManager.CreateAccount(*token, scope)
		if err != nil {
			return createOAuthRedirect(s.Config.WebRedirectUrl, "failed to get create an account: "+err.Error(), url.Values{})
		}
	}

	// user login
	if scope == "view" {
		user := models.User{
			Username: account.Username,
			Email:    account.Email,
			Picture:  account.Picture,
		}

		user, err := s.User.Login(user)
		if err != nil {
			return createOAuthRedirect(s.Config.WebRedirectUrl, "failed to login: "+err.Error(), url.Values{})
		}

		err = s.OauthManager.LinkBaseAccounts(user, *token)
		if err != nil {
			return createOAuthRedirect(s.Config.WebRedirectUrl, "failed to link base accounts: "+err.Error(), url.Values{})
		}

		parameters := url.Values{}
		parameters.Add("bearer", user.Session.Bearer)
		return createOAuthRedirect(s.Config.WebRedirectUrl, "", parameters)
	}

	// account is already linked as view only
	if account.PermissionLevel != "manage" {
		account, err = s.OauthManager.UpdateAccountToManage(account)
		if err != nil {
			return createOAuthRedirect(s.Config.WebRedirectUrl, "failed to upgrade account permissions: "+err.Error(), url.Values{})
		}
		return createOAuthRedirect(s.Config.WebRedirectUrl, "", url.Values{})
	}

	// account is being added under manage permissions
	user, err := s.User.GetUserFromBearer(bearer)
	if err != nil {
		return createOAuthRedirect(s.Config.WebRedirectUrl, "failed to get user from bearer: "+err.Error(), url.Values{})
	}

	err = s.OauthManager.LinkAccount(user, account)
	if err != nil {
		return createOAuthRedirect(s.Config.WebRedirectUrl, "failed to create account: "+err.Error(), url.Values{})
	}

	return createOAuthRedirect(s.Config.WebRedirectUrl, "", url.Values{})
}

func createOAuthRedirect(rawurl string, message string, params url.Values) Response {
	params.Add("success", strconv.FormatBool((message == "")))
	params.Add("reason", "user account was not linked")

	URL, err := url.Parse(rawurl)
	if err != nil {
		panic("google auth url was invalid")
	}

	URL.RawQuery = params.Encode()

	return NewRedirectResponse(URL.String(), message)
}

func getOAuthDetails(bearer string) ([]string, string, string) {
	scopes := []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/youtube.readonly"}
	accessType := "online"
	scope := "view"

	if bearer != "" {
		scopes = append(scopes, "https://www.googleapis.com/auth/youtube")
		accessType = "offline"
		scope = "manage"
	}

	return scopes, accessType, scope
}
