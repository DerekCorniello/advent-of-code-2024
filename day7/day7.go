package day7

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseFile(filename string) ([]int, [][]int) {
	file, err := os.Open(filename)
	if err != nil {
		panic(fmt.Sprintf("Failed to open file! %v. Aborting.", err))
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	results, operators := []int{}, [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, ": ")
		thisRes, err := strconv.Atoi(split[0])
		if err != nil {
			panic(fmt.Sprintf("Error parsing result int! %v", err))
		}
		thisOps := []int{}
		listOps := strings.Split(split[1], " ")
		for _, val := range listOps {
			parsedVal, err := strconv.Atoi(val)
			if err != nil {
				panic(fmt.Sprintf("Error parsing operand! %v", err))
			}
			thisOps = append(thisOps, parsedVal)
		}
		results = append(results, thisRes)
		operators = append(operators, thisOps)
	}

	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("Error scanning file! Error: %v. Aborting", err))
	}

	return results, operators
}

func concatInts(a, b int) int {
	result := fmt.Sprintf("%d%d", a, b)
	concatenatedResult, err := strconv.Atoi(result)
	if err != nil {
		fmt.Println("Error in concatenation:", err)
		return 0
	}
	return concatenatedResult
}

func ProcessDay7p1() int {
	res, ops := parseFile("day7/input.txt")
	totalSum := 0

	var recurse func(currVal int, opsLeft []int, target int) bool

	recurse = func(currVal int, opsLeft []int, target int) bool {
		if currVal == target {
			return true
		}
		if len(opsLeft) == 0 {
			return false
		}
        // recursively check if either adding or multiplying gets the solution
		return recurse(currVal+opsLeft[0], opsLeft[1:], target) || recurse(currVal*opsLeft[0], opsLeft[1:], target)
	}

	for i := range res {
		target := res[i]
		if recurse(ops[i][0], ops[i][1:], target) {
			totalSum += target
		}
	}

	return totalSum
}

func ProcessDay7p2() int {
	res, ops := parseFile("day7/input.txt")
	totalSum := 0

	var recurse func(currVal int, opsLeft []int, target int) bool
	recurse = func(currVal int, opsLeft []int, target int) bool {
		if len(opsLeft) == 0 {
			return currVal == target
		}
        // check each operator, separated by ifs dor readability
		last := opsLeft[0] // only concats the first and second int
		if recurse(currVal+last, opsLeft[1:], target) {
			return true
		}
		if recurse(currVal*last, opsLeft[1:], target) {
			return true
		}
		concatenatedValue := concatInts(currVal, last)
		if recurse(concatenatedValue, opsLeft[1:], target) {
			return true
		}
		return false
	}

	for i := range res {
		target := res[i]
		if recurse(ops[i][0], ops[i][1:], target) {
			totalSum += target
		}
	}

	return totalSum
}
