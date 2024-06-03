package functions

import (
	"bufio"
	"fmt"
	"os"
)
/*The maps function, opens the required file,
scans line by line and maps the line number to the strings found in the various
lines in the opened file. For all these,
each character line contents are concatenated into a single string
the new string is returned as a line.
*/

func Maps(n int)string{
	file, err := os.Open("standard.txt")
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	var count int
	var line, str string
	for scanner.Scan() {
		line = scanner.Text()
		count++
		ascMap := make(map[int]string)
		ascMap[count] = line
		if count == n {
			str = ascMap[n]
			
		}
	}
	
	defer file.Close()
	return str
}

func ProcessString(slice []int)[]string{
	var results []string
	var str string
	for _, number := range slice {
		str = Maps(number)
		results = append(results, str)
	}
	return results
}
func PrintStrings(str []string){
	for _, ch := range str {
		fmt.Println(ch)
	}

}