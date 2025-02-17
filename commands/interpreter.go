package commands

import (
	"fmt"
	"gopasskeeper/storage"
	"io"
	"log"
	"os"
	"strings"

	"github.com/chzyer/readline"
)

var historyFilePath string = "/tmp/gopasskeeper_history.tmp"

func ReturnConfiguredReadLine() *readline.Instance {
	rl, err := readline.NewEx(&readline.Config{
		Prompt:          ">>>",
		HistoryFile:     historyFilePath,
		AutoComplete:    nil,
		InterruptPrompt: "^C",
		EOFPrompt:       "^D",
	})
	if err != nil {
		log.Fatal("internal: readLine error:", err)
	}
	return rl
}

func GetCommandPrompt() string {
	rl := ReturnConfiguredReadLine()
	defer rl.Close()

	for {
		enter, err := rl.Readline()

		if err == readline.ErrInterrupt || err == io.EOF { // ctrl+C || ctrl+D
			fmt.Printf("\nuse '%s' to exit instead.\n", QUIT_COMMAND)
			continue
		}

		if err != nil {
			log.Fatal("internal: something went wrong with read input")
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
	log.SetPrefix("gopasskeeper: ")
}

func Run() {
	ConfigLog()
	passwordFile := storage.GetPasswordFile()

	defer passwordFile.Close()

	if passwordFile == nil {
		log.Fatal("unable to create password file in home directory")
	}

	defer func() {
		// this file must be deleted on appllication exit for security reasons
		err := os.Remove(historyFilePath)
		if err != nil {
			log.Fatalf("unable to delete history file %s, please do it manually for security reasons", historyFilePath)
		}
	}()

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
