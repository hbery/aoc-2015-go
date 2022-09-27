package hbery_aoc2015

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func re_toI(s string) int {
	val, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return int(val)
}

func re_init_grid(grid *[][]int, size_m, size_n int) {
	for i := 0; i < size_m; i++ {
		for j := 0; j < size_n; j++ {
			(*grid)[i][j] = 0
		}
	}
}

func re_turn_on_grid(grid *[][]int, s []string) {
	for i := 0; i < len(*grid); i++ {
		for j := 0; j < len(*grid); j++ {
			if i >= re_toI(s[0]) && j >= re_toI(s[1]) && i <= re_toI(s[2]) && j <= re_toI(s[3]) {
				(*grid)[i][j] = 1
			}
		}
	}
}

func re_turn_on_grid_b(grid *[][]int, s []string) {
	for i := 0; i < len(*grid); i++ {
		for j := 0; j < len(*grid); j++ {
			if i >= re_toI(s[0]) && j >= re_toI(s[1]) && i <= re_toI(s[2]) && j <= re_toI(s[3]) {
				(*grid)[i][j] += 1
			}
		}
	}
}

func re_turn_off_grid(grid *[][]int, s []string) {
	for i := 0; i < len(*grid); i++ {
		for j := 0; j < len(*grid); j++ {
			if i >= re_toI(s[0]) && j >= re_toI(s[1]) && i <= re_toI(s[2]) && j <= re_toI(s[3]) {
				(*grid)[i][j] = 0
			}
		}
	}
}

func re_turn_off_grid_b(grid *[][]int, s []string) {
	for i := 0; i < len(*grid); i++ {
		for j := 0; j < len(*grid); j++ {
			if i >= re_toI(s[0]) && j >= re_toI(s[1]) && i <= re_toI(s[2]) && j <= re_toI(s[3]) {
				if (*grid)[i][j] != 0 {
					(*grid)[i][j] -= 1
				}
			}
		}
	}
}

func re_toggle_grid(grid *[][]int, s []string) {
	for i := 0; i < len(*grid); i++ {
		for j := 0; j < len(*grid); j++ {
			if i >= re_toI(s[0]) && j >= re_toI(s[1]) && i <= re_toI(s[2]) && j <= re_toI(s[3]) {
				(*grid)[i][j] = ((*grid)[i][j] + 1) % 2
			}
		}
	}
}

func re_toggle_grid_b(grid *[][]int, s []string) {
	for i := 0; i < len(*grid); i++ {
		for j := 0; j < len(*grid); j++ {
			if i >= re_toI(s[0]) && j >= re_toI(s[1]) && i <= re_toI(s[2]) && j <= re_toI(s[3]) {
				(*grid)[i][j] += 2
			}
		}
	}
}

func re_count_ones(grid *[][]int) int64 {
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

func re_count_brightness(grid *[][]int) int64 {
	var light int64 = 0

	for i := 0; i < len(*grid); i++ {
		for j := 0; j < len(*grid); j++ {
			light += int64((*grid)[i][j])
		}
	}

	return light
}

func day06_p1_regex(input string) (int64, error) {
	// Init grid of lights
	var grid [][]int = make([][]int, 1000)
	for row := 0; row < 1000; row++ {
		grid[row] = make([]int, 1000)
	}
	init_grid(&grid, 1000, 1000)

	// turn on|turn off|toggle start_x,start_y through end_x,end_y
	re := regexp.MustCompile(`(turn on|turn off|toggle) (\d{1,3}),(\d{1,3}) through (\d{1,3}),(\d{1,3})`)

	input = strings.TrimSuffix(input, "\n")

	for _, line := range strings.Split(input, "\n") {
		refind := re.FindStringSubmatch(line)

		if refind[1] == "turn on" {
			re_turn_on_grid(&grid, refind[2:])
		}

		if refind[1] == "turn off" {
			re_turn_off_grid(&grid, refind[2:])
		}

		if refind[1] == "toggle" {
			re_toggle_grid(&grid, refind[2:])
		}
	}

	return re_count_ones(&grid), nil
}

func day06_p2_regex(input string) (int64, error) {
	// Init grid of lights
	var grid [][]int = make([][]int, 1000)
	for row := 0; row < 1000; row++ {
		grid[row] = make([]int, 1000)
	}
	init_grid(&grid, 1000, 1000)

	// turn on|turn off|toggle start_x,start_y through end_x,end_y
	re := regexp.MustCompile(`(turn on|turn off|toggle) (\d{1,3}),(\d{1,3}) through (\d{1,3}),(\d{1,3})`)

	input = strings.TrimSuffix(input, "\n")

	for _, line := range strings.Split(input, "\n") {
		refind := re.FindStringSubmatch(line)

		if refind[1] == "turn on" {
			re_turn_on_grid_b(&grid, refind[2:])
		}

		if refind[1] == "turn off" {
			re_turn_off_grid_b(&grid, refind[2:])
		}

		if refind[1] == "toggle" {
			re_toggle_grid_b(&grid, refind[2:])
		}
	}

	return re_count_brightness(&grid), nil
}

func Solution_Day06_regex(part int, input string) (int64, error) {
	if part == 1 {
		return day06_p1_regex(input)
	} else if part == 2 {
		return day06_p2_regex(input)
	} else {
		return 0, errors.New(fmt.Sprintf("Error: Hold on cowboy. No such part (%d) of this day task", part))
	}
}
