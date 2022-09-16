package hbery_aoc2015

import (
	"fmt"
)

func Solution(day int) {
	PrintlnCenterAndPad("Advent of Code 2015", 80, "=")
	PrintlnCenter(fmt.Sprintf("Day %02d", day), 80, "~")

	switch day {
	case 1:
		Solution_Day01()
	default:
		fmt.Println("Choose day! [1..25]")
	}
}
