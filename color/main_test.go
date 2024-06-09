package main

import (
	"os"
	"os/exec"
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

func TestError1(t *testing.T) {
	cmd := exec.Command("go", "run", ".", "Hello There!", "shadow")
	out, _ := cmd.Output()
	expectedOutput, _ := os.ReadFile("TestFile.txt")
	if string(out) != string(expectedOutput) {
		t.Errorf("Faild")
	}
}

func TestError2(t *testing.T) {
	cmd := exec.Command("go", "run", ".", "Hello There!", "thinkertoy.txt")
	out, _ := cmd.Output()
	expectedOutput, _ := os.ReadFile("TestFile1.txt")
	if string(out) != string(expectedOutput) {
		t.Errorf("Faild")
	}
}
