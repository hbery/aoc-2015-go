package hbery_aoc2015

import (
	"errors"
	"fmt"
	"strings"

	hbery_aoc2015 "github.com/hbery/aoc-2015-go/src/util"
)

func check_three_consecutive(pass []byte) bool {
	for i := 0; i < len(pass) - 2; i++ {
		if pass[i] == pass[i+1]-1 && pass[i] == pass[i+2]-2 {
			return true
		}
	}

	return false
}

func check_no_forbidden(pass []byte) bool {
	for _, p := range pass {
		if p == 'i' || p == 'o' || p == 'l' {
			return false
		}
	}

	return true
}

func check_two_pairs(pass []byte) bool {
	var pairs int = 0

	for i := 0; i < len(pass) - 1; i++ {
		if pass[i] == pass[i+1] {
			pairs++ // increment pair
			i++ // also skip one letter to avoid overlap
		}

		if pairs == 2 {
			return true
		}
	}

	return false
}

func verify_password(pass []byte) bool {
	if ! check_three_consecutive(pass) { return false }
    if ! check_no_forbidden(pass) { return false }
	if ! check_two_pairs(pass) { return false }

	return true
}

func increment_password(pass []byte) []byte {
	for i := len(pass) - 1; i >= 0; i-- {
		tmp := pass[i] + 1
		if tmp >= 'a' && tmp <= 'z' {
			pass[i] = tmp
			break
		}

		pass[i] = 'a'
	}

	return pass
}

func find_password(input_pass []byte) []byte {
	var current_pass []byte = input_pass

	current_pass = increment_password(current_pass)
	for next := true; next; next = !verify_password(current_pass) {
		current_pass = increment_password(current_pass)
	}

	return current_pass
}

func day11_p1(input string) (int64, error) {
	var result []byte

	result = find_password([]byte(strings.TrimSuffix(input, "\n")))

	hbery_aoc2015.PrintlnCenter(fmt.Sprintf("Result: %s", result), 80, "")
	return -1, nil
}

func day11_p2(input string) (int64, error) {
	var result []byte

	result = find_password([]byte(strings.TrimSuffix(input, "\n")))
	result = find_password(result)

	hbery_aoc2015.PrintlnCenter(fmt.Sprintf("Result: %s", result), 80, "")
	return -1, nil
}

func Solution_Day11(part int, input string) (int64, error) {
	if part == 1 {
		return day11_p1(input)
	} else if part == 2 {
		return day11_p2(input)
	} else {
		return 0, errors.New(fmt.Sprintf("Error: Hold on cowboy. No such part (%d) of this day task", part))
	}
}
