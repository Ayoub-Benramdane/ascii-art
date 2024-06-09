package main

import (
	"fmt"
	function "function/Functions"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	file, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(strings.Split(string(file), "\n")) != 856 {
		fmt.Println("Banner format not valid")
		return
	}
	split_File := strings.Split(string(file[1:]), "\n")
	split_Input := strings.Split(os.Args[1], "\\n")
	function.Split_and_Print(split_Input, split_File)
}
