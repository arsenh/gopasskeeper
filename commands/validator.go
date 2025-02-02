package commands

import (
	"errors"
	"gopasskeeper/actions"
	"regexp"
	"strings"
)

// commands
const (
	QUIT_COMMAND     = "quit"
	HELP_COMMAND     = "help"
	ADD_COMMAND      = "add"
	EDIT_COMMAND     = "edit"
	DELETE_COMMAND   = "delete"
	GET_COMMAND      = "get"
	GENERATE_COMMAND = "generate"
)

// error messages
const (
	INVALID_COMMAND = "invalid input. use the 'help' command to view detailed instructions and additional information"
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
	allowedArgs := map[string]bool{
		actions.SERVICE_ARG:  true,
		actions.USERNAME_ARG: true,
		actions.PASSWORD_ARG: true,
		actions.NOTES_ARG:    true, // Optional
	}

	// Required parameters
	serviceSet := false
	usernameSet := false
	passwordSet := false

	for k, v := range args {
		// Reject any invalid parameter
		if !allowedArgs[k] {
			return false
		}

		// Ensure required arguments are set (not empty)
		if k == actions.SERVICE_ARG && v != "" {
			serviceSet = true
		}
		if k == actions.USERNAME_ARG && v != "" {
			usernameSet = true
		}
		if k == actions.PASSWORD_ARG && v != "" {
			passwordSet = true
		}
	}

	// SERVICE_ARG, USERNAME_ARG, and PASSWORD_ARG must be set and non-empty
	return serviceSet && usernameSet && passwordSet
}

func isValidEditCommandParameters(args actions.Args) bool {

	allowedArgs := map[string]bool{
		actions.SERVICE_ARG:  true,
		actions.USERNAME_ARG: true,
		actions.PASSWORD_ARG: true,
		actions.NOTES_ARG:    true, // Optional
	}

	serviceSet := false
	usernameOrPasswordSet := false

	for k, v := range args {
		if !allowedArgs[k] {
			return false
		}

		if k == actions.SERVICE_ARG && v != "" {
			serviceSet = true
		}

		if (k == actions.USERNAME_ARG || k == actions.PASSWORD_ARG) && v != "" {
			usernameOrPasswordSet = true
		}
	}

	// SERVICE_ARG, USERNAME_ARG, or PASSWORD_ARG must be set and non-empty
	return serviceSet && usernameOrPasswordSet
}

func isOnlyServiceParameterProvided(args actions.Args) bool {
	value, ok := args[actions.SERVICE_ARG]
	return ok && value != "" && len(args) == 1
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
	case EDIT_COMMAND:
		isValid := isValidEditCommandParameters(args)
		if !isValid {
			return nil, errors.New(INVALID_COMMAND)
		}
		return actions.GetAction(actions.ACTION_EDIT, args), nil

	case DELETE_COMMAND:
		isValid := isOnlyServiceParameterProvided(args)
		if !isValid {
			return nil, errors.New(INVALID_COMMAND)
		}
		return actions.GetAction(actions.ACTION_DELETE, args), nil
	case GET_COMMAND:
		isValid := isOnlyServiceParameterProvided(args)
		if !isValid {
			return nil, errors.New(INVALID_COMMAND)
		}
		return actions.GetAction(actions.ACTION_GET, args), nil
	default:
		return nil, errors.New(INVALID_COMMAND)
	}
}
