package commands

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func GetCommandPrompt() string {
	for {
		fmt.Print(">>> ")
		reader := bufio.NewReader(os.Stdin)
		enter, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("something went wrong")
		}
		fmt.Print("\n")
		enter = strings.TrimSpace(enter)

		if enter == "" {
			continue
		}
		return enter
	}
}

func Run() {
	for {
		prompt := GetCommandPrompt()
		action, err := Validate(prompt)
		if err != nil {
			fmt.Println(err)
			continue
		}
		action.Run()
	}
}
