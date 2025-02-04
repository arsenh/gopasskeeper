package commands

import (
	"errors"
	"gopasskeeper/actions"
	"gopasskeeper/helpers"
	"regexp"
	"strconv"
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

// command arguments
const (
	SERVICE_ARG    = "service"
	USERNAME_ARG   = "username"
	PASSWORD_ARG   = "password"
	NOTE_ARG       = "note"
	COMPLEXITY_ARG = "complexity"
	LENGTH_ARG     = "length"

	COMPLEXITY_UPPERCASE_ARG = "uppercase"
	COMPLEXITY_NUMBERS_ARG   = "numbers"
	COMPLEXITY_SYMBOLS_ARG   = "symbols"
)

// error messages
const (
	INVALID_COMMAND = "invalid input. use the 'help' command to view detailed instructions and additional information"
)

func parseInput(input string) (string, map[string]string) {
	re := regexp.MustCompile(`--(\w+)=({[^}]+}|"[^"]*"|\S+)`)
	words := strings.Fields(input)
	if len(words) == 0 {
		return "", nil
	}

	command := words[0]
	args := make(map[string]string)
	matches := re.FindAllStringSubmatch(input, -1)

	for _, match := range matches {
		key, value := match[1], match[2]

		// Remove surrounding quotes or curly braces if present
		value = strings.Trim(value, `"{}"`)

		args[key] = value
	}
	return command, args
}

func isValidAddCommandParameters(args map[string]string) (bool, *actions.AddCommandArgs) {
	allowedArgs := map[string]bool{
		SERVICE_ARG:  true,
		USERNAME_ARG: true,
		PASSWORD_ARG: true,
		NOTE_ARG:     true, // Optional
	}

	// Required parameters
	serviceSet := false
	usernameSet := false
	passwordSet := false

	for k, v := range args {
		// Reject any invalid parameter
		if !allowedArgs[k] {
			return false, nil
		}

		// Ensure required arguments are set (not empty)
		if k == SERVICE_ARG && v != "" {
			serviceSet = true
		}
		if k == USERNAME_ARG && v != "" {
			usernameSet = true
		}
		if k == PASSWORD_ARG && v != "" {
			passwordSet = true
		}
	}

	// SERVICE_ARG, USERNAME_ARG, and PASSWORD_ARG must be set and non-empty
	if serviceSet && usernameSet && passwordSet {
		args := &actions.AddCommandArgs{
			Service:  args[SERVICE_ARG],
			Username: args[USERNAME_ARG],
			Password: args[PASSWORD_ARG],
			Note:     helpers.NewOptional(args[NOTE_ARG]),
		}
		return true, args
	}

	return false, nil
}

func isValidEditCommandParameters(args map[string]string) (bool, *actions.EditCommandArgs) {

	allowedArgs := map[string]bool{
		SERVICE_ARG:  true,
		USERNAME_ARG: true,
		PASSWORD_ARG: true,
		NOTE_ARG:     true, // Optional
	}

	serviceSet := false
	usernameOrPasswordSet := false

	for k, v := range args {
		if !allowedArgs[k] {
			return false, nil
		}

		if k == SERVICE_ARG && v != "" {
			serviceSet = true
		}

		if (k == USERNAME_ARG || k == PASSWORD_ARG) && v != "" {
			usernameOrPasswordSet = true
		}
	}

	// SERVICE_ARG, USERNAME_ARG, or PASSWORD_ARG must be set and non-empty
	//return serviceSet && usernameOrPasswordSet
	if serviceSet && usernameOrPasswordSet {
		args := &actions.EditCommandArgs{
			Service:  args[SERVICE_ARG],
			Username: helpers.NewOptional(args[USERNAME_ARG]),
			Password: helpers.NewOptional(args[PASSWORD_ARG]),
			Note:     helpers.NewOptional(args[NOTE_ARG]),
		}
		return true, args
	}
	return false, nil
}

func isOnlyServiceParameterProvided(args map[string]string) (bool, string) {
	value, ok := args[SERVICE_ARG]
	return ok && value != "" && len(args) == 1, value
}

func isValidGenereteParameters(args map[string]string) (bool, *actions.GenerateCommandArgs) {

	allowedArgs := map[string]bool{
		LENGTH_ARG:     true,
		COMPLEXITY_ARG: true,
	}
	validComplexity := map[string]bool{
		COMPLEXITY_NUMBERS_ARG:   true,
		COMPLEXITY_SYMBOLS_ARG:   true,
		COMPLEXITY_UPPERCASE_ARG: true,
	}

	var lengthValue int
	var complexityFields []string

	for k, v := range args {
		if !allowedArgs[k] {
			return false, nil
		}

		if k == LENGTH_ARG && v != "" {
			value, err := strconv.Atoi(v)
			if err != nil {
				return false, nil
			}
			lengthValue = value
		}

		if k == COMPLEXITY_ARG && v != "" {
			fields := strings.Split(v, ",")
			for _, field := range fields {
				if !validComplexity[field] {
					return false, nil
				}
				complexityFields = append(complexityFields, field)
			}
		}
	}

	generateArgs := &actions.GenerateCommandArgs{
		Length: lengthValue,
	}

	for _, field := range complexityFields {
		switch field {
		case COMPLEXITY_NUMBERS_ARG:
			generateArgs.Numbers = helpers.NewOptional(field)
		case COMPLEXITY_UPPERCASE_ARG:
			generateArgs.UpperCase = helpers.NewOptional(field)
		case COMPLEXITY_SYMBOLS_ARG:
			generateArgs.Symbols = helpers.NewOptional(field)
		}
	}

	return true, generateArgs
}

func Validate(prompt string) (*actions.Action, error) {
	if prompt == QUIT_COMMAND {
		return actions.GetAction(actions.ACTION_QUIT, nil), nil
	} else if prompt == HELP_COMMAND {
		return actions.GetAction(actions.ACTION_HELP, nil), nil
	}

	command, parsedArgs := parseInput(prompt)

	if command == "" || parsedArgs == nil {
		return nil, errors.New(INVALID_COMMAND)
	}

	switch command {
	case ADD_COMMAND:
		isValid, args := isValidAddCommandParameters(parsedArgs)
		if !isValid {
			return nil, errors.New(INVALID_COMMAND)
		}
		return actions.GetAction(actions.ACTION_ADD, args), nil
	case EDIT_COMMAND:
		isValid, args := isValidEditCommandParameters(parsedArgs)
		if !isValid {
			return nil, errors.New(INVALID_COMMAND)
		}
		return actions.GetAction(actions.ACTION_EDIT, args), nil
	case DELETE_COMMAND:
		isValid, service := isOnlyServiceParameterProvided(parsedArgs)
		if !isValid {
			return nil, errors.New(INVALID_COMMAND)
		}
		return actions.GetAction(actions.ACTION_DELETE, service), nil
	case GET_COMMAND:
		isValid, service := isOnlyServiceParameterProvided(parsedArgs)
		if !isValid {
			return nil, errors.New(INVALID_COMMAND)
		}
		return actions.GetAction(actions.ACTION_GET, service), nil
	case GENERATE_COMMAND:
		isValid, args := isValidGenereteParameters(parsedArgs)
		if !isValid {
			return nil, errors.New(INVALID_COMMAND)
		}
		return actions.GetAction(actions.ACTION_GENERATE, args), nil
	default:
		return nil, errors.New(INVALID_COMMAND)
	}
}
