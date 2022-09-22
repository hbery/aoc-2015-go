package hbery_aoc2015

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func min_array(arr []int64) int64 {
	var min_element int64 = arr[0]
	for _, element := range arr {
		if element < min_element {
			min_element = element
		}
	}

	return min_element
}

func day02_p1(input string) (int64, error) {
	var surface int64 = 0

	for idx, line := range strings.Split(strings.TrimSuffix(input, "\n"), "\n") {
		var dimensions []string = strings.Split(line, "x")
		var length, width, height int64
		var err error

		if length, err = strconv.ParseInt(dimensions[0], 10, 64); err != nil {
			return -1, errors.New(fmt.Sprintf("Error: parsing length string value to int64 failed at %d. line: %s", idx, line))
		}

		if width, err = strconv.ParseInt(dimensions[1], 10, 64); err != nil {
			return -1, errors.New(fmt.Sprintf("Error: parsing width string value to int64 failed at %d. line: %s", idx, line))
		}

		if height, err = strconv.ParseInt(dimensions[2], 10, 64); err != nil {
			return -1, errors.New(fmt.Sprintf("Error: parsing height string value to int64 failed at %d. line: %s", idx, line))
		}

		side_lw := length * width
		side_wh := width * height
		side_hl := height * length
		surface += 2*side_lw + 2*side_wh + 2*side_hl + min_array([]int64{side_lw, side_wh, side_hl})
	}

	return surface, nil
}

func day02_p2(input string) (int64, error) {
	return -1, nil
}

func Solution_Day02(part int, input string) (int64, error) {
	if part == 1 {
		return day02_p1(input)
	} else if part == 2 {
		return day02_p2(input)
	} else {
		return 0, errors.New(fmt.Sprintf("Error: Hold on cowboy. No such part (%d) of this day task", part))
	}
}
