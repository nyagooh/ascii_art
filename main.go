package main

import (
	"fmt"
	"os"
	"strings"

	art "ascii/functions"
)

func main() {
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
	input := os.Args[1]
	art.NonPrintable(input)
	lines := strings.Split(input, "\\n")
	for _, line := range lines {
		if line != "" {
			result := art.ProcessLine(line)
			result2 := art.ProcessString(result)
			art.PrintStrings(result2)
		} else {
			fmt.Println()
		}
	}
}
