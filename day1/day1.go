package day1

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func parseFile(filename string) ([]int, []int) {
	file, err := os.Open(filename)
	if err != nil {
		panic(fmt.Sprintf("Failed to open file! Error: %v. Aborting.", err.Error()))
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	locations1, locations2 := []int{}, []int{}

	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Fields(line)

		if len(splitLine) != 2 {
			panic(fmt.Sprintf("Invalid length, expected line length of 2, got %d.", len(splitLine)))
		}

		str, err := strconv.Atoi(splitLine[0])
		if err != nil {
			panic(fmt.Sprintf("Error parsing string: %v. Aborting.", str))
		}
		locations1 = append(locations1, str)

		str, err = strconv.Atoi(splitLine[1])
		if err != nil {
			panic(fmt.Sprintf("Error parsing string: %v. Aborting.", str))
		}
		locations2 = append(locations2, str)
	}

	err = scanner.Err()

	if err != nil {
		panic(fmt.Sprintf("Error scanning file! Error: %v. Aborting", err.Error()))
	}
	return locations1, locations2
}

func ProcessDay1p1() int {
	locations1, locations2 := parseFile("day1/input.txt")
	sort.Ints(locations1)
	sort.Ints(locations2)

	n := len(locations1)
	if n != len(locations2) {
		panic(fmt.Sprintf("Expected equal lengths, got %v and %v.", n, len(locations2)))
	}

	dist := 0
	for i := range n {
		dist += int(math.Abs(float64(locations1[i]) - float64(locations2[i])))
	}

	return dist
}

func ProcessDay1p2() int {
	locations1, locations2 := parseFile("day1/input.txt")
	countOccurences, simScores := map[int]int{}, map[int]int{}

	for _, key := range locations1 {
		_, ok := countOccurences[key]
		if !ok {
			countOccurences[key] = 0
		}
		countOccurences[key] += 1
	}

	for _, key := range locations2 {
		_, ok := simScores[key]
		if !ok {
			simScores[key] = 0
		}
		simScores[key] += key
	}

	simScore := 0

	for key, value := range countOccurences {
		simVal, ok := simScores[key]
		if !ok {
			continue
		}
		simScore += value * simVal
	}

	return simScore
}
