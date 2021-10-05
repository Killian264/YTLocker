package helpers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"

	"github.com/Killian264/YTLocker/golocker/models"
)

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

func EncryptString(str string, key string) string {
	keyBytes, _ := hex.DecodeString(key)
	plaintext := []byte(str)

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		panic(err.Error())
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return fmt.Sprintf("%x", ciphertext)
}

func DecryptString(str string, key string) string {
	keyBytes, _ := hex.DecodeString(key)
	enc, _ := hex.DecodeString(str)

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		panic(err.Error())
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	nonceSize := aesGCM.NonceSize()

	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	return fmt.Sprintf("%s", plaintext)
}
