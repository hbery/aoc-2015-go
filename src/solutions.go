package hbery_aoc2015

import (
	"fmt"

	hbery_aoc2015 "github.com/hbery/aoc-2015-go/src/util"
)

func Solution(day int) {
	hbery_aoc2015.PrintlnCenterAndPad("Advent of Code 2015", 80, "=")
	hbery_aoc2015.PrintlnCenter(fmt.Sprintf("Day %02d", day), 80, "~")

	switch day {
	case 1:
		Solution_Day01()
	default:
		fmt.Println("Choose day! [1..25]")
	}
}
