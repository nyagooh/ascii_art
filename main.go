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
	printable := art.NonPrintable(input)
	lines := strings.Split(printable, "\\n")
	result := art.ProcessLine(lines)
	result2 := art.ProcessString(result)
	art.PrintStrings(result2)
	fmt.Println()
}
