package actions

import (
	"fmt"
	"gopasskeeper/constants"
)

func ActionHelp(args any) {
	fmt.Println(constants.HelpMsg)
	fmt.Println()
}
