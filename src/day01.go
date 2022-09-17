package hbery_aoc2015

import (
	"errors"
	"fmt"
	"strings"
)

func day01_p1(input string) (int64, error) {
	var score int64 = 0

	input = strings.Replace(input, "\n", "", 1)

	for idx, char := range input {
		if char == '(' {
			score++
		} else if char == ')' {
			score--
		} else {
			return score, errors.New(fmt.Sprintf("Bad character input %q in column %d.", char, idx))
		}
	}

	return score, nil
}

func day01_p2(input string) (int64, error) {
	return 0, nil
	var score int64 = 0

	input = strings.Replace(input, "\n", "", 1)

	for idx, char := range input {
		if char == '(' {
			score++
		} else if char == ')' {
			score--
		} else {
			return score, errors.New(fmt.Sprintf("Bad character input %q in column %d.", char, idx))
		}
	}

	return score, nil
}

func Solution_Day01(part int, input string) (int64, error) {
	if part == 1 {
		return day01_p1(input)
	} else if part == 2 {
		return day01_p2(input)
	} else {
		return 0, errors.New(fmt.Sprintf("Hold on cowboy. No such part (%d) of this day task", part))
	}
}
