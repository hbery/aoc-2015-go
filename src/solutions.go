package hbery_aoc2015

import (
	"fmt"

	hbery_aoc2015 "github.com/hbery/aoc-2015-go/src/util"
)

func Solution(day, part int, input string) error {
	hbery_aoc2015.PrintlnCenterAndPad("Advent of Code 2015", 80, "=")
	hbery_aoc2015.PrintlnCenter(fmt.Sprintf("Day %02d", day), 80, "~")

	var score int64
	var err error
	switch day {
	case 1:
		score, err = Solution_Day01(part, input)
		if err != nil {
			return err
		}
	default:
		fmt.Println("Choose day! [1..25]")
	}

	// Print Score and return
	hbery_aoc2015.PrintlnCenter(fmt.Sprintf("Result: %d", score), 80, "")
	return nil
}
