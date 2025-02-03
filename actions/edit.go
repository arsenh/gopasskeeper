package actions

import (
	"fmt"
	"gopasskeeper/helpers"
)

type EditCommandArgs struct {
	Service  string
	Username helpers.Optional[string]
	Password helpers.Optional[string]
	Note     helpers.Optional[string]
}

func ActionEdit(args *EditCommandArgs) {
	fmt.Println("** Process Edit Command **")
	fmt.Println("service:", args.Service)

	if args.Username.HasValue() {
		fmt.Println("username:", args.Username.MustGet())
	} else {
		fmt.Println("username parameter is not provided")
	}

	if args.Password.HasValue() {
		fmt.Println("password:", args.Password.MustGet())
	} else {
		fmt.Println("password parameter is not provided")
	}

	if args.Note.HasValue() {
		fmt.Println("note:", args.Note.MustGet())
	} else {
		fmt.Println("note parameter is not provided")
	}

}
