package secure

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"gopasskeeper/constants"
	"io"
)

func EncryptAES(key []byte, plaintext string) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, 12) // GCM standard nonce size is 12 bytes
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	ciphertext := aesGCM.Seal(nonce, nonce, []byte(plaintext), nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func DecryptAES(key []byte, ciphertext string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	if len(data) < 12 {
		return "", errors.New(constants.ErrCiphertextTooShort)
	}

	nonce, ciphertext := data[:12], string(data[12:])
	plaintext, err := aesGCM.Open(nil, nonce, []byte(ciphertext), nil)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}

func GenerateAESKeyFromPassword(password string) []byte {

	// 32 byte is key length for AES256
	const keyLength = 32
	passwordBytes := []byte(password)

	if len(passwordBytes) == keyLength {
		// modifications is not needed, password lenght is ok
		return passwordBytes
	}

	if len(passwordBytes) > keyLength {
		// password is larger, need to take bytes equal to keyLength
		return passwordBytes[:32]
	}

	// remains case when password is small, need to add zeros
	newKeyBytes := []byte{}
	newKeyBytes = append(newKeyBytes, passwordBytes...)
	zeroBytes := make([]byte, keyLength-len(passwordBytes))
	newKeyBytes = append(newKeyBytes, zeroBytes...)
	return newKeyBytes
}
