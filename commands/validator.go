package commands

import (
	"errors"
	"gopasskeeper/actions"
	"regexp"
	"strings"
)

// commands
const (
	INVALID_COMMAND = "invalid input. use the 'help' command to view detailed instructions and additional information"
	QUIT_COMMAND    = "quit"
	HELP_COMMAND    = "help"
	ADD_COMMAND     = "add"
)

func parseInput(input string) (string, actions.Args) {
	re := regexp.MustCompile(`--(\w+)=({[^}]+}|"[^"]*"|\S+)`)
	words := strings.Fields(input)
	if len(words) == 0 {
		return "", nil
	}

	command := words[0]

	args := make(actions.Args)

	matches := re.FindAllStringSubmatch(input, -1)

	for _, match := range matches {
		key, value := match[1], match[2]

		// Remove surrounding quotes or curly braces if present
		value = strings.Trim(value, `"{}"`)

		args[key] = value
	}
	return command, args
}

func isValidAddCommandParameters(args actions.Args) bool {
	if args[actions.SERVICE_ARG] == "" || args[actions.USERNAME_ARG] == "" || args[actions.PASSWORD_ARG] == "" {
		return false
	}

	argsCount := len(args)
	_, hasNotes := args[actions.NOTES_ARG]

	return argsCount == 3 || (hasNotes && argsCount == 4)
}

func Validate(prompt string) (*actions.Action, error) {
	var args actions.Args = nil

	if prompt == QUIT_COMMAND {
		return actions.GetAction(actions.ACTION_QUIT, args), nil
	} else if prompt == HELP_COMMAND {
		return actions.GetAction(actions.ACTION_HELP, args), nil
	}

	command, args := parseInput(prompt)

	if command == "" || args == nil {
		return nil, errors.New(INVALID_COMMAND)
	}

	switch command {
	case ADD_COMMAND:
		isValid := isValidAddCommandParameters(args)
		if !isValid {
			return nil, errors.New(INVALID_COMMAND)
		}
		return actions.GetAction(actions.ACTION_ADD, args), nil
	default:
		return nil, errors.New(INVALID_COMMAND)
	}
}
