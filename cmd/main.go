package main

import (
	"fmt"
	. "gopasskeeper/colors"
	"gopasskeeper/commands"
)

func printBanner() {
	fmt.Println(Bold + Blue + "***********************************************************" + Reset)
	fmt.Println(Bold + Green + "🚀  Welcome to GoPassKeeper - Offline Password Manager  🔒" + Reset)
	fmt.Println(Bold + Blue + "***********************************************************" + Reset)
	fmt.Println(Yellow + "🔑 Securely store and manage your passwords offline!" + Reset)
	fmt.Println()
}

func main() {
	printBanner()
	commands.Run()
}
