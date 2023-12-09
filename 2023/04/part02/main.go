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

	fmt.Println("Sum of Points:", getCountOfScratchCards(filePath))
}

func getScanner(path string) *bufio.Scanner {
	readFile, err := os.Open(path)

	if err != nil {
		log.Fatalln(err)
	}

	scanner := bufio.NewScanner(readFile)

	return scanner
}

func getCountOfScratchCards(path string) int {
	lineWinningNumbersCount := getWinningNumberCountOverview(path)
	initialCardsCount := len(lineWinningNumbersCount)

	countOfScratchCards := initialCardsCount // initialize with count if initial scratch cards

	linesToCheck := []int{}
	for l := 0; l < initialCardsCount; l++ {
		linesToCheck = append(linesToCheck, l)
	}

	i := 0
	for true {
		curLineIndex := linesToCheck[i]
		curCount := lineWinningNumbersCount[curLineIndex]

		fmt.Printf("%.1f %%\n", float64(i)/float64(len(linesToCheck))*100)

		newLinesToCheck := []int{}
		for j := 0; j < curCount; j++ {
			newLineIndex := curLineIndex + j + 1

			if newLineIndex < initialCardsCount {
				countOfScratchCards++
				newLinesToCheck = append(newLinesToCheck, newLineIndex)
			}
		}

		linesToCheck = append(linesToCheck, newLinesToCheck...)
		i++
		if i == len(linesToCheck) {
			break
		}
	}

	return countOfScratchCards
}

func getWinningNumberCountOverview(path string) []int {
	lineWinningNumbersCount := []int{}

	scanner := getScanner(path)
	scanner.Split(bufio.ScanLines)

	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		lineWinningNumbersCount = append(lineWinningNumbersCount, getCountOfWinningNumbers(line))
	}

	return lineWinningNumbersCount
}

func getCountOfWinningNumbers(line string) int {
	var countOfWinningNumbers int = 0

	_, numbersStr, _ := strings.Cut(line, ":")
	cardWinningNumbersStr, guessedNumbersStr, _ := strings.Cut(numbersStr, "|")

	cardWinningNumbers := parseNumbersString(cardWinningNumbersStr)
	guessedNumbers := parseNumbersString(guessedNumbersStr)

	for _, winningNumber := range cardWinningNumbers {
		for _, guessedNumber := range guessedNumbers {
			if winningNumber == guessedNumber {
				countOfWinningNumbers++
				continue
			}
		}
	}

	return countOfWinningNumbers
}

func parseNumbersString(numberStr string) []int {
	var numberArray []int

	for i := 0; i <= len(numberStr)-3; i += 3 {
		curNumberStr := strings.TrimSpace(numberStr[i : i+3])
		curNumber, err := strconv.Atoi(curNumberStr)
		if err != nil {
			log.Fatalln(err)
		}
		numberArray = append(numberArray, curNumber)
	}

	return numberArray
}
