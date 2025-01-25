package commands

import (
	"errors"
	"gopasskeeper/actions"
)

const QUIT_PROMPT = "quit"
const HELP_PROMPM = "help"

func Validate(command string) (*actions.Action, error) {
	var commandId int = 0
	var args actions.Args
	if command == QUIT_PROMPT {
		commandId = actions.ACTION_QUIT
		return actions.GetAction(commandId, args), nil
	} else if command == HELP_PROMPM {
		commandId = actions.ACTION_HELP
		return actions.GetAction(commandId, args), nil
	}
	return nil, errors.New("invalid input. please enter valid command")
}
