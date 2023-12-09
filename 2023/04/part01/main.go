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
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatalln("could not get file path of caller")
	}

	filePath := filepath.Join(filepath.Dir(filename), "input.txt")

	fmt.Println("Sum of Points:", getSumOfPoints(filePath))
}

func getScanner(path string) *bufio.Scanner {
	readFile, err := os.Open(path)

	if err != nil {
		log.Fatalln(err)
	}

	scanner := bufio.NewScanner(readFile)

	return scanner
}

func getSumOfPoints(path string) uint {
	var sumOfPoints uint = 0

	scanner := getScanner(path)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		sumOfPoints += getPointsFromLine(line)
	}

	return sumOfPoints
}

func getPointsFromLine(line string) uint {
	winningNumbers := getWinningNumbers(line)
	return getPointsFromWinningNumbers(winningNumbers)
}

func getPointsFromWinningNumbers(winningNumbers []uint) uint {
	amountOfWinningNumbers := len(winningNumbers)

	if amountOfWinningNumbers == 0 {
		return 0
	}

	return 1 << (amountOfWinningNumbers - 1) // I have almost no idea 💡
}

func getWinningNumbers(line string) []uint {
	var winningNumbers = []uint{}

	_, numbersStr, _ := strings.Cut(line, ":")
	cardWinningNumbersStr, guessedNumbersStr, _ := strings.Cut(numbersStr, "|")

	cardWinningNumbers := parseNumbersString(cardWinningNumbersStr)
	guessedNumbers := parseNumbersString(guessedNumbersStr)
	fmt.Println(cardWinningNumbers)
	fmt.Println(guessedNumbers)

	for _, winningNumber := range cardWinningNumbers {
		for _, guessedNumber := range guessedNumbers {
			if winningNumber == guessedNumber {
				winningNumbers = append(winningNumbers, winningNumber)
				continue
			}
		}
	}

	return winningNumbers
}

func parseNumbersString(numberStr string) []uint {
	var numberArray []uint

	for i := 0; i <= len(numberStr)-3; i += 3 {
		curNumberStr := strings.TrimSpace(numberStr[i : i+3])
		curNumber, err := strconv.ParseUint(curNumberStr, 10, 8)
		if err != nil {
			log.Fatalln(err)
		}
		numberArray = append(numberArray, uint(curNumber))
	}

	return numberArray
}
