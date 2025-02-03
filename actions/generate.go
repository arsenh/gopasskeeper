package actions

import (
	"fmt"
	"gopasskeeper/helpers"
)

type GenerateCommandArgs struct {
	Length    int
	UpperCase helpers.Optional[string]
	Numbers   helpers.Optional[string]
	Symbols   helpers.Optional[string]
}

func ActionGenerate(args *GenerateCommandArgs) {
	fmt.Println("Process Generate Password.")
	fmt.Println("length:", args.Length)
	if args.UpperCase.HasValue() {
		fmt.Println("uppercase:", args.UpperCase.MustGet())
	} else {
		fmt.Println("uppercase parameter is not provided")
	}

	if args.Numbers.HasValue() {
		fmt.Println("numbers:", args.Numbers.MustGet())
	} else {
		fmt.Println("numbers parameter is not provided")
	}

	if args.Symbols.HasValue() {
		fmt.Println("symbols:", args.Symbols.MustGet())
	} else {
		fmt.Println("symbols parameter is not provided")
	}
}
