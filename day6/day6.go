package day6

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sync"
)

func parseFile(filename string) map[[2]int]string {
	file, err := os.Open(filename)
	if err != nil {
		panic(fmt.Sprintf("Failed to open file! %v. Aborting.", err))
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	// Initialize the map with coordinates as keys and empty strings as values
	board := make(map[[2]int]string)

	// Read each line and update the map based on the rune encountered
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		for x, ch := range line {
			switch ch {
			case '.':
				// Safe spot
				board[[2]int{x, y}] = "safe"
			case '#':
				// Blocked spot
				board[[2]int{x, y}] = "blocked"
			case '^':
				// Guard spot
				board[[2]int{x, y}] = "guard"
			}
		}
		y++
	}

	err = scanner.Err()
	if err != nil {
		panic(fmt.Sprintf("Error scanning file! Error: %v. Aborting", err.Error()))
	}

	return board
}

// TIL, maps are *always* passed as reference

func makeNextMove(board map[[2]int]string, dir *[2]int, guard *[2]int) bool {
	// set the current guard position as unsafe
	board[[2]int{guard[0], guard[1]}] = "unsafe"

	nextPos, ok := board[[2]int{guard[0] + dir[0], guard[1] + dir[1]}]

	// this handles if the next line is oob
	if !ok {
		return true
	}

	if nextPos != "blocked" {
		// move forward
		guard[0] += dir[0]
		guard[1] += dir[1]
	} else {
		// turn right (90 degrees clockwise)
		dir[0], dir[1] = -dir[1], dir[0]

		//                [0,-1]
		//
		//     [-1,0]       x       [1,0]
		//
		//                [0, 1]

	}

	return false
}
func findGuard(board map[[2]int]string) [2]int {
	for coord, label := range board {
		if label == "guard" {
			return coord
		}
	}
	panic("Guard not found!")
}

func ProcessDay6p1() int {
	board := parseFile("day6/input.txt")

	// always starts up
	dir := [2]int{0, -1}
	movesDone := false
	guardPos := findGuard(board)

	for !movesDone {
		movesDone = makeNextMove(board, &dir, &guardPos)
	}

	badSpots := 0
	var wg sync.WaitGroup
	var mut sync.Mutex
	for _, val := range board {
		wg.Add(1)
		go func(val string) {
			defer wg.Done()
			if val == "unsafe" {
				mut.Lock()
				badSpots++
				mut.Unlock()
			}
		}(val)
	}
	wg.Wait()
	return badSpots
}

func deepCopyBoard(board map[[2]int]string) map[[2]int]string {
	newBoard := make(map[[2]int]string)
	for key, value := range board {
		newBoard[key] = value
	}
	return newBoard
}

// *sigh* brute force method is fastest
// other approach was to create a 'ghost' guard
// that goes on the same path as the current ghost, but 
// with a blocker in front of it. Tried both on my
// machine, maybe go's wg stuff is really good 
func ProcessDay6p2() int {
	board := parseFile("day6/input.txt")

	// always starts up
	dir := [2]int{0, -1}
	guardPos := findGuard(board)

	var wg sync.WaitGroup
	var mut sync.Mutex
	cycles := 0

	// iterate through all "safe" spaces and treat each as a blocker
	for coord, label := range board {
		if label == "safe" {
			// create a deep copy of the board to avoid modifying the original one
			// mark this "safe" space as "blocked"
			boardCopy := deepCopyBoard(board)
			boardCopy[coord] = "blocked"

			wg.Add(1)
			go func(board map[[2]int]string, dir [2]int, guard [2]int) {
				defer wg.Done()
				guardDone := false

				// initialize visited map to track positions and directions
				visited := make(map[[2]int][][2]int)
				visited[guard] = [][2]int{dir}

				for !guardDone {
					guardDone = makeNextMove(board, &dir, &guard)
					if guardDone {
						break
					}

					// check if this position and direction have been visited
					dirs, ok := visited[guard]
					if ok {
						if slices.Contains(dirs, dir) {
							// cycle detected, increment the cycle count
							mut.Lock()
							cycles++
							mut.Unlock()
							guardDone = true
						} else {
							// add this new direction to the visited map
							visited[guard] = append(dirs, dir)
						}
					} else {
						// first time visiting this position with this direction
						visited[guard] = [][2]int{dir}
					}
				}
			}(boardCopy, [2]int{dir[0], dir[1]}, [2]int{guardPos[0], guardPos[1]})
		}
	}

	// Wait for all goroutines to finish
	wg.Wait()

	return cycles
}
