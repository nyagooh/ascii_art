package functions

import "fmt"

// liner function calculates mathematically the line number of each rune contained in the input string.
func FndLine(r rune) map[int]int {
	lineNumbers := make(map[int]int)
	endLine := 9
	for i := 2; i <= endLine; i++ {
		lineNumbers[i] = int(((r - 32) * 9)) + i
	}

	return lineNumbers
}

func ProcessLine(line []string) map[string][]int {
	result := make(map[string][]int)

	for _, values := range line {
		for _, char := range values {
		lineNumbers := FndLine(char)
		for i := 2; i <= 9; i++ {
			key := fmt.Sprintf("Line%d", i)
			result[key] = append(result[key], lineNumbers[i])
		}
	}
	}

	return result
}
