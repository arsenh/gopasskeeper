package storage

import (
	"errors"
	"fmt"
	"gopasskeeper/constants"
	"gopasskeeper/secure"
	"io"
	"log"
	"os"
	"sync"
)

var passwordsFile *os.File
var once sync.Once

var (
	ErrPasswordFileIsEmpty = errors.New(constants.ErrPasswordFileIsEmpty)
	ErrInvalidJsonFormat   = errors.New(constants.ErrInvalidJsonFormat)
	ErrMasterHashIsEmpty   = errors.New(constants.ErrMasterHashIsEmpty)
)

func IsPasswordFileIsEmpty(err error) bool {
	return errors.Is(err, ErrPasswordFileIsEmpty)
}

func IsInvalidJsonFormat(err error) bool {
	return errors.Is(err, ErrInvalidJsonFormat)
}

func IsMasterHashIsEmpty(err error) bool {
	return errors.Is(err, ErrMasterHashIsEmpty)
}

func GetPasswordFile() *os.File {
	once.Do(func() {
		passwordsFile = setupPasswordFileIfNeeded()
		if passwordsFile == nil {
			log.Fatal(constants.ErrCantCreateFileInHomeDir)
		}
	})
	return passwordsFile
}

func IsMasterPasswordHashAlreadyExist() (*PasswordJson, error) {
	passwordFileContent := getPasswordFileContent()

	if passwordFileContent == "" {
		return nil, ErrPasswordFileIsEmpty
	}

	passwordJson, err := DeserializePasswordDataFromJson(passwordFileContent)

	if err != nil {
		return nil, ErrInvalidJsonFormat
	}

	if passwordJson.MasterKeyHash == "" {
		return nil, ErrMasterHashIsEmpty
	}

	return &passwordJson, nil
}

func RestoreMasterPasswordHash(masterPassword string) bool {

	passwordJson, err := IsMasterPasswordHashAlreadyExist()
	if err != nil {
		return false
	}

	// restore password hash
	return secure.RestorePasswordHash(masterPassword, passwordJson.MasterKeyHash)
}

func StoreMasterPassword(masterPassword string) {
	/*
		from user password will be generated hash,
		hash will encrypted using AES256
		encrypted data will be used for
		encrypt/decrypt operations for credentials.
	*/
	masterPasswordHash := secure.InitializePasswordHash(masterPassword)

	// create initial data for password json file
	data := PasswordJson{
		MasterKeyHash: masterPasswordHash,
		Data:          "",
	}

	jsonData := SerializePasswordDataToJson(data)
	StoreJsonDataToPasswordFile(jsonData)
}

func StoreJsonDataToPasswordFile(jsonData string) {
	passwordsFile.WriteString(jsonData)
}

func GetPasswordFilePath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(constants.ErrGetHomaDir)
	}
	return homeDir + string(os.PathSeparator) + constants.PasswordFileName
}

func getPasswordFileContent() string {
	passwordsFile.Seek(0, 0)
	bytes, _ := io.ReadAll(passwordsFile)
	return string(bytes)
}

func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil || !os.IsNotExist(err)
}

func setupPasswordFileIfNeeded() *os.File {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(constants.ErrGetHomaDir)
	}

	applicationDirPath := homeDir + string(os.PathSeparator) + constants.ApplicationFolderName
	passwordFilePath := homeDir + string(os.PathSeparator) + constants.PasswordFileName

	if !fileExists(applicationDirPath) {
		err = os.Mkdir(applicationDirPath, 0755)
		if err != nil {
			log.Fatal(constants.ErrGetHomaDir)
		}
		fmt.Printf(constants.ApplicationDirNotExistCreateNew, applicationDirPath)
	}

	if !fileExists(passwordFilePath) {
		passwordsFile, err = os.OpenFile(passwordFilePath, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0755)
		if err != nil {
			log.Fatal(constants.ErrCreateFileInHomeDir)
		}
		fmt.Printf(constants.PasswordFileCreatedMsg, passwordFilePath)
	} else {
		passwordsFile, err = os.OpenFile(passwordFilePath, os.O_RDWR, 0755)
		if err != nil {
			log.Fatalf(constants.ErrOpenFile, passwordFilePath)
		}
	}
	return passwordsFile
}
