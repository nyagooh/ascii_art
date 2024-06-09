package main

import (
	"fmt"
	"os"
	"strings"

	art "ascii/functions"
)

func main() {
	// check for the correct passing of arguments
	// function for unprintable characters
	// handle two lines
	//handle test files
	//understand your code
	// do chunk with the same knowledge as printing strings menu
	arguments := len(os.Args)
	if arguments != 2 {
		fmt.Println("Wrong passing of arguments")
		return
	}
	if os.Args[1] == "" {
		return
	}
	if os.Args[1] == "\\n" {
		fmt.Println()
		return
	}
	input := os.Args[1] //hello
	for _, ch := range input {
		if !(ch >= 32 && ch <= 126) {
			fmt.Println("An unprinttable character was detected in your string")
		}
		if strings.Contains(os.Args[1], "\\t") {
			input = strings.ReplaceAll(os.Args[1], "\\t", "    ")
		}
	}

	lines := strings.Split(input, "\\n")
	result := art.ProcessLine(lines)
	result2 := art.ProcessString(result)
	art.PrintStrings(result2)
	fmt.Println()
}
