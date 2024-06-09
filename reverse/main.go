package main

import (
	"flag"
	"fmt"
	function "function/Functions"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . [OPTION]\n\nEX: go run . --reverse=<fileName>")
		return
	}
	file, _ := os.ReadFile("standard.txt")
	reverse_Flag := flag.String("reverse", "", "Reverse file")
	flag.Parse()
	res_Final := function.Split_and_Print(os.Args[1], file)
	fmt.Println(res_Final)
}
