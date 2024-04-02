package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func process(line string) string {
	return line
}

func main() {
	args := os.Args
	var fileName string
	if len(args) < 2 {
		fileName = "input.txt"
	} else {
		fileName = args[1]
	}
	fileRelativePath := fmt.Sprintf("./%s", fileName)

	filePtr, err := os.Open(fileRelativePath)
	check(err)
	defer filePtr.Close()

	scanner := bufio.NewScanner(filePtr)
	for scanner.Scan() {
		line := scanner.Text()
		// Skip empty lines
		if len(line) == 0 {
			continue
		}
		fmt.Println(process(line))
	}
}
