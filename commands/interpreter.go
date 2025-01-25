package commands

import (
	"bufio"
	"fmt"
	"gopasskeeper/actions"
	"os"
	"strings"
)

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

func GetAction(command string) (*actions.Action, error) {
	action, err := Validate(command)

	if err != nil {
		return nil, err
	}

	return action, nil
}

func Run() {
	for {
		command, _ := GetCommandPrompt()
		action, err := GetAction(command)
		if err != nil {
			fmt.Println(err)
			continue
		}
		action.Run()
	}
}
