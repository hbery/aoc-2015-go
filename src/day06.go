package hbery_aoc2015

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func toI(s string) int {
	val, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return int(val)
}

func init_grid(grid *[][]int, size_m, size_n int) {
	for i := 0; i < size_m; i++ {
		for j := 0; j < size_n; j++ {
			(*grid)[i][j] = 0
		}
	}
}

func turn_on_grid(grid *[][]int, line string) {
	s := strings.Split(line, " ")
	start := strings.Split(s[2], ",")
	end := strings.Split(s[4], ",")

	for i := 0; i < len(*grid); i++ {
		for j := 0; j < len(*grid); j++ {
			if i >= toI(start[0]) && j >= toI(start[1]) && i <= toI(end[0]) && j <= toI(end[1]) {
				(*grid)[i][j] = 1
			}
		}
	}
}

func turn_off_grid(grid *[][]int, line string) {
	s := strings.Split(line, " ")
	start := strings.Split(s[2], ",")
	end := strings.Split(s[4], ",")

	for i := 0; i < len(*grid); i++ {
		for j := 0; j < len(*grid); j++ {
			if i >= toI(start[0]) && j >= toI(start[1]) && i <= toI(end[0]) && j <= toI(end[1]) {
				(*grid)[i][j] = 0
			}
		}
	}
}

func toggle_grid(grid *[][]int, line string) {
	s := strings.Split(line, " ")
	start := strings.Split(s[1], ",")
	end := strings.Split(s[3], ",")

	for i := 0; i < len(*grid); i++ {
		for j := 0; j < len(*grid); j++ {
			if i >= toI(start[0]) && j >= toI(start[1]) && i <= toI(end[0]) && j <= toI(end[1]) {
				(*grid)[i][j] = ((*grid)[i][j] + 1) % 2
			}
		}
	}
}

func count_ones(grid *[][]int) int64 {
	var ones int64 = 0

	for i := 0; i < len(*grid); i++ {
		for j := 0; j < len(*grid); j++ {
			if (*grid)[i][j] == 1 {
				ones++
			}
		}
	}

	return ones
}

func day06_p1(input string) (int64, error) {
	// Init grid of lights
	var grid [][]int = make([][]int, 1000)
	for row := 0; row < 1000; row++ {
		grid[row] = make([]int, 1000)
	}
	init_grid(&grid, 1000, 1000)

	// turn on start_x,start_y through end_x,end_y
	re_1 := regexp.MustCompile(`turn on`)
	// turn off start_x,start_y through end_x,end_y
	re_0 := regexp.MustCompile(`turn of`)
	// toggle start_x,start_y through end_x,end_y
	re_t := regexp.MustCompile(`toggle`)

	input = strings.TrimSuffix(input, "\n")

	for _, line := range strings.Split(input, "\n") {
		if re_1.MatchString(line) {
			turn_on_grid(&grid, line)
		}

		if re_0.MatchString(line) {
			turn_off_grid(&grid, line)
		}

		if re_t.MatchString(line) {
			toggle_grid(&grid, line)
		}
	}

	return count_ones(&grid), nil
}

func day06_p2(input string) (int64, error) {
	return -1, nil
}

func Solution_Day06(part int, input string) (int64, error) {
	if part == 1 {
		return day06_p1(input)
	} else if part == 2 {
		return day06_p2(input)
	} else {
		return 0, errors.New(fmt.Sprintf("Error: Hold on cowboy. No such part (%d) of this day task", part))
	}
}
