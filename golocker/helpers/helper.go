package helpers

import "github.com/Killian264/YTLocker/golocker/models"

func IsKeyInArray(array []uint64, key uint64) bool {
	for _, num := range array {
		if num == key {
			return true
		}
	}

	return false
}

func IsAccountInArray(array []models.YoutubeAccount, key models.YoutubeAccount) bool {
	for _, account := range array {
		if account.ID == key.ID {
			return true
		}
	}

	return false
}
