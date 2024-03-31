package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

type Match struct {
	index, number int
	parentString  string
}

type Comparator struct {
	digit, index, length int
}

// func RemoveIndex(c []Comparator, index int) []Comparator {
// 	ret := []Comparator{}
// 	ret = append(ret, c[:index]...)
// 	return append(ret, c[index+1:]...)
// }

func findMatch(subString string) []int {
	rawMatches := []Comparator{}
	matches := []int{}
	numStr := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	// Comparison Logic => Move to function and impliment multiple sorted comparisons
	// Find each match => NEED TO ENSURE YOU FIND EACH MATCH!!
	for i := 0; i < len(numStr); i++ {
		found := strings.Index(string(subString), numStr[i])
		// Keep slicing and finding till you have all matches
		if found != -1 {
			// fmt.Println("Initial String: ", subString)
			// fmt.Println("FOUND initial string match : ", numStr[i])
			rawMatches = append(rawMatches, Comparator{digit: i + 1, index: found, length: len(numStr[i])})
			if (found + len(numStr[i])) >= len(subString) {
				continue
			}
			currentString := subString[found+1:]
			increment := found
			for nestedFound := strings.Index(string(currentString), numStr[i]); nestedFound != -1; {
				// fmt.Println("FOUND repeated STRING MATCH: ", i+1)
				// fmt.Println("Substring: ", currentString)
				// fmt.Printf("Found: %v, Increment: %v, NestedFound: %v\n", found, increment, nestedFound)
				// rawMatches = append(rawMatches, Comparator{digit: i + 1, index: (increment + nestedFound), length: len(numStr[i])})
				rawMatches = append(rawMatches, Comparator{digit: i + 1, index: nestedFound + increment + 1, length: len(numStr[i])})
				if (nestedFound + len(numStr[i])) >= len(currentString) {
					break
				}
				currentString = currentString[nestedFound+1:]
				increment += nestedFound
				nestedFound = strings.Index(string(currentString), numStr[i])
				// fmt.Println("Abount to compare substring:", currentString[found+1:])
				// found = strings.Index(string(currentString[found+1:]), numStr[i])
				// fmt.Println("Found:", found)
			}
		}
	}
	// NO NEED TO SORT: You're already adding in sorted sequence, need to check length and delete while adding
	// NVM  : Need to sort because you're not adding in sequence, you're searching for individual numbers
	sort.Slice(rawMatches, func(i, j int) bool {
		return rawMatches[i].index < rawMatches[j].index
	})
	// Remove overlapping matches
	for i := 0; i < len(rawMatches)-1; i++ {
		// fmt.Printf("Raw matches: %#v", rawMatches)
		m1 := rawMatches[i].index + rawMatches[i].length - 1
		// fmt.Println("m1: ", m1)
		// fmt.Println("m2: ", rawMatches[i+1].index)
		if m1 >= rawMatches[i+1].index {
			// Delete i+1
			// fmt.Printf("Raw Matches PRE %+v: ", rawMatches)
			rawMatches = append(rawMatches[:i], rawMatches[i+1:]...)
			// fmt.Printf("Raw Matches POST DELETE: %+v", rawMatches)
		}
	}
	for _, match := range rawMatches {
		// fmt.Println(rawMatches)
		matches = append(matches, match.digit)
	}
	return matches
}

func process() int {
	// Open file -> returns a file pointer with R access
	filePtr, err := os.Open("./input.txt")
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
				// Trigger a comparison and reset substring, you've found a number
				// fmt.Println("About to compare NUMERIC DIGIT : ", string(subStr))
				match := findMatch(string(subStr))
				if len(match) > 0 {
					fmt.Println("STRING: ", match)
					digits = append(digits, match...)
				}
				// fmt.Println("NO MATCH FOUND")
				subStr = nil
				// Store the number
				digit, err = strconv.Atoi(string(char))
				if err == nil {
					fmt.Println("DIGIT: ", digit)
					digits = append(digits, digit)
					// fmt.Println("Found: ", digit)
				}
				continue
			}
			// Not a digit, append rune to subString
			subStr = append(subStr, char)
		}
		// Account for leftover substring
		if len(subStr) > 0 {
			// Trigger a comparison here
			// fmt.Println("About to compare END OF LINE: ", string(subStr))
			match := findMatch(string(subStr))
			if len(match) > 0 {
				fmt.Println("String: ", match)
				digits = append(digits, match...)
			}
			// fmt.Println("NO MATCH FOUND")
			subStr = nil
		}
		fmt.Println("Digits: :", digits)
		if len(digits) == 0 {
			fmt.Println("NO DIGITS FOUND!!")
			continue
		}
		// fmt.Println("Digits: ", digits[0]*10, digits[len(digits)-1])
		sum = sum + (digits[0]*10 + digits[len(digits)-1])
		// fmt.Println("Sum: ", sum)
	}
	return sum
}

func main() {
	fmt.Printf("Result: %v\n", process())
}

/* BUG : Should be 62, need the char buffer to slide back
---------------------
fzrpfhbfvj6dbxbtfs7twofksfbshrzkdeightwoqg
Found:  6
Found:  7
Found:  two
Found:  eight
[6 7 2 8]
Digits:  60 8
Sum:  866
*/
