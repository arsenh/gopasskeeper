package actions

import "fmt"

func ActionAdd(args Args) {
	//TODO: delete dump of args.
	fmt.Println("Run Add Command with Args:")
	fmt.Println("service: ", args[SERVICE_ARG])
	fmt.Println("username: ", args[USERNAME_ARG])
	fmt.Println("password: ", args[PASSWORD_ARG])

	notes, ok := args[NOTES_ARG]
	if ok {
		fmt.Println("notes: ", notes)
	} else {
		fmt.Println("notes argument is not set")
	}
}
