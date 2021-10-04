package oauthmanager

import (
	"fmt"

	"github.com/Killian264/YTLocker/golocker/data"
	"github.com/Killian264/YTLocker/golocker/helpers"
	"github.com/Killian264/YTLocker/golocker/helpers/parsers"
	"github.com/Killian264/YTLocker/golocker/models"
	"golang.org/x/oauth2"
)

// GetBaseConfig gets the base ytlocker config
func (m *OauthManager) GetBaseConfig() oauth2.Config {
	return m.config
}

// GetBaseYoutubeAccount gets the base ytlocker account that should be added to every user
func (m *OauthManager) GetBaseYoutubeAccount() (models.YoutubeAccount, error) {
	return m.GetAccountByEmail("dev-locker@ytlocker.com")
}

func (m *OauthManager) InitializeYTService(service IYoutubeService, accountId uint64) (IYoutubeService, error) {
	account, err := m.GetAccountById(accountId)
	if err != nil {
		return service, fmt.Errorf("failed to get account by id: %d, err: %s", accountId, err.Error())
	}

	token, err := service.Initialize(m.config, parsers.ParseYoutubeToken(account.YoutubeToken))
	if err != nil {
		return service, fmt.Errorf("failed to initalize yt service: %s", err.Error())
	}

	_, err = m.RefreshToken(account, token, account.PermissionLevel)
	if err != nil {
		return service, fmt.Errorf("failed to refresh token: %s", err.Error())
	}

	return service, nil
}

func (m *OauthManager) GetLoginAccount(token oauth2.Token, permissionLevel string) (models.YoutubeAccount, error) {
	if permissionLevel != PERMISSION_LEVEL_VIEW && permissionLevel != PERMISSION_LEVEL_MANAGE {
		return models.YoutubeAccount{}, fmt.Errorf("Invalid permission level")
	}

	token, err := m.youtube.Initialize(m.config, token)
	if err != nil {
		return models.YoutubeAccount{}, err
	}

	userDetails, err := m.youtube.GetUser()
	if err != nil {
		return models.YoutubeAccount{}, err
	}

	account, err := m.GetAccountByEmail(userDetails.Email)
	if err != nil && err != data.ErrorNotFound {
		return models.YoutubeAccount{}, err
	}
	if err == nil {
		account, err := m.RefreshToken(account, token, permissionLevel)
		if err != nil {
			return models.YoutubeAccount{}, err
		}
		return account, nil
	}

	account, err = m.createAccount(userDetails, token, permissionLevel)
	if err != nil {
		return models.YoutubeAccount{}, err
	}

	return account, nil
}

// CreateAccount creates a youtube account, takes in a token and a permissionLevel that can be view or manage
func (m *OauthManager) createAccount(user models.OAuthUserInfo, token oauth2.Token, permissionLevel string) (models.YoutubeAccount, error) {
	channel, err := m.youtube.GetChannel()
	if err != nil {
		return models.YoutubeAccount{}, err
	}

	account := models.YoutubeAccount{
		Username:        channel.Snippet.Title,
		Email:           user.Email,
		Picture:         parsers.ParseYTChannelGetThumbnail(channel),
		PermissionLevel: permissionLevel,
		YoutubeToken:    parsers.ParseYoutubeTokenToModel(token),
	}

	account, err = m.data.NewYoutubeAccount(account)
	if err != nil {
		return models.YoutubeAccount{}, err
	}

	return account, nil
}

// LinkAccount links an account to a user
func (m *OauthManager) LinkAccount(user models.User, account models.YoutubeAccount) error {
	return m.data.NewUserYoutubeAccount(user.ID, account.ID)
}

// GetUserAccountList gets all the accounts of a user
func (m *OauthManager) GetUserAccountList(user models.User) ([]models.YoutubeAccount, error) {
	return m.data.GetUserYoutubeAccounts(user)
}

// GetAccountById gets an account by an id
func (m *OauthManager) GetAccountById(accountID uint64) (models.YoutubeAccount, error) {
	return m.data.GetAccount(accountID)
}

// GetUserAccount gets a single user account
func (m *OauthManager) GetUserAccount(user models.User, accountId uint64) (models.YoutubeAccount, error) {
	accountList, err := m.GetUserAccountList(user)
	if err != nil {
		return models.YoutubeAccount{}, err
	}

	account, err := m.GetAccountById(accountId)
	if err != nil {
		return models.YoutubeAccount{}, err
	}

	if !helpers.IsAccountInArray(accountList, account) {
		return models.YoutubeAccount{}, data.ErrorNotFound
	}

	return account, nil
}

// GetAccountByEmail gets an account by email
func (m *OauthManager) GetAccountByEmail(email string) (models.YoutubeAccount, error) {
	return m.data.GetAccountByEmail(email)
}

// RefreshToken refreshes a token, used to update an account permssion level or to set the token after it is refreshed
func (m *OauthManager) RefreshToken(account models.YoutubeAccount, token oauth2.Token, permissionLevel string) (models.YoutubeAccount, error) {
	if permissionLevel != PERMISSION_LEVEL_VIEW && permissionLevel != PERMISSION_LEVEL_MANAGE {
		return models.YoutubeAccount{}, fmt.Errorf("invalid permission level passed to RefreshToken")
	}

	account.PermissionLevel = permissionLevel

	return m.data.UpdateAccount(account, parsers.ParseYoutubeTokenToModel(token))
}
