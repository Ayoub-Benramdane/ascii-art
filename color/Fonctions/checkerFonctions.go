package Fonctions

import (
	"fmt"
	"os"
	"strings"
)

// Capter les lettres à colorier

func CheckLetters(letters_to_be_colored []string, arg []string) ([]string, []string) { // arg = color(red)
	for i := len(letters_to_be_colored) - 1; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if letters_to_be_colored[i] == letters_to_be_colored[j] {
				arg[j] = arg[i]
			}
		}
	}
	return letters_to_be_colored, arg
}

func CheckIsBanner(arg string) bool {
	if len(arg) < 4 || arg[len(arg)-4:] != ".txt" {
		arg += ".txt"
	}
	_, err := os.ReadFile(arg)
	return err == nil
}

func CheakFormatBanner(Banner []byte, str string, nb int) bool {
	var maxAscii rune
	for _, c := range str {
		if maxAscii == 0 && c >= 32 && c <= 126 {
			maxAscii = c
		} else if c > maxAscii && c >= 32 && c <= 126 {
			maxAscii = c
		}
	}
	Len_File := int(maxAscii-31) * (nb + 1)
	return len(strings.Split(string(Banner), "\n")) > Len_File && len(strings.Split(string(Banner), "\n")) <= 95*(nb+1)+1
}

func IsHexa(str string) bool {
	listeHexa := "0123456789abcdefABCDEF"
	count := 0
	for _, c := range listeHexa {
		for _, v := range str {
			if c == v {
				count++
			}
		}
	}
	return count == 6
}

func CheckAnsi(str string) bool {
	listeAnsi := []string{"", "black", "red", "green", "yellow", "blue", "purple", "cyan", "white", "pink", "orange", "gray", "rainbow"}
	for _, c := range listeAnsi {
		if c == strings.ToLower(str) {
			return true
		}
	}
	return false
}

func Empty(Lines []string) bool {
	for _, str := range Lines {
		if str != "" {
			return false
		}
	}
	return true
}

func ErrExit() {
	fmt.Println("\x1b[38;5;214mUsage: go run. [OPTION] [STRING]\n\nEX: go run. --color=<color> <letters to be colored> something\033[0m")
	os.Exit(0)
}
