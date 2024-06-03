package main

import (
	"fmt"
	"os"
	"strings"

	art "ascii/functions"
)

func main() {
	// check for the correct passing of arguments
	arguments := len(os.Args)
	if arguments != 2 {
		fmt.Println("Wrong passing of arguments")
		return
	}

	// check if the passed string is an empty string
	if os.Args[1] == "" {
		return
	}

	// check if the string passed is a new line character
	if os.Args[1] == "\\n" {
		fmt.Println()
		return
	}
	input := os.Args[1]
	for _, ch := range input {
		if !(ch >= 32 && ch <= 126) {
			fmt.Println("An unprinttable character was detected in your string")
		}
		if strings.Contains(os.Args[1], "\\t") {
			input = strings.ReplaceAll(os.Args[1], "\\t", "    ")
		}
		lines := strings.Split(input, "\\n")
		result := art.ProcessLine(lines)
		result2 := art.ProcessString(result)
		art.PrintStrings(result2)
	}
	fmt.Println()
}
