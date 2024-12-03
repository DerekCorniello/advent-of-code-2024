package day2

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func parseFile(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(fmt.Sprintf("Failed to open file! Error: %v. Aborting.", err.Error()))
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := [][]int{}

	for scanner.Scan() {
		nums := []int{}
		line := scanner.Text()
		splitLine := strings.Fields(line)

		for i := range len(splitLine) {
			entry, err := strconv.Atoi(splitLine[i])
			if err != nil {
				panic(fmt.Sprintf("Error parsing string: %v. Aborting.", entry))
			}
			nums = append(nums, entry)
		}
		lines = append(lines, nums)
	}

	err = scanner.Err()

	if err != nil {
		panic(fmt.Sprintf("Error scanning file! Error: %v. Aborting", err.Error()))
	}

	return lines
}

func processLineP1(ints []int) bool {
	negFlag := false
	n := len(ints)
	for index, val := range ints {
		// if it is the last entry, it has to be good
		if index == n-1 {
			continue
		}

		diff := val - ints[index+1]

		if index == 0 {
			if diff < 0 {
				negFlag = true
			}
		}

		// is difference within constraints of 1 <= diff <= 3
		if !((math.Abs(float64(diff)) > 0) && (math.Abs(float64(diff)) < 4)) {
			return false
		}
		// is difference in correct direction?
		if !((negFlag && diff < 0) || (!negFlag && diff > 0)) {
			return false
		}
	}
	return true
}

func processLineP2(ints []int) bool {
	// Check if the report is safe without modification
	if processLineP1(ints) {
		return true
	}

    // brute force method...
	for i := 0; i < len(ints); i++ {
		firstHalf := append([]int{}, ints[:i]...)    
		secondHalf := append([]int{}, ints[i+1:]...)

		// Combine both halves and check if the result is safe
		modifiedReport := append(firstHalf, secondHalf...)
		if processLineP1(modifiedReport) {
			fmt.Printf("Report is safe by removing element at index %d: %v\n", i, modifiedReport)
			return true
		} else {
			fmt.Printf("Trying by removing element at index %d: %v\n", i, modifiedReport)
		}
	}

	// If no valid modification found, report remains unsafe
	fmt.Println("No valid modification found. Report remains unsafe.")
	return false
}

func ProcessDay2p1() int {
	lines := parseFile("day2/input.txt")
	goodCount := 0
	for _, line := range lines {
		if processLineP1(line) {
			goodCount += 1
		}
	}
	return goodCount
}

func ProcessDay2p2() int {
	lines := parseFile("day2/input.txt")
	goodCount := 0
	for _, line := range lines {
		if processLineP2(line) {
			goodCount += 1
		}
	}
	return goodCount
}
