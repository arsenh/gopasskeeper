package commands

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

const (
	COMMAND_QUIT = iota
)

func CommandQuit() {
	os.Exit(0)
}

type actions map[int]func()

var actionRegistry = actions{
	COMMAND_QUIT: CommandQuit,
}

func GetCommandPrompt() (string, error) {
	for {
		fmt.Print(">>> ")
		reader := bufio.NewReader(os.Stdin)
		enter, err := reader.ReadString('\n')
		fmt.Print("\n")
		enter = strings.TrimSpace(enter)

		if enter == "" {
			continue
		}
		return enter, err
	}
}

func GetAction(command string) (func(), error) {
	// TODO: need to validate command variable content to determine which action need to call.
	if command == "quit" {
		return actionRegistry[COMMAND_QUIT], nil
	}
	return nil, errors.New("invalid input. please enter valid command")
}

func Run() {
	for {
		command, _ := GetCommandPrompt()
		action, err := GetAction(command)
		if err != nil {
			fmt.Println(err)
			continue
		}
		action()
	}
}
