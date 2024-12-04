package day4

import (
	"bufio"
	"fmt"
	"os"
)

func parseFile(filename string) ([][]rune, int, int) {
	file, err := os.Open(filename)
	if err != nil {
		panic(fmt.Sprintf("Failed to open file! %v. Aborting.", err))
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var grid [][]rune

	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}

	return grid, len(grid[0]), len(grid)
}

func ProcessDay4p1() int {
	findCount := 0

	grid, width, height := parseFile("day4/input.txt")
	word := []rune("XMAS")

	// Possible movement directions (dx, dy)
	directions := [][2]int{
		{0, 1},
		{1, 0},
		{1, 1},
		{1, -1},
		{0, -1},
		{-1, 0},
		{-1, 1},
		{-1, -1},
	}

	var dfsOnXmas func(x, y, index int, dir [2]int) bool
	dfsOnXmas = func(x, y, index int, dir [2]int) bool {
		if index == len(word) {
			return true
		}

		// check boundaries and character match
		if x < 0 || y < 0 || x >= height || y >= width || grid[x][y] != word[index] {
			return false
		}

		// mark cell as visited
		temp := grid[x][y]
		grid[x][y] = rune('#')

		// move in the given direction
		res := dfsOnXmas(x+dir[0], y+dir[1], index+1, dir)

		grid[x][y] = temp
		return res
	}

	// start a DFS from each cell
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			// do a weird, directed dfs
			for _, dir := range directions {
				if dfsOnXmas(i, j, 0, dir) {
					findCount++
				}
			}
		}
	}

	return findCount
}

func ProcessDay4p2() int {
	findCount := 0

	grid, width, height := parseFile("day4/input.txt")

	checkCrossMas := func(x int, y int) bool {
        // if its not an A in the middle, its not a cross
		if grid[x][y] != rune('A') {
			return false
		}

        // check bounds
		if x <= 0 || y <= 0 || x >= height-1 || y >= width-1 {
			return false
		}

        // create a string of both corsses
		diagonal1 := []rune{grid[x-1][y-1], grid[x+1][y+1]}
		diagonal2 := []rune{grid[x-1][y+1], grid[x+1][y-1]}

		// Helper function to check if a slice forms 'MAS'
		isMas := func(d []rune) bool {
			return (d[0] == 'M' && d[1] == 'S') || (d[0] == 'S' && d[1] == 'M')
		}

		// Both diagonals must match 'MAS' pattern
		return isMas(diagonal1) && isMas(diagonal2)
	}

	// Iterate through the grid
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if checkCrossMas(i, j) {
				findCount++
			}
		}
	}

	return findCount
}
