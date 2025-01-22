package main

import (
	"fmt"
	"gopasskeeper/commands"
)

func main() {

	fmt.Println("*** GoPassKeeper Offline Password Manager ***")
	commands.Run()
}
