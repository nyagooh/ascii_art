package functions

// import "fmt"

// liner function calculates mathematically the line number of each rune contained in the input string.
func FndLine(r rune) []int {
	lineNumbers := make([]int, 8)
	for i := 0; i < 8; i++ {
		lineNumbers[i] = (int(r-32) * 9) + i + 2
	}
	// fmt.Println(lineNumbers)

	return lineNumbers
}

func ProcessLine(line []string) []int {
	var result []int
	for _, values := range line {
		for _, char := range values {
			lineNumbers := FndLine(char)
			result = append(result, lineNumbers...)

		}
	}
	// fmt.Println(result)
	return result
}
