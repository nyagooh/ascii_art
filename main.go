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

		// check if the string passed contains only the printable characters
		if !(ch >= 32 && ch <= 126) {
			fmt.Println("An unprinttable character was detected in your string")
		} else {
			if strings.Contains(os.Args[1], "\\t") {
				input = strings.ReplaceAll(os.Args[1], "\\t", "    ")
			}
			var lines []string
			lines = strings.Split(input, "\\n")

			for _, line := range lines {
				if line != "" {

					/*calaculate the line of each character in your input string
					  and store the numbers of each character to its corresponding line number from the file
					*/
					var nums, a, b, c, d, e, f, g []int
					for _, char := range line {
						nums = append(nums, art.FindHeadLine(char))
					}
					for _, char := range line {
						a = append(a, art.Line2(char))
					}
					for _, char := range line {
						b = append(b, art.Line3(char))
					}
					for _, char := range line {
						c = append(c, art.Line4(char))
					}
					for _, char := range line {
						d = append(d, art.Line5(char))
					}
					for _, char := range line {
						e = append(e, art.Line6(char))
					}
					for _, char := range line {
						f = append(f, art.Line7(char))
					}
					for _, char := range line {
						g = append(g, art.Line8(char))
					}

					/*
						concatenate the characters of every line to a string
					*/
					var str, s, stri, strn, trn, tr, sring, sr string
					for _, num := range nums {
						str += art.Maps(num)
					}
					for _, ray := range a {
						s += art.Maps(ray)
					}
					for _, ry := range b {
						stri += art.Maps(ry)
					}
					for _, rata := range c {
						strn += art.Maps(rata)
					}
					for _, dig := range d {
						trn += art.Maps(dig)
					}
					for _, digi := range e {
						tr += art.Maps(digi)
					}
					for _, digit := range f {
						sring += art.Maps(digit)
					}
					for _, kala := range g {
						sr += art.Maps(kala)
					}
					fmt.Println(str)
					fmt.Println(s)
					fmt.Println(stri)
					fmt.Println(strn)
					fmt.Println(trn)
					fmt.Println(tr)
					fmt.Println(sring)
					fmt.Println(sr)
				} else if line == "\\n" {
					fmt.Println()
				} else {
					fmt.Println()
				}
			}

		}
		return
	}
}
