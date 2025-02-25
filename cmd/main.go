package main

import (
	"fmt"
	"gopasskeeper/constants"
	"gopasskeeper/interpreter"
)

func main() {
	fmt.Println(constants.AppBannerMsg)
	fmt.Println()
	interpreter.Run()
}
