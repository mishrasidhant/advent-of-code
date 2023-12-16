package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	// Open file
	// filePtr, err := os.Open("./test.txt")
	filePtr, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer filePtr.Close()

	// Create file scanner
	scanner := bufio.NewScanner(filePtr)
	// Split file into lins using newline as deliminator
	// scanner.Split(bufio.ScanLines)
	// Process each line
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		var firstDigit, lastDigit int
		isFirst := true
		for _, char := range line {
			if unicode.IsDigit(char) && isFirst {
				isFirst = false
				firstDigit, err = strconv.Atoi(string(char))
			}
			if unicode.IsDigit(char) && !isFirst {
				lastDigit, err = strconv.Atoi(string(char))
			}
		}
		sum += firstDigit*10 + lastDigit
		firstDigit, lastDigit = 0, 0
		isFirst = true
	}
	return sum
}

func main() {
	fmt.Printf("Result: %v\n", process())
}
