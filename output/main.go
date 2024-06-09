package main

import (
	"ascii-art/Fonctions"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		return
	} 
	output, str, banner := Fonctions.GetArgs(os.Args[1:])
	res_Final := Fonctions.AsciiArt(banner, str)
	if output != "" {
		err := os.WriteFile(output, []byte(res_Final), 0o644)
		if err != nil {
			return
		}
	} else {
		fmt.Print(res_Final)
	}
}
