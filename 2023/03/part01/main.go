package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
)

type Number struct {
	number int
	index  []int
}

func main() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatalln("could not get file path of caller")
	}

	filePath := filepath.Join(filepath.Dir(filename), "input.txt")

	fmt.Println("Sum of Game Powers:", getSumOfPartNumbers(filePath))
}

func getSumOfPartNumbers(path string) int {
	sumOfPartNumbers := 0

	readFile, err := os.Open(path)

	if err != nil {
		log.Fatalln(err)
	}

	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	for i, line := range lines {
		numbers := getNumbersFromLine(line)

		for _, number := range numbers {
			checkLines := []string{line}
			if i > 0 {
				checkLines = append(checkLines, lines[i-1])
			}

			if i < len(lines)-1 {
				checkLines = append(checkLines, lines[i+1])
			}

			isValid := validateNumbers(checkLines, number.index)
			if isValid {
				sumOfPartNumbers += number.number
			}
		}

	}

	return sumOfPartNumbers
}

func validateNumbers(lines []string, index []int) bool {
	isValid := false

	for _, line := range lines {
		startIndex := index[0]
		if startIndex > 0 {
			startIndex -= 1
		}

		endIndex := index[1]
		if endIndex < len(line) {
			endIndex += 1
		}

		if checkSymbols(line[startIndex:endIndex]) {
			isValid = true
			return isValid
		}
	}
	return isValid
}

func checkSymbols(linePart string) bool {
	regex, _ := regexp.Compile(`[^\.\d]`)
	foundSymbol := regex.Match([]byte(linePart))

	return foundSymbol
}

func getNumbersFromLine(line string) []Number {
	var numbers []Number

	regex, _ := regexp.Compile(`\d+`)

	foundNumberIndices := regex.FindAllIndex([]byte(line), -1)

	for _, numberIndex := range foundNumberIndices {
		number, err := strconv.Atoi(line[numberIndex[0]:numberIndex[1]])
		if err != nil {
			log.Fatal(err)
		}

		numbers = append(numbers, Number{number: number, index: numberIndex})
	}

	return numbers
}
