package actions

import "fmt"

func ActionGet(args *struct{ Service string }) {
	fmt.Println("** Process Get Command **")
	fmt.Println("service:", args.Service)
}
