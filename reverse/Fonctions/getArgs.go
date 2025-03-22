package Fonctions

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func GetArgs(args []string, clo, letters_to_be_colored, output *[]string, align, str, banner *string) { // clo slice of color
	var index int
	var clout bool
	var color string
	var reverse string
	for i := 0; i < len(args) && strings.Index(args[i], "--") == 0; i++ {
		index++
		*output, color, *align, clout = GetOption(args[i], color, *align, *output, &clout, &reverse)
		if reverse != "" {
			if len(args) != 1 {
				fmt.Println("\x1b[38;5;214mUsage: go run. [--reverse=filename.txt] [STRING]")
				os.Exit(0)
			}
			Reverse(reverse)
			os.Exit(0)
		}
		if clout && i < len(args)-1 && strings.Index(args[i+1], "--") != 0 {
			*letters_to_be_colored, *clo = CheckFlag(*clo, *letters_to_be_colored, args[i+1], color)
			i++
			index++
			clout = false
		} else if i+1 == len(args) {
			fmt.Println("\x1b[38;5;214mUsage: go run. [OPTION] [STRING]\n\nEX: go run. --color=<color> <letters to be colored> something\033[0m")
			os.Exit(0)
		}
	}
	*str, *banner = GetStrBanner(args, index)
}

func GetOption(arg, color, align string, output []string, clout *bool, reverse *string) ([]string, string, string, bool) {
	if strings.Index(arg, "--color=") == 0 && len(arg) > 8 {
		*clout = true
		return output, arg[8:], align, *clout
	} else if strings.Index(arg, "--output=") == 0 && len(arg) > 9 && arg[len(arg)-4:] == ".txt" {
		*clout = false
		output = append(output, arg[9:])
	} else if strings.Index(arg, "--align=") == 0 && len(arg) > 8 {
		align = arg[8:]
	} else if strings.Index(arg, "--reverse=") == 0 && len(arg) > 10 && arg[len(arg)-4:] == ".txt" {
		*reverse = arg[10:]
	} else {
		fmt.Println("\x1b[38;5;214mUsage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> something\033[0m")
		os.Exit(0)
	}
	return output, color, align, *clout
}

func CheckFlag(clo, ltc []string, arg, color string) ([]string, []string) {
	for i := 0; i < len(ltc); i++ {
		if i+1 < len(ltc) && ltc[i] == arg {
			ltc = slices.Delete(ltc, i, i+1)
			clo = slices.Delete(clo, i, i+1)
		} else if ltc[i] == arg {
			ltc = ltc[:i]
			clo = clo[:i]
		}
	}
	ltc = append(ltc, arg)
	clo = append(clo, color)
	return ltc, clo
}

func GetStrBanner(args []string, index int) (string, string) {
	if index == len(args) {
		return args[index-1], "standard.txt"
	} else if index == len(args)-2 && CheckIsBanner(args[index+1]) {
		return args[index], args[index+1]
	} else if index == len(args)-1 {
		if index > 0 && CheckIsBanner(args[index]) && !strings.Contains(args[index-1], "--") {
			return args[index-1], args[index]
		}
		return args[index], "standard.txt"
	} else {
		fmt.Println("\x1b[38;5;214mUsage: go run. [OPTION] [STRING]\n\nEX: go run. --color=<color> <letters to be colored> something\033[0m")
		os.Exit(0)
	}
	return "", ""
}
