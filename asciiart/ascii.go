package asciiart

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func AsciiTable(input, banner string) (string, error) {
	str := []rune(input)
	lnum := []int{}
	data, err := ioutil.ReadFile(banner)
	if err != nil {
		return "", fmt.Errorf("error reading file: %v", err)
	}

	// Calculate the line numbers for each character
	for i := 0; i < len(str); i++ {
		fline := ((int(str[i]) - 32) * 9) + 2
		lnum = append(lnum, fline)
	}

	for a := 0; a < len(lnum); a++ {
		if a+1 < len(lnum) && lnum[a] == 542 && lnum[a+1] == 704 {
			lnum[a] = 0
			lnum[a+1] = 0
		}
	}
	// Split the lnum array when there are two consecutive zeros
	var parts [][]int
	var currentPart []int
	for _, num := range lnum {
		if num == 0 && len(currentPart) > 0 && currentPart[len(currentPart)-1] == 0 {
			parts = append(parts, currentPart)
			currentPart = []int{}
		} else {
			currentPart = append(currentPart, num)
		}
	}
	if len(currentPart) > 0 {
		parts = append(parts, currentPart)
	}

	var result strings.Builder
	// Print each part and send it to the Table function
	for _, part := range parts {
		if EqualToZero(part) {
			result.WriteString("\n")
		} else {
			result.WriteString(Table(part, data))
		}
	}
	if checkLastElement(parts) {
		result.WriteString("\n")
	}
	return result.String(), nil
}

func Table(lnum []int, data []byte) string {
	var result strings.Builder
	// Convert file content to string
	text := string(data)

	// Split the content into lines
	lines := strings.Split(text, "\n")

	// Print the lines corresponding to the line numbers
	for k := 0; k < 8; k++ {
		for _, lineNum := range lnum {
			if lineNum != 0 && lineNum-1 < len(lines) {
				result.WriteString(lines[lineNum-1])
			} else {
				break
			}
		}
		result.WriteString("\n")

		// Increment the line numbers
		for j := 0; j < len(lnum); j++ {
			if lnum[j] != 0 {
				lnum[j]++
			}
		}
	}
	return result.String()
}

// check if the part is equal to [0], add new
func EqualToZero(arr []int) bool {
	if len(arr) != 1 {
		return false
	}
	return arr[0] == 0
}

// check if the last element of the last part is =0, then add new line
func checkLastElement(arrays [][]int) bool {
	if len(arrays) == 0 {
		return false
	}
	lastArray := arrays[len(arrays)-1]
	if len(lastArray) == 1 {
		return false
	}
	if len(lastArray) == 0 {
		return false
	}

	lastElement := lastArray[len(lastArray)-1]
	return lastElement == 0
}
