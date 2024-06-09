package function

import (
	"fmt"
)

func Split_and_Print(split_Input, split_File []string) {
	for i := 0; i < len(split_Input); i++ {
		if split_Input[i] == "" {
			if i != 0 || !Empty(split_Input) {
				fmt.Println()
			}
		} else {
			PrintFinal(split_Input[i], split_File)
		}
	}
}

func Empty(str []string) bool {
	for i := 0; i < len(str); i++ {
		if str[i] != "" {
			return false
		}
	}
	return true
}

func PrintFinal(split_Input string, split_File []string) {
	output := make([][]string, 8)
	err := false
	for _, c := range split_Input {
		for i := 0; i < 8; i++ {
			if c < 32 || c > 126 {
				if !err {
					err = true
					fmt.Println("Check your input you have a character not printible!")
				}
			} else {
				output[i] = append(output[i], split_File[(int(c)-32)*9+i])
			}
		}
	}
	if len(output[0]) == 0 {
		return
	}
	for _, c := range output {
		for _, v := range c {
			fmt.Print(string(v))
		}
		fmt.Println()
	}
}
