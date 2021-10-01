package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

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
	scope := r.URL.Query().Get("scope")
	bearer := r.URL.Query().Get("bearer")

	scopes, accessType, scope := getOAuthDetails(scope)

	state := stateDetails{
		Bearer:    bearer,
		ScopeType: scope,
	}

	if state.ScopeType == "manage" && state.Bearer == "" {
		return createRedirect(s.Config.WebRedirectUrl, "scope 'manage' was provided without a bearer", url.Values{})
	}

	stateDetailsJson, err := json.Marshal(state)
	if err != nil {
		return createRedirect(s.Config.WebRedirectUrl, "failed to marshal state json data", url.Values{})
	}

	config := s.OauthManager.GetBaseConfig()

	parameters := url.Values{}
	parameters.Add("client_id", config.ClientID)
	parameters.Add("scope", strings.Join(scopes, " "))
	parameters.Add("access_type", accessType)
	parameters.Add("redirect_uri", config.RedirectURL)
	parameters.Add("response_type", "code")
	parameters.Add("include_granted_scopes", "true")
	parameters.Add("state", string(stateDetailsJson))

	return createRedirect(google.Endpoint.AuthURL, "", parameters)
}

func OAuthAuthenticateCallback(w http.ResponseWriter, r *http.Request, s *services.Services) Response {
	youtubeService := &ytservice.YTPlaylist{}
	config := s.OauthManager.GetBaseConfig()
	code := r.FormValue("code")

	state, err := parseOAuthCallbackDetails(r.FormValue("state"), code)
	if err != nil {
		return createRedirect(s.Config.WebRedirectUrl, err.Error(), url.Values{})
	}

	token, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		return createRedirect(s.Config.WebRedirectUrl, "failed to exchange oauth code to token err: "+err.Error(), url.Values{})
	}

	err = youtubeService.Initialize(config, *token)
	if err != nil {
		return createRedirect(s.Config.WebRedirectUrl, "failed to initialize youtube service : "+err.Error(), url.Values{})
	}

	userDetails, err := youtubeService.GetUser()
	if err != nil {
		return createRedirect(s.Config.WebRedirectUrl, "failed to get user info: "+err.Error(), url.Values{})
	}

	channel, err := youtubeService.GetChannel()
	if err != nil {
		return createRedirect(s.Config.WebRedirectUrl, "failed to get user channel information: "+err.Error(), url.Values{})
	}

	if state.ScopeType == "view" {
		user := models.User{
			Username: channel.Snippet.Title,
			Email:    userDetails.Email,
			Picture:  userDetails.Picture,
		}

		user, err := s.User.Login(user)
		if err != nil {
			return createRedirect(s.Config.WebRedirectUrl, "failed to login"+err.Error(), url.Values{})
		}

		// note that this is refreshed on page load for security
		http.SetCookie(w, &http.Cookie{
			Name:    "bearer",
			Value:   user.Session.Bearer,
			Expires: time.Now().Add(24 * time.Hour),
		})

		parameters := url.Values{}
		parameters.Add("bearer", user.Session.Bearer)
		return createRedirect(s.Config.WebRedirectUrl, "", parameters)
	}

	return BlankResponse(nil)
}

func createRedirect(rawurl string, message string, params url.Values) Response {
	params.Add("success", strconv.FormatBool((message == "")))
	params.Add("reason", "user account was not linked")

	URL, err := url.Parse(rawurl)
	if err != nil {
		panic("google auth url was invalid")
	}

	URL.RawQuery = params.Encode()

	return NewRedirectResponse(URL.String(), message)
}

func getOAuthDetails(scope string) ([]string, string, string) {
	scopes := []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/youtube.readonly"}
	accessType := "online"

	if scope != "view" {
		scopes = append(scopes, "https://www.googleapis.com/auth/youtube")
		accessType = "offline"
		scope = "manage"
	}

	return scopes, accessType, scope
}

func parseOAuthCallbackDetails(state string, code string) (stateDetails, error) {
	details := stateDetails{}

	if code == "" {
		return details, fmt.Errorf("code was not returned on callback")
	}

	err := json.Unmarshal([]byte(state), &details)
	if err != nil {
		return details, fmt.Errorf("failed to parse state details, error: " + err.Error())
	}

	if details.ScopeType != "view" && (details.ScopeType != "manage" || details.Bearer == "") {
		return details, fmt.Errorf("invalid state type and or bearer")
	}

	return details, nil
}
