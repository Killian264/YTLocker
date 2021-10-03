package data

import (
	"github.com/Killian264/YTLocker/golocker/models"
)

func (d *Data) NewYoutubeAccount(account models.YoutubeAccount) (models.YoutubeAccount, error) {
	account.ID = d.rand.ID()

	account.YoutubeToken.ID = d.rand.ID()

	result := d.db.Create(&account)

	return account, result.Error
}

func (d *Data) NewUserYoutubeAccount(userID uint64, accountID uint64) error {
	user := models.User{ID: userID}
	account := models.YoutubeAccount{ID: accountID}

	return d.db.Model(&user).Association("YoutubeAccounts").Append(&account)
}

func (d *Data) GetUserYoutubeAccounts(user models.User) ([]models.YoutubeAccount, error) {
	accounts := []models.YoutubeAccount{}

	result := d.db.Model(&user).Preload("YoutubeToken").Association("YoutubeAccounts").Find(&accounts)

	return accounts, result
}

func (d *Data) GetAccountByEmail(email string) (models.YoutubeAccount, error) {
	account := models.YoutubeAccount{}

	result := d.db.Preload("YoutubeToken").Where("email = ?", email).First(&account)

	return account, result.Error
}

func (d *Data) GetAccount(accountId uint64) (models.YoutubeAccount, error) {
	account := models.YoutubeAccount{}

	result := d.db.Preload("YoutubeToken").Where("id = ?", accountId).First(&account)

	return account, result.Error
}

func (d *Data) UpdateAccount(account models.YoutubeAccount, token models.YoutubeToken) (models.YoutubeAccount, error) {
	result := d.db.Model(&account).Select("PermissionLevel").Updates(account)
	if result.Error != nil {
		return models.YoutubeAccount{}, result.Error
	}

	token.ID = account.YoutubeToken.ID

	result = d.db.Model(&token).Select("AccessToken", "TokenType", "RefreshToken", "Expiry").Updates(&token)

	account.YoutubeToken = token

	return account, result.Error
}
