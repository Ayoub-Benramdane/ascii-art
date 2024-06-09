package Fonctions

import (
	"fmt"
	"os"
	"strings"
)

func AsciiArt(Banner string, output, res_Color []string, str string) (string,int) {
	if len(Banner) < 4 || Banner[len(Banner)-4:] != ".txt" {
		Banner += ".txt"
	}
	fichier, err := os.ReadFile(Banner)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	res_Color = GetColor(res_Color)
	return PrintAsciiArt(string(fichier), output, res_Color, str)
}

func PrintAsciiArt(banner string, output, res_Color []string, arg string) (string, int) {
	var resultat_Final string
	banner = strings.ReplaceAll(banner, "\r", "")
	tableau, nb := AsciiTable(strings.Split(banner[1:], "\n\n"))
	if !CheakFormatBanner([]byte(banner), arg, nb) {
		fmt.Println("\x1b[38;5;190mBanner format not valid!\033[0m")
		os.Exit(0)
	}
	Lines := strings.Split(arg, "\\n")
	if len(Lines) > 0 && Empty(Lines) {
		Lines = Lines[1:]
	}
	for i, str := range Lines {
		if i != 0 && len(res_Color) != 0 {
			res_Color = res_Color[2:]
		}
		if str == "" {
			resultat_Final += "\n"
			continue
		} else {
			resultat_Final += AddLine(tableau, str, output, res_Color, nb)
		}

		if len(res_Color) != 0 {
			res_Color = res_Color[len(str):]
		}
	}
	return resultat_Final, nb
}

func AsciiTable(split_File []string) ([][]string, int) {
	var tableau [][]string
	for i := 0; i < len(split_File); i++ {
		tableau = append(tableau, strings.Split(split_File[i], "\n"))
	}
	for i := 0; i < len(tableau); i++ {
		if i == len(tableau) - 1 &&  len(tableau[i]) == len(tableau[0]) + 1 {
			continue
		}else if len(tableau[i]) != len(tableau[0]) {
			fmt.Println("\x1b[38;5;190mBanner format not valid!kk\033[0m")
			os.Exit(0)
		}
	}
	return tableau, len(tableau[0])
}

func AddLine(tableau [][]string, str string, output, res_Color []string, nb int) string {
	var resultat_Final string
	var err bool
	for k := 0; k < nb*len(str); k++ {
		Line := k / len(str) % nb
		Character := k % len(str)
		if int(str[Character]) < 32 || int(str[Character]) > 126 {
			if !err {
				err = true
				fmt.Println("\x1b[38;5;190mCheck your input you have a character not printable!\033[0m")
			}
		} else {
			if len(res_Color) != 0 && output == nil {
				resultat_Final = PrintColor(resultat_Final, str, res_Color, tableau, k, Character, Line)
			} else {
				resultat_Final += (tableau[int(str[Character])-32][Line])
			}
		}
		if (k+1)%len(str) == 0 && resultat_Final != "" {
			resultat_Final += "\n"
		}
	}
	return resultat_Final
}

func PrintColor(resultat_Final, str string, res_Color []string, tableau [][]string, k, Character, Line int) string {
	if res_Color[Character] == "rainbow" {
		Rin := []string{"\033[36m", "\033[32m", "\033[33m", "\033[34m", "\033[35m", "\033[31m", "\033[35m", "\033[33m"}
		resultat_Final += Rin[Line] + (tableau[int(str[Character])-32][Line]) + "\033[0m"
	} else {
		resultat_Final += res_Color[Character] + (tableau[int(str[Character])-32][Line]) + "\033[0m"
	}
	return resultat_Final
}
