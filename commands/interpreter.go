package commands

import (
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/chzyer/readline"
)

func ReturnConfiguredReadLine() *readline.Instance {
	rl, err := readline.NewEx(&readline.Config{
		Prompt:          ">>>",
		HistoryFile:     "/tmp/gopasskeeper_history.tmp",
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

func Run() {
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
