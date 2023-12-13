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

type SeedMap struct {
	seeds                 []int
	seedToSoil            []int
	soilToFertilizer      []int
	fertilizerToWater     []int
	waterToLight          []int
	lightToTemperature    []int
	temperatureToHumidity []int
	humidityToLocation    []int
}

func main() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatalln("could not get file path of caller")
	}

	filePath := filepath.Join(filepath.Dir(filename), "input.txt")

	fmt.Println("Sum of Points:", getNearestLocation(filePath))
}

func getScanner(path string) *bufio.Scanner {
	readFile, err := os.Open(path)

	if err != nil {
		log.Fatalln(err)
	}

	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)

	return scanner
}

func getNearestLocation(path string) int {
	seedMap := parseAlmanac(path)

	fmt.Println("seed map", seedMap)
	return 1
}

func parseAlmanac(path string) SeedMap {
	seedMap := SeedMap{}
	var curMap string

	scanner := getScanner(path)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()

		if i == 0 {
			seedMap.seeds = getSeeds(line)
		}

		if i > 0 {
			splitLine := strings.Split(line, " ")
			splitLineLen := len(splitLine)
			if splitLineLen == 3 {
				getMap(splitLine)
			}
			if splitLineLen == 2 {
				curMap = splitLine[0]
				fmt.Println(curMap)
				continue
			}
			if splitLineLen == 0 {
				continue
			}
		}

		fmt.Println("line", line)
	}

	return seedMap
}

func getMap(splitLine []string) []int {

	source, err := strconv.Atoi(splitLine[0])
	if err != nil {
		log.Fatalln(err)
	}
	destination, err := strconv.Atoi(splitLine[0])
	if err != nil {
		log.Fatalln(err)
	}
	mapRange, err := strconv.Atoi(splitLine[0])
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(source, destination, mapRange)
	mapSlice := []int{source, destination, mapRange}
	return mapSlice
}

func getSeeds(line string) []int {
	seedLineString := strings.Split(line, " ")
	seeds := []int{}
	for _, seedStr := range seedLineString[1:] {
		seed, err := strconv.Atoi(seedStr)
		if err != nil {
			log.Fatalln(err)
		}
		seeds = append(seeds, seed)
	}

	return seeds
}
