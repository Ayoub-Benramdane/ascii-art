package Fonctions

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func GetTerminalSize() int {
	width := 0
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	size, _ := cmd.Output()
	width, _ = strconv.Atoi(strings.Trim(strings.Split(string(size), " ")[1], "\n"))
	return width
}

func AlignText(text, align string, countLetter *[]int, width int) string {
	lines := strings.Split(text, "\n")
	var res_Final strings.Builder
	for i := 0; i < len(lines)-1; i++ {
		line := lines[i]
		padding := width + (*countLetter)[0] - len(line)
		fmt.Println(countLetter)
		if len(line) > width+(*countLetter)[0] {
			align = "left"
		}
		switch align {
		case "center":
			res_Final.WriteString(strings.Repeat(" ", padding/2))
		case "right":
			res_Final.WriteString(strings.Repeat(" ", padding))
		case "justify":
			words := strings.Split(line, "      $")
			if len(words) > 1 {
				total_Spaces := width - len(line) + 7*(len(words)-1) + (*countLetter)[0]
				if len(line) > width+7*(len(words)-1)+(*countLetter)[0] {
					total_Spaces = -total_Spaces
				}
				space_Between_Words := total_Spaces / (len(words) - 1)
				for i, word := range words {
					if i > 0 {
						res_Final.WriteString(strings.Repeat(" ", space_Between_Words))
					}
					res_Final.WriteString(word)
				}
				res_Final.WriteString("\n")
				continue
			}
		}
		res_Final.WriteString(line)
		if i < len(lines)-1 {
			res_Final.WriteString("\n")
		}
	}
	lines = nil
	return res_Final.String()
}
