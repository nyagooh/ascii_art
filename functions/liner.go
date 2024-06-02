package functions

// liner function calculates mathematically the line number of each rune contained in the input string.
func FndLine(r rune) int {
	var lineNumbers int
	endLine := 9
	for i := 2; i <= endLine; i++ {
		lineNumbers = int(((r - 32) * 9)) + i
	}
	return lineNumbers
}

func ProcessLine(line []string) []int {
	var result[]int
	for _, values := range line {
		for _, char := range values {
		lineNumbers := FndLine(char)
		for i := 2; i <= 9; i++ {
			result = append(result, lineNumbers)
		}
	}
	}
return result
}
