package data

import (
	"github.com/Killian264/YTLocker/golocker/models"
)

var (
	YTLOCKER_TOKEN_TYPE = "ytlocker"
	USER_TOKEN_TYPE     = "user"
)

// NewYoutubeClientConfig creates a new client config
func (d *Data) NewYoutubeClientConfig(config *models.YoutubeClientConfig) error {
	config.ID = d.rand.ID()

	result := d.db.Create(&config)

	return result.Error
}

func (d *Data) NewYoutubeAccount(account models.YoutubeAccount) (models.YoutubeAccount, error) {
	account.ID = d.rand.ID()

	account.YoutubeToken.ID = d.rand.ID()

	result := d.db.Create(&account)

	return account, result.Error
}

func (d *Data) NewUserYoutubeAccount(userID uint64, accountID uint64) error {
	user := &models.User{ID: userID}
	account := &models.YoutubeAccount{ID: accountID}

	return d.db.Model(user).Association("YoutubeAccount").Append(account)
}

func (d *Data) GetUserYoutubeAccounts(user models.User) ([]models.YoutubeAccount, error) {
	accounts := []models.YoutubeAccount{}

	result := d.db.Model(&user).Association("YoutubeAccounts").Find(&accounts)

	return accounts, result
}

func (d *Data) GetAccountFromToken(token models.YoutubeToken) (models.YoutubeAccount, error) {
	result := d.db.Where("access_token = ?", token.AccessToken).First(&token)
	if result.Error != nil {
		return models.YoutubeAccount{}, result.Error
	}
	account := models.YoutubeAccount{}

	result = d.db.First(&account, token.YoutubeAccountID)

	return account, result.Error
}

func (d *Data) GetAccountByEmail(email string) (models.YoutubeAccount, error) {
	account := models.YoutubeAccount{}

	result := d.db.Where("email = ?", email).First(&account)

	return account, result.Error
}

func (d *Data) GetAccount(accountId uint64) (models.YoutubeAccount, error) {
	account := models.YoutubeAccount{}

	result := d.db.First(&account, accountId)

	return account, result.Error
}

func (d *Data) UpdateAccount(account models.YoutubeAccount) (models.YoutubeAccount, error) {
	account.PermissionLevel = "manage"

	result := d.db.Model(&account).Select("PermissionLevel").Updates(account)

	return account, result.Error
}
