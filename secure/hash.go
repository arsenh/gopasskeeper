package secure

import (
	"gopasskeeper/constants"
	"log"
	"sync"

	"golang.org/x/crypto/bcrypt"
)

// This is hash generated from user password
// Hash will used for encrypt user credentials
var masterPasswordHash string
var once sync.Once

func InitializePasswordHash(masterPassword string) string {
	once.Do(func() {
		passwordHash := GeneratePasswordHash(masterPassword)
		aesKey := GenerateAESKeyFromPassword(masterPassword)
		encryptedPasswordHash, err := EncryptAES(aesKey, passwordHash)
		if err != nil {
			log.Fatal(constants.ErrInternalCrypto)
		}
		masterPasswordHash = encryptedPasswordHash
	})
	return masterPasswordHash
}

func RestorePasswordHash(masterPassword, encryptedPasswordHash string) bool {
	aesKey := GenerateAESKeyFromPassword(masterPassword)
	decryptedPasswordHash, err := DecryptAES(aesKey, encryptedPasswordHash)
	if err != nil {
		return false
	}
	if !CheckHashAndPasssword(masterPassword, decryptedPasswordHash) {
		return false
	}
	masterPasswordHash = decryptedPasswordHash
	return true
}

func GetMasterPasswordHash() string {
	return masterPasswordHash
}

func GeneratePasswordHash(password string) string {
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(constants.ErrInternalCrypto)
	}
	return string(hashBytes)
}

func CheckHashAndPasssword(hash string, password string) bool {
	return nil == bcrypt.CompareHashAndPassword([]byte(password), []byte(hash))
}
