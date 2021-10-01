package oauthmanager

import (
	"golang.org/x/oauth2"
)

func (m OauthManager) GetBaseConfig() oauth2.Config {
	return m.config
}

func (m OauthManager) GetBaseToken() oauth2.Token {
	return m.token
}
