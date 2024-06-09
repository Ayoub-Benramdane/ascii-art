package main

import (
	"fmt"
	"os"
	"ascii-art/Fonctions"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
		return
	}
	Banner := Fonctions.SelectBanner(os.Args)
	resFinal := Fonctions.FichierTxt(Banner, os.Args[1])
	fmt.Print(resFinal)
}
