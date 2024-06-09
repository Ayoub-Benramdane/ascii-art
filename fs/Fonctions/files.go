package Fonctions

import (
	"fmt"
	"os"
	"strings"
)

func SelectBanner(str []string) string {
	Banner := ""
	if len(str) == 2 {
		Banner = "standard.txt"
	} else {
		Banner = str[2]
	}
	return Banner
}

func FichierTxt(Banner, str string) string {
	if len(Banner) < 4 {
		Banner += ".txt"
	} else if Banner[len(Banner)-4:] != ".txt" {
		Banner += ".txt"
	}
	fichier, err := os.ReadFile(Banner)
	if err != nil || !CheakBanner(fichier, str) {
		if err != nil {
			fmt.Println(err)
		}
		return ""
	}
	return Principale(string(fichier), os.Args[1])
}

func CheakBanner(Banner []byte, str string) bool {
	var maxAscii rune
	for _, c := range str {
		if maxAscii == 0 && c >= 32 && c <= 126 {
			maxAscii = c
		} else if c > maxAscii && c >= 32 && c <= 126 {
			maxAscii = c
		}
	}
	Len_File := int((maxAscii-31) * 9)
	if len(strings.Split(string(Banner), "\n")) < Len_File {
		fmt.Println("Banner format not valid")
		return false
	}
	return true
}
