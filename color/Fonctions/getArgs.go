package Fonctions

import (
	"slices"
	"strings"
)

func GetArgs(args []string, clo, letters_to_be_colored, output *[]string, str, banner *string) { // clo slice of color
	var index int
	var clout bool
	var color string
	for i := 0; i < len(args) && strings.Index(args[i], "--") == 0; i++ {
		index++
		*output, color, clout = GetOption(args[i], color, *output)
		if clout && i < len(args)-1 && strings.Index(args[i+1], "--") != 0 { // lettres_to_be_colored != nil
			*letters_to_be_colored, *clo = CheckFlag(*clo, *letters_to_be_colored, args[i+1], color)
			i++
			index++
		} else if i+1 == len(args) {
			ErrExit()
		}
	}
	*str, *banner = GetStrBanner(args, index)
}

func GetOption(arg, color string, output []string) ([]string, string, bool) {
	clout := true
	if strings.Index(arg, "--color=") == 0 && len(arg) > 8 {
		return output, arg[8:], clout
	} else if strings.Index(arg, "--output=") == 0 && len(arg) > 9 && arg[len(arg)-4:] == ".txt" {
		clout = false
		output = append(output, arg[9:])
	} else {
		ErrExit()
	}
	return output, color, clout
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

// Test if the banner is specified or not

func GetStrBanner(args []string, index int) (string, string) {
	if index == len(args) {
		return args[index-1], "standard.txt"
	} else if index == len(args)-2 && CheckIsBanner(args[index+1]) {
		return args[index], args[index+1]
	} else if index == len(args)-1 {
		if index > 0 && CheckIsBanner(args[index]) && strings.Index(args[index-1], "--") != 0 {
			return args[index-1], args[index]
		}
		return args[index], "standard.txt"
	} else {
		ErrExit()
	}
	return "", ""
}
