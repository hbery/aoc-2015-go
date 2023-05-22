package hbery_aoc2015

import (
	"errors"
	"fmt"
	"strings"
)

func parse_light_matrix(input string, msize int) *[][]int {
	var light_matrix [][]int = make([][]int, msize)

	for n, line := range strings.Split(input, "\n") {
		var row []int = make([]int, msize)
		for k, c := range strings.Split(line, "") {
			if c == "#" {
				row[k] = 1
			} else {
				row[k] = 0
			}
		}
		light_matrix[n] = row
	}

	return &light_matrix
}

func prep_next_light_cycle(light_matrix *[][]int) *[][]int {
	var nsize int = len(*light_matrix)
	var ksize int = len((*light_matrix)[0])
	var new_lmatrix [][]int = make([][]int, len(*light_matrix))

	// * - k ->
	// |
	// n
	// |
	// V
	for n, line := range *light_matrix {
		var row []int = make([]int, len(line))
		for k, c := range line {
			var around_lit int = 0

			if n - 1 < 0 && k - 1 < 0 {
			// upper-left corner
				around_lit = (*light_matrix)[n][k+1] + (*light_matrix)[n+1][k+1] + (*light_matrix)[n+1][k]
			} else if n - 1 < 0 && k + 1 >= ksize {
			// upper-right corner
				around_lit = (*light_matrix)[n][k-1] + (*light_matrix)[n+1][k-1] + (*light_matrix)[n+1][k]
			} else if n + 1 >= nsize && k - 1 < 0 {
			// lower-left corner
				around_lit = (*light_matrix)[n][k+1] + (*light_matrix)[n-1][k+1] + (*light_matrix)[n-1][k]
			} else if n + 1 >= nsize && k + 1 >= ksize {
			// lower-right corner
				around_lit = (*light_matrix)[n][k-1] + (*light_matrix)[n-1][k-1] + (*light_matrix)[n-1][k]
			} else if n - 1 < 0 {
			// upper edge
				around_lit = (*light_matrix)[n][k-1] + (*light_matrix)[n+1][k-1] + (*light_matrix)[n+1][k] + (*light_matrix)[n+1][k+1] + (*light_matrix)[n][k+1]
			} else if k + 1 >= ksize {
			// right edge
				around_lit = (*light_matrix)[n+1][k] + (*light_matrix)[n+1][k-1] + (*light_matrix)[n][k-1] + (*light_matrix)[n-1][k-1] + (*light_matrix)[n-1][k]
			} else if n + 1 >= nsize {
			// lower edge
				around_lit = (*light_matrix)[n][k-1] + (*light_matrix)[n-1][k-1] + (*light_matrix)[n-1][k] + (*light_matrix)[n-1][k+1] + (*light_matrix)[n][k+1]
			} else if k - 1 < 0 {
			// left edge
				around_lit = (*light_matrix)[n-1][k] + (*light_matrix)[n-1][k+1] + (*light_matrix)[n][k+1] + (*light_matrix)[n+1][k+1] + (*light_matrix)[n+1][k]
			} else {
			// rest
				around_lit = (*light_matrix)[n+1][k] + (*light_matrix)[n+1][k-1] + (*light_matrix)[n][k-1] + (*light_matrix)[n-1][k-1] + (*light_matrix)[n-1][k] + (*light_matrix)[n-1][k+1] + (*light_matrix)[n][k+1] + (*light_matrix)[n+1][k+1]
			}

			if c == 1 && (around_lit == 2 || around_lit == 3) {
				row[k] = 1
			} else if c == 0 && around_lit == 3 {
				row[k] = 1
			} else {
				row[k] = 0
			}
		}
		new_lmatrix[n] = row
	}

	return &new_lmatrix
}

func light_back_corners(light_matrix *[][]int, msize *int) {
	(*light_matrix)[0][*msize-1] = 1
	(*light_matrix)[*msize-1][*msize-1] = 1
	(*light_matrix)[*msize-1][0] = 1
	(*light_matrix)[0][0] = 1
}

func count_turned_on_lights(light_matrix *[][]int) int64 {
	var lights_on int = 0

	for _, line := range *light_matrix {
		for _, l := range line {
			lights_on += l
		}
	}
	
	return int64(lights_on)
}

func day18_p1(input string) (int64, error) {
	var result int64 = 0
	var msize = 100
	var cycles int =100
	var light_matrix *[][]int

	light_matrix = parse_light_matrix(strings.TrimSuffix(input, "\n"), msize)

	for i := 0; i < cycles; i++ {
		light_matrix = prep_next_light_cycle(light_matrix)
	}

	result = count_turned_on_lights(light_matrix)

	return result, nil
}

func day18_p2(input string) (int64, error) {
	var result int64 = 0
	var msize = 100
	var cycles int =100
	var light_matrix *[][]int

	light_matrix = parse_light_matrix(strings.TrimSuffix(input, "\n"), msize)

	light_back_corners(light_matrix, &msize)
	for i := 0; i < cycles; i++ {
		light_matrix = prep_next_light_cycle(light_matrix)
		light_back_corners(light_matrix, &msize)
	}

	result = count_turned_on_lights(light_matrix)

	return result, nil
}

func Solution_Day18(part int, input string) (int64, error) {
	if part == 1 {
		return day18_p1(input)
	} else if part == 2 {
		return day18_p2(input)
	} else {
		return 0, errors.New(fmt.Sprintf("Error: Hold on cowboy. No such part (%d) of this day task", part))
	}
}
