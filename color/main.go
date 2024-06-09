package main

import (
	"fmt"
	"os"

	"ascii-art/Fonctions"
)

func main() {
	if len(os.Args) < 2 {
		Fonctions.ErrExit()
	}
	var color, letters_to_be_colored, output, res_Color []string
	var str, banner string
	Fonctions.GetArgs(os.Args[1:], &color, &letters_to_be_colored, &output, &str, &banner)
	if output == nil && color != nil {
		res_Color = Fonctions.GetLetColor(color, letters_to_be_colored, str)
	}
	res_Final, nb := Fonctions.AsciiArt(banner, output, res_Color, str)
	if output != nil {
		for _, r := range output {
			bannerOutput, _ := os.ReadFile(r)
			if !Fonctions.CheakFormatBanner(bannerOutput, str, nb) {
				err := os.WriteFile(r, []byte(res_Final), 0o644)
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
