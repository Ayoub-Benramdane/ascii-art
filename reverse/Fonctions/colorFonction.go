package Fonctions

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetColor(color []string) []string {
	colorTable := []string{}
	for _, c := range color {
		if len(c) > 11 && len(c) < 19 && c[:4] == "rgb(" && c[len(c)-1] == ')' {
			colorTable = append(colorTable, RgbHsl(c, "rgb"))
		} else if len(c) > 11 && len(c) < 22 && c[:4] == "hsl(" && c[len(c)-1] == ')' {
			colorTable = append(colorTable, RgbHsl(c, "hsl"))
		} else if len(c) == 7 && c[0] == '#' {
			colorTable = append(colorTable, HexColor(c))
		} else if CheckAnsi(c) {
			colorTable = append(colorTable, AnsiColor(strings.ToLower(c)))
		} else {
			ErrColor(c)
		}
	}
	return colorTable
}

func RgbHsl(str, chek string) string {
	var rgb [3]int
	var RGB [3]string
	var err error
	k := 0
	var st string
	for i := 4; i < len(str); i++ {
		if (str[i] == ',' && str[i+1] == ' ' && k < 3) || str[i] == ')' {
			if chek == "hsl" && st != "0" {
				if st[len(st)-1] != '%' {
					ErrColor(str)
				}
				st = st[:len(st)-1]
			}
			rgb[k], err = strconv.Atoi(st)
			if rgb[k] < 0 || rgb[k] > 255 || err != nil{
				ErrColor(str)
			}
			RGB[k] = st
			k++
			i++
			st = ""
			continue
		}
		if chek == "hsl" && len(st) > 4 || chek == "rgb" && len(st) > 3 {
			ErrColor(str)
		}
		st += string(str[i])
	}
	if st != "" || k != 3 {
		ErrColor(str)
	}
	return "\033[38;2;" + RGB[0] + ";" + RGB[1] + ";" + RGB[2] + "m"
}

func HexColor(str string) string {
	if !IsHexa(str) {
		ErrColor(str)
	}
	var rgb [3]int64
	var err error
	j := 0
	for i := 1; i <= 5; i += 2 {
		rgb[j], err = strconv.ParseInt(str[i:i+2], 16, 32)
		if err != nil {
			ErrColor(str)
		}
		j++
	}
	rtn := fmt.Sprintf("\033[38;2;%d;%d;%dm", rgb[0], rgb[1], rgb[2])
	return rtn
}

func AnsiColor(color string) string {
	colorCodes := map[string]string{
		"black":  "\033[30m",
		"red":    "\033[31m",
		"green":  "\033[32m",
		"yellow": "\033[33m",
		"blue":   "\033[34m",
		"purple": "\033[35m",
		"cyan":   "\033[36m",
		"white":  "\033[37m",
		"pink":   "\033[38;5;205m",
		"orange": "\033[38;5;208m",
		"gray":   "\033[38;5;240m",
		"rainbow": "rainbow",
	}
	code, exists := colorCodes[color]
	if exists {
		return code
	}
	return ""
}

// If Color not valid

func ErrColor(str string) {
	fmt.Printf("Color \033[31m%s\033[0m not defined!", str)
	fmt.Println()
	os.Exit(0)
}
