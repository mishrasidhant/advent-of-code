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
			c. Combine 1st and last digit to create a number
		3. Add all numbers (1 per line of input in file) and return as solution
	*/

	// Use os.open() function to open the file.
	// filePtr, err := os.Open("./test.txt")
	filePtr, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer filePtr.Close()

	// Use bufio.NewScanner() function to create the file scanner.
	scanner := bufio.NewScanner(filePtr)
	// Use bufio.ScanLines() function with the scanner to split the file into lines => DEFAULT behaviour, not required
	// scanner.Split(bufio.ScanLines)
	// Then use the scanner Scan() function in a for loop to get each line and process it.
	sum := 0
	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		line := scanner.Text()
		// fmt.Println(reflect.TypeOf(line))
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
		// fmt.Println(line)
		// fmt.Printf("First digit: %d \nLast digit: %d\n", firstDigit, lastDigit)
		sum += firstDigit*10 + lastDigit
		firstDigit, lastDigit = 0, 0
		isFirst = true

		// for i := 0; i < len(line); i++ {
		// 	fmt.Println(reflect.TypeOf(line[i]))
		// 	break
		// 	fmt.Println(line[i])
		// }
		// Iterate over string and identify first and last digit
		// Combine to create a number
		// Add to sum
	}
	return sum
}

func main() {
	fmt.Printf("Result: %v\n", process())
}
