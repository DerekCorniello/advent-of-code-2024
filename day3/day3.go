package day3

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func parseFile(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		panic(fmt.Sprintf("Failed to open file! Error: %v. Aborting.", err.Error()))
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var inputString string
	for scanner.Scan() {
		inputString = inputString + scanner.Text()
	}

	return inputString
}

func getLineSums(line string) int {
	// matches the mul(num,num) syntax
	re := regexp.MustCompile(`mul\((\d{0,3}),(\d{0,3})\)`)

	matches := re.FindAllStringSubmatch(line, -1)

	var sum int
	for _, match := range matches {
		i1, err := strconv.Atoi(match[1])
		if err != nil {
			panic(fmt.Sprintf("Error Parsing Int: %v", i1))
		}
		i2, err := strconv.Atoi(match[2])
		if err != nil {
			panic(fmt.Sprintf("Error Parsing Int: %v", i2))
		}
		sum += i1 * i2
	}

	return sum
}

func ProcessDay3p1() int {
	inputString := parseFile("day3/input.txt")
	return getLineSums(inputString)
}

func ProcessDay3p2() int {
	inputString := parseFile("day3/input.txt")
	re := regexp.MustCompile(`(?:do|don't)\(\)`)

    // we can separate commands by execute vs do not execute by
    // creating lines that start with do or don't, we can use more regex
	// this does a find and replace, returns string
	processedString := re.ReplaceAllStringFunc(inputString,
        // anonymous func that takes the match and prepends newline to format
		func(match string) string {
			return "\n" + match
		})
	commandsList := strings.Split(processedString, "\n")
	fmt.Printf("%v", commandsList)

	var sum int
	for _, line := range commandsList {
		if !strings.HasPrefix(line, "don't") {
			sum += getLineSums(line)
		}
	}

	return sum
}
