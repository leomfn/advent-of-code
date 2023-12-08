package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	readFile, _ := os.Open("2023/01/part01/input.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	sum := 0
	for fileScanner.Scan() {
		sum += int(getLineSum(fileScanner.Text()))
	}
	fmt.Println("Total sum", sum)
}

func getLineSum(line string) int64 {
	reNumber := regexp.MustCompile(`\d`)
	numbers := reNumber.FindAllString(line, -1)

	firstNumber := numbers[0]
	lastNumber := numbers[len(numbers)-1]

	twoDigitNumber, _ := strconv.ParseInt(firstNumber+lastNumber, 10, 16)

	return twoDigitNumber
}
