package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	_, curPath, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("unable to get the current file path")
	}

	dataPath := filepath.Join(filepath.Dir(curPath), "input.txt")

	fmt.Println("Total sum:", totalSum(dataPath))
}

func totalSum(path string) int {
	readFile, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	sum := 0

	f, err := os.Create(filepath.Join(filepath.Dir(path), "log.txt"))

	if err != nil {
		log.Fatal(err)
	}

	for fileScanner.Scan() {
		sum += int(getLineNumber(fileScanner.Text(), &f))
		fmt.Fprintln(f, "sum", sum)
	}

	f.Close()

	return sum
}

func getLineNumber(line string, f **os.File) int64 {
	var firstNumber string
	var firstNumberIndex int = len(line)
	var lastNumber string
	var lastNumberIndex int = -1

	options := [18]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for _, v := range options {
		firstIndex := strings.Index(line, v)
		if firstIndex != -1 && firstIndex < firstNumberIndex {
			firstNumber = v
			firstNumberIndex = firstIndex
		}

		lastIndex := strings.LastIndex(line, v)
		if lastIndex > lastNumberIndex {
			lastNumber = v
			lastNumberIndex = lastIndex
		}
	}

	// A written out number string has a length > 1, so translate it to a one-digit number string
	if len(firstNumber) > 1 {
		firstNumber = translateNumberString(firstNumber)
	}
	if len(lastNumber) > 1 {
		lastNumber = translateNumberString(lastNumber)
	}

	twoDigitNumber, _ := strconv.ParseInt(firstNumber+lastNumber, 10, 16)

	fmt.Fprintln(*f, line, "firstNumber", firstNumber, "lastNumber", lastNumber, "twoDigitNumber", twoDigitNumber)

	return twoDigitNumber
}

func translateNumberString(numberString string) string {
	m := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	return m[numberString]
}
