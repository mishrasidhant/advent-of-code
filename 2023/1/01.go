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
	/*
		1. Open file
		2. Read file line by line
			a. store line as string
			b. Iterate over string and identify 1st and last digit in string
				- Update: need to maintain a slice of words and run string comparison with words
					a) create slice of strings that store each word []string
					b)
			c. Combine 1st and last digit to create a number
		3. Add all numbers (1 per line of input in file) and return as solution
	*/

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
		// fmt.Printf("%s", "---------------------")
		// fmt.Printf("%v", reflect.TypeOf(line))
		// fmt.Printf("%s", "---------------------")
		// fmt.Print(line)
		fmt.Println("---------------------")
		fmt.Println(line)
		for _, char := range line {
			// fmt.Printf("%v", reflect.TypeOf(line[i]))
			// fmt.Println(string(line[i]))
			// Check for digit!
			// TODO: Flow of logic needs changing
			/*
				if you encounter digit check substring (add if found) and clear it then append the digit
				else add char to substring and check if substring length > 3 trigger a comparison
				comparison needs to be a function
					loop through the substring and find all matches
					store the first and last index of matches
					append the digit that has the smallest first index
					resize the substring to start from last index of first match
						repeat comparison
			*/
			if unicode.IsDigit(char) {
				// if len(subStr) != 0 {
				// 	// Trigger a comparison and store any digits you find here
				// }
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
		// Account for remaining substring if line doesn't end in digit
		// if len(subStr) != 0 {
		// 	for i := 0; i < len(numStr); i++ {
		// 		if strings.Contains(string(subStr), numStr[i]) {
		// 			digits = append(digits, i+1)
		// 		}
		// 	}
		// 	// fmt.Println(string(subStr))
		// 	subStr = nil
		// }
		// sum += firstDigit*10 + lastDigit
		// firstDigit, lastDigit = 0, 0
		// isFirst = true
		fmt.Println(digits)
		if len(digits) != 0 {
			sum = sum + (digits[0]*10 + digits[len(digits)-1])
			continue
		}
		sum = sum + (digits[0]*10 + digits[0])
	}
	// fmt.Printf("digits: %v", digits)
	return sum
}

func main() {
	fmt.Printf("Result: %v\n", process())
}
