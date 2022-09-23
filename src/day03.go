package hbery_aoc2015

import (
	"errors"
	"fmt"
	"strings"
)

type Pos struct {
	x, y int64
}

func day03_p1(input string) (int64, error) {
	var visited_houses = make(map[Pos]int64)
	var pos_now = Pos{0, 0}

	visited_houses[pos_now] = 1

	input = strings.TrimSuffix(input, "\n")

	for idx, char := range input {
		switch char {
		case '>':
			pos_now.x++
		case '<':
			pos_now.x--
		case '^':
			pos_now.y++
		case 'v':
			pos_now.y--
		default:
			return -1, errors.New(fmt.Sprintf("Error: Character %q on position %d is not on the list of '^>V<'", char, idx))
		}

		if _, exists := visited_houses[pos_now]; exists {
			visited_houses[pos_now]++
		} else {
			visited_houses[pos_now] = 1
		}
	}

	return int64(len(visited_houses)), nil
}

func day03_p2(input string) (int64, error) {
	var visited_houses = make(map[Pos]int64)
	var pos1_now = Pos{0, 0}
	var pos2_now = Pos{0, 0}

	visited_houses[pos1_now] = 2

	input = strings.TrimSuffix(input, "\n")

	for idx, char := range input {
		var pos_now *Pos
		if idx%2 == 0 {
			pos_now = &pos1_now
		} else {
			pos_now = &pos2_now
		}

		switch char {
		case '>':
			pos_now.x++
		case '<':
			pos_now.x--
		case '^':
			pos_now.y++
		case 'v':
			pos_now.y--
		default:
			return -1, errors.New(fmt.Sprintf("Error: Character %q on position %d is not on the list of '^>V<'", char, idx))
		}

		if _, exists := visited_houses[*pos_now]; exists {
			visited_houses[*pos_now]++
		} else {
			visited_houses[*pos_now] = 1
		}
	}

	return int64(len(visited_houses)), nil
}

func Solution_Day03(part int, input string) (int64, error) {
	if part == 1 {
		return day03_p1(input)
	} else if part == 2 {
		return day03_p2(input)
	} else {
		return 0, errors.New(fmt.Sprintf("Error: Hold on cowboy. No such part (%d) of this day task", part))
	}
}
