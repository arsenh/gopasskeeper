package main

import (
	"fmt"
	"gopasskeeper/commands"
	"gopasskeeper/constants"
)

func main() {
	fmt.Println(constants.AppBannerMsg)
	fmt.Println()
	commands.Run()
}
