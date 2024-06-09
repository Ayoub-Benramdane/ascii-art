package Fonctions

import (
	"fmt"
	"strings"
)

func PrintAsciiArt(banner, arg, BannerName string) string {
	var resultat string
	if BannerName == "thinkertoy.txt" {
		banner = strings.ReplaceAll(banner, "\r", "")
	}
	split := strings.Split(banner[1:], "\n\n")
	tableau := AsciiTable(split)
	Lines := strings.Split(arg, "\\n")
	if Empty(Lines) {
		Lines = Lines[1:]
	}
	for _, str := range Lines {
		if str == "" {
			resultat += "\n"
		} else {
			resultat += AddLine(tableau, str)
		}
	}
	return resultat
}

func AsciiTable(split_File []string) [][]string {
	var tableau [][]string
	for i := 0; i < len(split_File); i++ {
		tableau = append(tableau, strings.Split(split_File[i], "\n"))
	}
	return tableau
}

func Empty(Lines []string) bool {
	for _, str := range Lines {
		if str != "" {
			return false
		}
	}
	return true
}

func AddLine(tableau [][]string, str string) string {
	resultat := ""
	err := false
	for k := 0; k < 8*len(str); k++ {
		if int(str[k%len(str)]) < 32 || int(str[k%len(str)]) > 126 {
			if !err {
				err = true
				fmt.Println("Check your input you have a character not printible!")
			}
		} else {
			resultat += (tableau[int(str[k%len(str)])-32][k/len(str)%8])
		}
		if (k+1)%len(str) == 0 {
			resultat += "\n"
		}
	}
	return resultat
}
