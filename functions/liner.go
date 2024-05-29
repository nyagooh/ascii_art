package functions

//liner function calculates mathematically the line number of each rune contained in the input string.
func FndLine(r rune) map[int]int {
	lineNumbers := make(map[int]int)
	endLine := 9
	for i := 2; i <= endLine; i++ {
		lineNumbers[i] = int(((r - 32) * 9)) + i
	}

	return lineNumbers
}

func CalculateLine(){
	
}