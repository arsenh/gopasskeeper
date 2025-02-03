package actions

import (
	"fmt"
	"gopasskeeper/helpers"
)

type AddCommandArgs struct {
	Service  string
	Username string
	Password string
	Note     helpers.Optional[string]
}

func ActionAdd(args *AddCommandArgs) {

	fmt.Println("** Process Add Command **")
	fmt.Println("service:", args.Service)
	fmt.Println("username:", args.Username)
	fmt.Println("password:", args.Password)

	if args.Note.HasValue() {
		fmt.Println("note:", args.Note.MustGet())
	} else {
		fmt.Println("note parameter is not provided")
	}
}
