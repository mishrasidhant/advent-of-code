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

type Comparator struct {
	digit, index, length int
}

func RemoveIndex(c []Comparator, index int) []Comparator {
	ret := []Comparator{}
	ret = append(ret, c[:index]...)
	return append(ret, c[index+1:]...)
}

func findMatch(subString string, isFirst bool) []int {
	var localBool = isFirst
	rawMatches := []Comparator{}
	matches := []int{}
	numStr := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for i := 0; i < len(numStr); i++ {
		found := strings.Index(string(subString), numStr[i])
		// Keep slicing and finding till you have all matches
		if found != -1 {
			rawMatches = append(rawMatches, Comparator{digit: i + 1, index: found, length: len(numStr[i])})
			if (found + len(numStr[i])) >= len(subString) {
				continue
			}
			currentString := subString[found+1:]
			increment := found
			for nestedFound := strings.Index(string(currentString), numStr[i]); nestedFound != -1; {
				rawMatches = append(rawMatches, Comparator{digit: i + 1, index: nestedFound + increment + 1, length: len(numStr[i])})
				if (nestedFound + len(numStr[i])) >= len(currentString) {
					break
				}
				currentString = currentString[nestedFound+1:]
				increment += nestedFound
				nestedFound = strings.Index(string(currentString), numStr[i])
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
		m1 := rawMatches[i].index + rawMatches[i].length - 1
		if m1 >= rawMatches[i+1].index {
			if localBool {
				localBool = false
				rawMatches = append(rawMatches[:i+1], rawMatches[i+2:]...)
			} else {
				rawMatches = append(rawMatches[:i], rawMatches[i+1:]...)
			}
		}
	}
	for _, match := range rawMatches {
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
		isFirst := true
		for _, char := range line {
			// Check for digit!
			if unicode.IsDigit(char) {
				// Trigger a comparison and reset substring, you've found a number
				match := findMatch(string(subStr), isFirst)
				if len(match) > 0 {
					fmt.Println("STRING: ", match)
					digits = append(digits, match...)
					if len(digits) > 0 {
						isFirst = false
					}
				}
				subStr = nil
				// Store the number
				digit, err = strconv.Atoi(string(char))
				if err == nil {
					fmt.Println("DIGIT: ", digit)
					digits = append(digits, digit)
					if len(digits) > 0 {
						isFirst = false
					}
				}
				continue
			}
			// Not a digit, append rune to subString
			subStr = append(subStr, char)
		}
		// Account for leftover substring
		if len(subStr) > 0 {
			// Trigger a comparison here
			match := findMatch(string(subStr), isFirst)
			if len(match) > 0 {
				fmt.Println("String: ", match)
				digits = append(digits, match...)
				if len(digits) > 0 {
					isFirst = false
				}
			}
			subStr = nil
		}
		fmt.Println("Digits: :", digits)
		if len(digits) == 0 {
			fmt.Println("NO DIGITS FOUND!!")
			continue
		}
		sum = sum + (digits[0]*10 + digits[len(digits)-1])
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

/* BUG : twotwo => need first two is getting deleted, fix index off by 1
twotwosevenonelzlpnmkdqq2rqthff
STRING:  [2 7 1]
DIGIT:  2
Digits: : [2 7 1 2]
*/
