package main

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	originalStdout := os.Stdout
	file, _ := os.Create("testFile")
	os.Stdout = file
	os.Args = []string{".", "Hello"}
	main()
	os.Stdout = originalStdout
	output, _ := os.ReadFile("testFile")
	expectedOutput := " _    _          _   _          \n| |  | |        | | | |         \n| |__| |   ___  | | | |   ___   \n|  __  |  / _ \\ | | | |  / _ \\  \n| |  | | |  __/ | | | | | (_) | \n|_|  |_|  \\___| |_| |_|  \\___/  \n                                \n                                \n"
	if string(output) != expectedOutput {
		t.Errorf("Faild")
	}
	os.Remove("testFile")
}
