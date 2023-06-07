package hbery_aoc2015

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ELF_PRESENT_MULTIPLIER int = 10
var ELF_PRESENT_NEW_MULTIPLIER int = 11
var ELF_HOUSE_VISIT_LIMIT int = 50

func day20_p1(input string) (int64, error) {
	var result int64 = 0
	var presents_target int64
	var err error

	if presents_target, err = strconv.ParseInt(strings.TrimSuffix(input, "\n"), 10, 64); err != nil {
		return -1, errors.New("Error: Failed parsing input.")
	}

	houses := make([]int, presents_target)

	for elf := 1; elf <= int(presents_target)/ELF_PRESENT_MULTIPLIER; elf++ {
		boundry := int(presents_target) / elf / ELF_PRESENT_MULTIPLIER

		for num := 1; num <= boundry; num++ {
			houses[num * elf] += elf * ELF_PRESENT_MULTIPLIER
		}
	}

	for i, house := range houses {
		if house >= int(presents_target) {
			result = int64(i)
			break
		}
	}

	return result, nil
}

func day20_p2(input string) (int64, error) {
	var result int64 = 0
	var presents_target int64
	var err error

	if presents_target, err = strconv.ParseInt(strings.TrimSuffix(input, "\n"), 10, 64); err != nil {
		return -1, errors.New("Error: Failed parsing input.")
	}

	houses := make([]int, presents_target)

	for elf := 1; elf < int(int(presents_target)/ELF_PRESENT_NEW_MULTIPLIER); elf++ {
		boundry := int(presents_target) / elf / ELF_PRESENT_MULTIPLIER

		if boundry > ELF_HOUSE_VISIT_LIMIT {
			boundry = ELF_HOUSE_VISIT_LIMIT
		}

		for num := 1; num <= boundry; num++ {
			houses[num * elf] += elf * ELF_PRESENT_NEW_MULTIPLIER
		}
	}

	for i, house := range houses {
		if house >= int(presents_target) {
			result = int64(i)
			break
		}
	}

	return result, nil
}

func Solution_Day20(part int, input string) (int64, error) {
	if part == 1 {
		return day20_p1(input)
	} else if part == 2 {
		return day20_p2(input)
	} else {
		return 0, errors.New(fmt.Sprintf("Error: Hold on cowboy. No such part (%d) of this day task", part))
	}
}
