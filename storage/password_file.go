package storage

import (
	"errors"
	"fmt"
	"gopasskeeper/secure"
	"io"
	"log"
	"os"
	"sync"
)

var fileCorruptedMsg string = `
	File Error: Data Corruption Detected. path: %s

	It appears that the file containing your data has become corrupted and cannot be restored. To resolve this issue, please follow these steps:

	Delete the Corrupted File: Locate the file and delete it manually from your system.
	Restart the Program: Relaunch the application, which will create a new file from scratch.
`

var passwordsFile *os.File
var once sync.Once

const (
	applicationFolderName = ".gopasskeeper"
	passwordFileName      = applicationFolderName + "/gopasskeeper_passwords.json"
)

var (
	ErrPasswordFileIsEmpty = errors.New("password file is empty")
	ErrInvalidJsonFormat   = errors.New("invalid json format")
	ErrMasterHashIsEmpty   = errors.New("master hash is empty")
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
			log.Fatal("unable to setup system files in home directory")
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
		log.Fatal("unable to get home directory")
	}
	return homeDir + string(os.PathSeparator) + passwordFileName
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
		log.Fatal("unable to get home directory")
	}

	applicationDirPath := homeDir + string(os.PathSeparator) + applicationFolderName
	passwordFilePath := homeDir + string(os.PathSeparator) + passwordFileName

	if !fileExists(applicationDirPath) {
		err = os.Mkdir(applicationDirPath, 0755)
		if err != nil {
			log.Fatal("unable to create home directory")
		}
		fmt.Printf("application directory not exist, created new one in %s\n", applicationDirPath)
	}

	if !fileExists(passwordFilePath) {
		passwordsFile, err = os.OpenFile(passwordFilePath, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0755)
		if err != nil {
			log.Fatal("unable to create file in home directory")
		}
		fmt.Printf("password file is created %s\n", passwordFilePath)
	} else {
		passwordsFile, err = os.OpenFile(passwordFilePath, os.O_RDWR, 0755)
		if err != nil {
			log.Fatalf("unable to open existing file in path %s\n", passwordFilePath)
		}
	}
	return passwordsFile
}
