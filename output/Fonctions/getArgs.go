package Fonctions

import (
	"fmt"
	"os"
	"strings"
)

func GetArgs(args []string) (string, string, string) {
	var output, str, banner string
	index := 0
	for i := 0;i < len(args) && strings.Index(args[i], "--") == 0; i++ {
		index++
		output = GetOption(args[i], "--output=")
	}
	if len(args) == index+2 {
		str = args[index]
		banner = args[index+1]
	} else if len(args) == index+1 {
		str = args[index]
		banner = "standard.txt"
	} else {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		os.Exit(0)
	}
	return output, str, banner
}

func GetOption(arg, output string) string {
	if !(strings.Index(arg, output) == 0 && len(arg) > 9 && arg[len(arg)-4:] == ".txt") {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		os.Exit(0)
	}
	return arg[9:]
}
