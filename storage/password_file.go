package storage

import (
	"fmt"
	"log"
	"os"
	"sync"
)

var passwordsFile *os.File

const (
	applicationFolderName = ".gopasskeeper"
	passwordFileName      = applicationFolderName + "/gopasskeeper.passwords"
)

var once sync.Once

func GetPasswordFile() *os.File {
	once.Do(func() {
		passwordsFile = setupPasswordFileIfNeeded()
		if passwordsFile == nil {
			log.Fatal("unsable to setup system files in home directory")
		}
	})
	return passwordsFile
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
		passwordsFile, err = os.Create(passwordFilePath)
		if err != nil {
			log.Fatal("unable to create file in home directory")
		}
		fmt.Printf("password file is created %s\n", passwordFilePath)
	} else {
		/*
			TODO: check that file is not modified manually.
		*/
	}
	return passwordsFile
}
