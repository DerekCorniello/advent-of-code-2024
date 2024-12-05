package day5

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseFile(filename string) ([][]string, [][]string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(fmt.Sprintf("Failed to open file! %v. Aborting.", err))
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	rels, updates := [][]string{}, [][]string{}

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "|") {
			rel := strings.Split(line, "|")
			rels = append(rels, []string{rel[0], rel[1]})
		} else if strings.Contains(line, ",") {
			update := strings.Split(line, ",")
			updates = append(updates, update)
		}
	}

	err = scanner.Err()

	if err != nil {
		panic(fmt.Sprintf("Error scanning file! Error: %v. Aborting", err.Error()))
	}

	return rels, updates
}

func buildGraph(rels [][]string) map[string][]string {
	graph := map[string][]string{}
	for _, rule := range rels {
		from, to := rule[0], rule[1]
		_, ok := graph[from]
		if !ok {
			graph[from] = []string{}
		}
		graph[from] = append(graph[from], to)
	}
	return graph
}

func buildReverseGraph(rels [][]string) map[string][]string {
	reverseGraph := map[string][]string{}
	for _, rule := range rels {
		from, to := rule[0], rule[1]
		_, ok := reverseGraph[to]
		if !ok {
			reverseGraph[to] = []string{}
		}
		reverseGraph[to] = append(reverseGraph[to], from)
	}
	return reverseGraph
}

func isTopologicallySorted(update []string, graph map[string][]string) bool {
	// Create a map to check which pages have been processed
	processed := map[string]bool{}
	for _, page := range update {
		// For the current page, check its dependencies in the graph
		// If the dependency is after the current page, it's invalid
		for _, dep := range graph[page] {
			if _, exists := processed[dep]; exists {
				return false
			}
		}
		// Mark the current page as processed
		processed[page] = true
	}
	return true
}

// Kahn's alg
// step by step given because this is new to me!
func resortGraph(update []string, graph map[string][]string) []string {
	// step 1: Initialize an in-degree map
	inDegree := map[string]int{}
	for _, page := range update {
		inDegree[page] = 0
	}
	// calculate in-degrees for each page in the update
	for _, page := range update {
		for _, dep := range graph[page] {
			if _, exists := inDegree[dep]; exists {
				inDegree[dep]++
			}
		}
	}

	// step 2: Initialize a queue with all pages having 0 in-degree
	queue := []string{}
	for page, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, page)
		}
	}

	// step 3: Perform Kahn's algorithm (topological sort)
	sortedPages := []string{}
	for len(queue) > 0 {
		// get the next page with 0 in-degree
		page := queue[0]
		queue = queue[1:]
		sortedPages = append(sortedPages, page)

		// decrease the in-degree of its neighbors
		for _, dep := range graph[page] {
			inDegree[dep]--
			if inDegree[dep] == 0 {
				queue = append(queue, dep)
			}
		}
	}

	// if sortedPages doesn't have all pages, it means there's a cycle
	if len(sortedPages) != len(update) {
		panic("Cycle detected or incomplete sorting!")
	}

	// step 4: Return the newly sorted pages
	return sortedPages
}

func ProcessDay5p1() int {
	relations, updates := parseFile("day5/input.txt")
	graph := buildGraph(relations)

	validMiddleValues := []int{}

	for _, update := range updates {
		if isTopologicallySorted(update, graph) {
			n := len(update)
			middleValue := update[n/2]

			// take left val if even
			if n%2 == 0 {
				middleValue = update[(n/2)-1]
			}

			mid, err := strconv.Atoi(middleValue)
			if err != nil {
				panic(fmt.Sprintf("Invalid middle value: %v. Aborting.", err))
			}
			validMiddleValues = append(validMiddleValues, mid)
		}
	}

	sum := 0
	for _, val := range validMiddleValues {
		sum += val
	}

	return sum
}

func ProcessDay5p2() int {
	relations, updates := parseFile("day5/input.txt")
	graph := buildGraph(relations)

	validMiddleValues := []int{}

	for _, update := range updates {
		if !isTopologicallySorted(update, graph) {
			// Resort the update if it's not in order
			update = resortGraph(update, graph)
            // now do middle val
            n := len(update)
            middleValue := update[n/2]

            // take left val if even
            if n%2 == 0 {
                middleValue = update[(n/2)-1]
            }

            mid, err := strconv.Atoi(middleValue)
            if err != nil {
                panic(fmt.Sprintf("Invalid middle value: %v. Aborting.", err))
            }
            validMiddleValues = append(validMiddleValues, mid)
		}
	}

	// Compute the sum of all valid middle values
	sum := 0
	for _, val := range validMiddleValues {
		sum += val
	}

	return sum
}
