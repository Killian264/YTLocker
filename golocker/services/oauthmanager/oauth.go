package oauthmanager

import (
	"fmt"

	"github.com/Killian264/YTLocker/golocker/data"
	"github.com/Killian264/YTLocker/golocker/helpers"
	"github.com/Killian264/YTLocker/golocker/helpers/parsers"
	"github.com/Killian264/YTLocker/golocker/models"
	"golang.org/x/oauth2"
)

func (m *OauthManager) GetBaseConfig() oauth2.Config {
	return m.config
}

func (m *OauthManager) GetBaseToken() oauth2.Token {
	return m.token
}

func (m *OauthManager) GetBaseYoutubeAccount() models.YoutubeAccount {
	return m.account
}

// CreateAccount creates a youtube account, takes in a token and a permissionLevel that can be view or manage
func (m *OauthManager) CreateAccount(token oauth2.Token, permissionLevel string) (models.YoutubeAccount, error) {
	m.youtube.Initialize(m.config, token)

	if permissionLevel != "view" && permissionLevel != "manage" {
		return models.YoutubeAccount{}, fmt.Errorf("invalid permission level passed to CreateAccount")
	}

	user, err := m.youtube.GetUser()
	if err != nil {
		return models.YoutubeAccount{}, err
	}

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

// LinkBaseAccounts links the base account and login account to a user
func (m *OauthManager) LinkBaseAccounts(user models.User, token oauth2.Token) error {
	account, err := m.GetAccountByEmail(user.Email)
	if err != nil && err != data.ErrorNotFound {
		return err
	}

	if err == data.ErrorNotFound {
		account, err = m.CreateAccount(token, "view")
		if err != nil {
			return err
		}

		err = m.LinkAccount(user, account)
		if err != nil {
			return err
		}
	}

	accountList, err := m.GetUserAccountList(user)
	if err != nil {
		return err
	}

	if helpers.IsAccountInArray(accountList, m.account) {
		return nil
	}

	return m.LinkAccount(user, m.account)
}

// LinkAccount links an account to a user
func (m *OauthManager) LinkAccount(user models.User, account models.YoutubeAccount) error {
	return m.data.NewUserYoutubeAccount(user.ID, account.ID)
}

func (m *OauthManager) GetUserAccountList(user models.User) ([]models.YoutubeAccount, error) {
	return m.data.GetUserYoutubeAccounts(user)
}

func (m *OauthManager) GetAccountById(accountID uint64) (models.YoutubeAccount, error) {
	return m.data.GetAccount(accountID)
}

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

// GetAccountFromAccessToken very specifically for the base account only
func (m *OauthManager) GetAccountByAccessToken(token models.YoutubeToken) (models.YoutubeAccount, error) {
	return m.data.GetAccountFromToken(token)
}

func (m *OauthManager) GetAccountByEmail(email string) (models.YoutubeAccount, error) {
	return m.data.GetAccountByEmail(email)
}

func (m *OauthManager) UpdateAccountToManage(account models.YoutubeAccount) (models.YoutubeAccount, error) {
	return m.data.UpdateAccount(account)
}
