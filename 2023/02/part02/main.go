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

	fmt.Println("Sum of Game Powers:", getSumOfPowers(filePath))
}

func getSumOfPowers(path string) int {
	sumOfPowers := 0

	readFile, err := os.Open(path)

	if err != nil {
		log.Fatalln(err)
	}

	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		sumOfPowers += getPowerFromLine(line)
	}

	return sumOfPowers
}

func getPowerFromLine(line string) int {
	colorFactors := map[string]int{
		"red":   1,
		"green": 1,
		"blue":  1,
	}

	// fmt.Println("line:", line)

	idGames := strings.Split(line, ":")

	game := strings.Split(idGames[1], ";")

	for _, pulls := range game {
		for _, pull := range strings.Split(pulls, ",") {
			colorPullSlice := strings.Split(strings.TrimSpace(pull), " ")

			color := colorPullSlice[1]
			pullAmount, err := strconv.Atoi(colorPullSlice[0])
			if err != nil {
				log.Fatal(err)
			}

			if pullAmount > colorFactors[color] {
				colorFactors[color] = pullAmount
			}

		}
	}

	power := colorFactors["red"] * colorFactors["green"] * colorFactors["blue"]

	return power
}
