package actions

import "fmt"

func ActionDelete(args *struct{ Service string }) {
	fmt.Println("** Process Delete Command **")
	fmt.Println("service:", args.Service)
}
