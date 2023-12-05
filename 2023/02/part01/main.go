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

	getSumOfGameIds(filePath)
}

func getSumOfGameIds(path string) int {
	readFile, err := os.Open(path)

	if err != nil {
		log.Fatalln(err)
	}

	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)

	maxCounts := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	var validGameSum int = 0

	for scanner.Scan() {
		line := scanner.Text()
		idGames := strings.Split(line, ":")

		id, err := strconv.Atoi(strings.Replace(idGames[0], "Game ", "", 1))
		if err != nil {
			log.Fatal(err)
		}

		game := strings.Split(idGames[1], ";")

		gameIsValid := true

		for _, pulls := range game {
			for _, pull := range strings.Split(pulls, ",") {
				colorPullSlice := strings.Split(strings.TrimSpace(pull), " ")

				color := colorPullSlice[1]
				pullAmount, err := strconv.Atoi(colorPullSlice[0])
				if err != nil {
					log.Fatal(err)
				}

				if pullAmount > maxCounts[color] {
					gameIsValid = false
				}

			}
		}

		if gameIsValid {
			validGameSum += id
		}
	}

	fmt.Println("Sum of valid Game IDs:", validGameSum)

	readFile.Close()

	return validGameSum
}
