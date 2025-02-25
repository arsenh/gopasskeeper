package interpreter

import (
	"fmt"
	"gopasskeeper/constants"
	"gopasskeeper/secure"
	"gopasskeeper/storage"
	"io"
	"log"
	"os"
	"strings"
	"syscall"

	"github.com/chzyer/readline"
	"golang.org/x/term"
)

func ReturnConfiguredReadLine() *readline.Instance {
	rl, err := readline.NewEx(&readline.Config{
		Prompt:          constants.UserPromptMsg,
		HistoryFile:     constants.HistoryFilePath,
		AutoComplete:    nil,
		InterruptPrompt: "^C",
		EOFPrompt:       "^D",
	})
	if err != nil {
		log.Fatal(constants.ErrInternalReadline, err)
	}
	return rl
}

func GetCommandPrompt() string {
	rl := ReturnConfiguredReadLine()
	defer rl.Close()

	for {
		enter, err := rl.Readline()

		if err == readline.ErrInterrupt || err == io.EOF { // ctrl+C || ctrl+D
			fmt.Printf(constants.UseExistCommandInsteadMsg, QUIT_COMMAND)
			continue
		}

		if err != nil {
			log.Fatal(constants.ErrInternalReadline)
		}

		fmt.Print("\n")
		enter = strings.TrimSpace(enter)

		if enter == "" {
			continue
		}
		return enter
	}
}

func ConfigLog() {
	log.SetFlags(0)
	log.SetOutput(os.Stderr)
	log.SetPrefix(constants.LoggerPrefixMsg)
}

func readPassword(prompt string) (string, error) {
	fmt.Print(prompt)
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", err
	}
	fmt.Println()
	return string(bytePassword), nil
}

func SetupFirstTimeUserMasterPassword() {
	masterPassword := ""
	masterConfirmPassword := ""
	var err error
	fmt.Println(constants.FirstTimeMasterPasswordSetupMsg)
	for {
		masterPassword, err = readPassword(constants.EnterMasterPasswordMsg)
		if err != nil {
			log.Fatal(constants.ErrFatalReadingMasterPassword)
		}

		if !secure.IsStrongPassword(masterPassword) {
			fmt.Println(constants.StrongPasswordMsg)
			continue
		}

		masterConfirmPassword, err = readPassword(constants.ConfirmMasterPasswordMsg)
		if err != nil {
			log.Fatal(constants.ErrFatalReadingMasterPassword)
		}

		if masterPassword != masterConfirmPassword {
			fmt.Println(constants.ErrPasswordMismatch)
			continue
		}
		break
	}
	storage.StoreMasterPassword(masterPassword)
	fmt.Printf(constants.PasswordMatchMessageMsg, storage.GetPasswordFilePath())
}

func SetupUserMasterPassword() {
	for {
		masterPassword, err := readPassword(constants.EnterMasterPasswordMsg)
		if err != nil {
			log.Fatal(constants.ErrInternalReadline)
		}
		if ok := storage.RestoreMasterPasswordHash(masterPassword); !ok {
			fmt.Println(constants.InvalidPasswordMsg)
			continue
		} else {
			fmt.Printf(constants.PasswordCorrectMsg, storage.GetPasswordFilePath())
			break
		}
	}
}

func SetupMasterPasswordIfNeeded() {
	_, err := storage.IsMasterPasswordHashAlreadyExist()

	if storage.IsPasswordFileIsEmpty(err) {
		SetupFirstTimeUserMasterPassword()
	} else if storage.IsInvalidJsonFormat(err) || storage.IsMasterHashIsEmpty(err) {
		fmt.Printf(constants.PasswordFileEmptyCorruptedMsg, storage.GetPasswordFilePath())
		os.Exit(1)
	} else {
		// get master password from user and restore master password hash.
		SetupUserMasterPassword()
	}
}

func Run() {
	ConfigLog()
	passwordFile := storage.GetPasswordFile()

	defer passwordFile.Close()

	if passwordFile == nil {
		log.Fatal(constants.ErrCantCreateFileInHomeDir)
	}

	defer func() {
		// this file must be deleted on appllication exit for security reasons
		err := os.Remove(constants.HistoryFilePath)
		if err != nil {
			log.Fatalf(constants.ErrUnableDeleteHistoryFile, constants.HistoryFilePath)
		}
	}()

	SetupMasterPasswordIfNeeded()

	for {
		prompt := GetCommandPrompt()
		action, err := Validate(prompt)
		if err != nil {
			fmt.Println(err)
			continue
		}
		action.Run()
	}
}
