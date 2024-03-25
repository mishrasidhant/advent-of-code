package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func process() int {
	// Open file -> returns a file pointer with R access
	filePtr, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer filePtr.Close()

	// Create file scanner  -> takes a os.Reader file pointer, defaults to /n delimiter
	scanner := bufio.NewScanner(filePtr)
	// Split file into lins using newline as deliminator
	// scanner.Split(bufio.ScanLines)

	// Process each line
	sum := 0
	numStr := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for scanner.Scan() {
		var digit int
		digits := []int{}
		subStr := []rune{}
		line := scanner.Text()
		// Guard for empty lines
		if len(line) == 0 {
			continue
		}
		fmt.Println("---------------------")
		fmt.Println(line)
		for _, char := range line {
			// Check for digit!
			if unicode.IsDigit(char) {
				digit, err = strconv.Atoi(string(char))
				if err == nil {
					digits = append(digits, digit)
				}
				continue
			}
			// Not a digit, begin to store substring
			subStr = append(subStr, char)
			for i := 0; i < len(numStr); i++ {
				if strings.Contains(string(subStr), numStr[i]) {
					digits = append(digits, i+1)
					fmt.Println("Found: ", numStr[i])
					subStr = nil
					break
				}
			}
		}
		fmt.Println(digits)
		if len(digits) != 0 {
			sum = sum + (digits[0]*10 + digits[len(digits)-1])
			continue
		}
		sum = sum + (digits[0]*10 + digits[0])
	}
	return sum
}

func main() {
	fmt.Printf("Result: %v\n", process())
}
