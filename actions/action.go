package actions

import (
	"gopasskeeper/constants"
	"log"
)

const (
	ACTION_QUIT int = iota
	ACTION_HELP
	ACTION_ADD
	ACTION_EDIT
	ACTION_DELETE
	ACTION_GET
	ACTION_GENERATE
)

type Action struct {
	Args any
	Fn   func(any)
}

func (a *Action) Run() {
	if a.Fn != nil {
		a.Fn(a.Args)
	} else {
		log.Fatal(constants.ErrNoFunctionAssigned)
	}
}

func NewAction[T any](args T, fn func(T)) *Action {
	return &Action{
		Args: args,
		Fn:   func(v any) { fn(v.(T)) },
	}
}

func GetAction(actionNumber int, args any) *Action {
	switch actionNumber {
	case ACTION_QUIT:
		var dummy any = struct{}{}
		return NewAction(dummy, ActionQuit)
	case ACTION_HELP:
		var dummy any = struct{}{}
		return NewAction(dummy, ActionHelp)
	case ACTION_ADD:
		addargs, ok := args.(*AddCommandArgs)
		if !ok {
			log.Fatal(constants.ErrInvalidOperationAction)
		}
		return NewAction(addargs, ActionAdd)
	case ACTION_EDIT:
		editargs, ok := args.(*EditCommandArgs)
		if !ok {
			log.Fatal(constants.ErrInvalidOperationAction)
		}
		return NewAction(editargs, ActionEdit)
	case ACTION_GET:
		service, ok := args.(string)
		if !ok {
			log.Fatal(constants.ErrInvalidOperationAction)
		}
		return NewAction(service, ActionGet)
	case ACTION_DELETE:
		service, ok := args.(string)
		if !ok {
			log.Fatal(constants.ErrInvalidOperationAction)
		}
		return NewAction(service, ActionDelete)
	case ACTION_GENERATE:
		getargs, ok := args.(*GenerateCommandArgs)
		if !ok {
			log.Fatal(constants.ErrInvalidOperationAction)
		}
		return NewAction(getargs, ActionGenerate)

	default:
		log.Fatal(constants.ErrInvalidOperationAction)
		return nil
	}
}
