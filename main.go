package main

import (
	"fmt"
	"github.com/DerekCorniello/advent-of-code-2024/day1"
	"github.com/DerekCorniello/advent-of-code-2024/day2"
	"github.com/DerekCorniello/advent-of-code-2024/day3"
	"github.com/DerekCorniello/advent-of-code-2024/day4"
)

func main() {
	fmt.Printf("Output of Day 1: `%v`, `%v`.\n", day1.ProcessDay1p1(), day1.ProcessDay1p2())
	fmt.Printf("Output of Day 2: `%v`, `%v`.\n", day2.ProcessDay2p1(), day2.ProcessDay2p2())
	fmt.Printf("Output of Day 3: `%v`, `%v`.\n", day3.ProcessDay3p1(), day3.ProcessDay3p2())
	fmt.Printf("Output of Day 4: `%v`, `%v`.\n", day4.ProcessDay4p1(), day4.ProcessDay4p2())
}
