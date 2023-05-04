package hbery_aoc2015

import (
	"errors"
	"fmt"

	hbery_aoc2015 "github.com/hbery/aoc-2015-go/src/util"
)

type SolutionDay func(int, string) (int64, error)

var SolutionsMap = map[int]SolutionDay{
	1:  Solution_Day01,
	2:  Solution_Day02,
	3:  Solution_Day03,
	4:  Solution_Day04,
	5:  Solution_Day05,
	35: Solution_Day05_regex,
	6:  Solution_Day06,
	36: Solution_Day06_regex,
	7:  Solution_Day07,
	8:  Solution_Day08,
	9:  Solution_Day09,
	10:  Solution_Day10,
	11:  Solution_Day11,
}

func Solution(day, part int, input string) error {
	hbery_aoc2015.PrintlnCenterAndPad("Advent of Code 2015", 80, "=")
	hbery_aoc2015.PrintlnCenter(fmt.Sprintf("Day %02d", day), 80, "~")

	var result int64
	var err error
	if fn, key_exists := SolutionsMap[day]; key_exists {
		if result, err = fn(part, input); err != nil {
			return err
		}
	} else {
		return errors.New("Choose day! [1..25]")
	}

	// Print Score and return
	hbery_aoc2015.PrintlnCenter(fmt.Sprintf("Result: %d", result), 80, "")
	return nil
}
