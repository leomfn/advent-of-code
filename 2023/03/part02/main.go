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

	fmt.Println("Sum of Game Powers:", getSumOfGearRatios(filePath))
}

func getSumOfGearRatios(path string) int {
	sumOfGearRatios := 0

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
		starIndices := getStarsFromLine(line)

		for _, starIndex := range starIndices {
			checkLines := []string{line}
			if i > 0 {
				checkLines = append(checkLines, lines[i-1])
			}

			if i < len(lines)-1 {
				checkLines = append(checkLines, lines[i+1])
			}

			adjacentNumbers := findAdjacentNumbers(checkLines, starIndex[0])

			if len(adjacentNumbers) == 2 {
				// fmt.Println("Adding gear ratio from these adjacent numbers", adjacentNumbers)
				sumOfGearRatios += adjacentNumbers[0] * adjacentNumbers[1]
			}
		}
	}

	return sumOfGearRatios
}

func findAdjacentNumbers(lines []string, starIndex int) []int {
	var adjacentNumbers = []int{}
	regex, _ := regexp.Compile(`\d+`)

	for _, line := range lines {
		foundNumbers := regex.FindAllIndex([]byte(line), -1)

		for _, numberIndex := range foundNumbers {
			if starIndex >= numberIndex[0]-1 && starIndex <= numberIndex[1] {
				number, err := strconv.Atoi(line[numberIndex[0]:numberIndex[1]])
				if err != nil {
					log.Fatal(err)
				}

				adjacentNumbers = append(adjacentNumbers, number)
			}
		}
	}

	return adjacentNumbers
}

func getStarsFromLine(line string) [][]int {
	regex, _ := regexp.Compile(`\*`)
	foundStarIndices := regex.FindAllIndex([]byte(line), -1)

	return foundStarIndices
}
