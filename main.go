package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/DerekCorniello/advent-of-code-2024/day1"
	"github.com/DerekCorniello/advent-of-code-2024/day2"
	"github.com/DerekCorniello/advent-of-code-2024/day3"
	"github.com/DerekCorniello/advent-of-code-2024/day4"
	"github.com/DerekCorniello/advent-of-code-2024/day5"
	"github.com/DerekCorniello/advent-of-code-2024/day6"
)

func main() {
    numDays := 6

 	if len(os.Args) < 2 {
 		fmt.Println("Please provide a day number as an argument.")
 		return
 	}

	day, err := strconv.Atoi(os.Args[1])
	if err != nil || day < 1 || day > numDays {
		fmt.Printf("Please provide a valid day number (1-%d).\n", numDays)
		return
	}

	switch day {
	case 1:
		fmt.Printf("Output of Day 1: `%v`, `%v`.\n", day1.ProcessDay1p1(), day1.ProcessDay1p2())
	case 2:
		fmt.Printf("Output of Day 2: `%v`, `%v`.\n", day2.ProcessDay2p1(), day2.ProcessDay2p2())
	case 3:
		fmt.Printf("Output of Day 3: `%v`, `%v`.\n", day3.ProcessDay3p1(), day3.ProcessDay3p2())
	case 4:
		fmt.Printf("Output of Day 4: `%v`, `%v`.\n", day4.ProcessDay4p1(), day4.ProcessDay4p2())
	case 5:
		fmt.Printf("Output of Day 5: `%v`, `%v`.\n", day5.ProcessDay5p1(), day5.ProcessDay5p2())
	case 6:
		fmt.Printf("Output of Day 6: `%v`, `%v`.\n", day6.ProcessDay6p1(), day6.ProcessDay6p2())
	}
}
