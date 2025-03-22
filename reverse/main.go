package main

import (
	"fmt"
	"os"

	"ascii-art/Fonctions"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("\x1b[31mUsage: go run . [OPTION] [STRING] [BANNER]\n\nExample1: go run . --color=<color> <letters to be colored> something\n\nExample2: go run . --output=<fileName.txt> something standard\033[0m")
		return
	}
	var color, letters_to_be_colored, output, res_Color []string
	var str, banner, align string
	var countLetter []int
	Fonctions.GetArgs(os.Args[1:], &color, &letters_to_be_colored, &output, &align, &str, &banner)
	if output == nil && color != nil {
		res_Color = Fonctions.GetLetColor(color, letters_to_be_colored, str)
	}
	res_Final, nb := Fonctions.AsciiArt(banner, output, res_Color, str, align, &countLetter)
	if output == nil && align != "" {
		res_Final = Fonctions.AlignText(res_Final, align, &countLetter, Fonctions.GetTerminalSize(), nb)
	}
	if output != nil {
		for _, r := range output {
			bannerOutput, _ := os.ReadFile(r)
			if !Fonctions.CheakFormatBanner(bannerOutput, str, nb) {
				err := os.WriteFile(r, []byte(res_Final), 0644)
				if err != nil {
					return
				}
			} else {
				fmt.Println("\033[31mimpossible to modify this banner\033[0m")
			}
		}
	} else {
		fmt.Print(res_Final)
	}
}
