package Fonctions

import (
	"fmt"
	"os"
	"strings"
)

func AsciiArt(Banner, str string) string {
	if len(Banner) < 4 || len(Banner) >= 4 && Banner[len(Banner)-4:] != ".txt" {
		Banner += ".txt"
	}
	fichier, err := os.ReadFile(Banner)
	if err != nil || !CheakBanner(fichier, str) {
		if err != nil {
			fmt.Println(err)
		}
		os.Exit(0)
	}
	return PrintAsciiArt(string(fichier), str, Banner)
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
		fmt.Println("Banner format not valid!")
		return false
	}
	return true
}
