package actions

import (
	"log"
)

const (
	ACTION_QUIT int = iota
	ACTION_ADD
	ACTION_HELP
)

// command arguments
const (
	SERVICE_ARG  = "service"
	USERNAME_ARG = "username"
	PASSWORD_ARG = "password"
	NOTES_ARG    = "notes"
)

type Args map[string]string

type ActionFunction func(args Args)

var actionRegistry = map[int]ActionFunction{
	ACTION_QUIT: ActionQuit,
	ACTION_ADD:  ActionAdd,
	ACTION_HELP: ActionHelp,
}

type Action struct {
	function ActionFunction
	args     Args
}

func (a *Action) Run() {
	a.function(a.args)
}

func GetAction(actionNumber int, args Args) *Action {

	f, ok := actionRegistry[actionNumber]

	if !ok {
		log.Fatal("invalid operation (action) is selected")
	}

	return &Action{
		function: f,
		args:     args,
	}
}
