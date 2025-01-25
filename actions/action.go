package actions

import (
	"log"
)

const (
	ACTION_QUIT int = iota
	ACTION_ADD
	ACTION_HELP
)

type Args []string

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
