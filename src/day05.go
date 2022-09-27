package hbery_aoc2015

import (
	"errors"
	"fmt"
	"strings"
)

func str_contains_3chars_of(s string, chars string) bool {
	ct := 3

	for ct != 0 {
		if idx := strings.IndexAny(s, chars); idx != -1 {
			s = strings.Replace(s, string(s[idx]), "", 1)
			ct--
		} else {
			return false
		}
	}

	return true
}

func str_contains_any_letter_twice_in_row(s []rune) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			return true
		}
	}

	return false
}

func str_contains_one_of(s string, phrases []string) bool {
	for _, p := range phrases {
		if strings.Contains(s, p) {
			return true
		}
	}
	return false
}

func str_contains_pair_of_2letters(s []rune) bool {
	for i := 0; i < len(s)-2; i++ {
		if strings.Count(string(s), string(s[i:i+2])) >= 2 {
			return true
		}
	}

	return false
}

func str_contains_same_char_letter_apart(s []rune) bool {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] {
			return true
		}
	}

	return false
}

func day05_p1(input string) (int64, error) {
	var nice_strings_num int64 = 0

	input = strings.TrimSuffix(input, "\n")

	for _, str := range strings.Split(input, "\n") {
		// Check if contains at least 3 vowels
		if !str_contains_3chars_of(str, "aeiou") {
			continue
		}

		// Check if contains any letter that appears twice in a row
		if !str_contains_any_letter_twice_in_row([]rune(str)) {
			continue
		}

		// Check for forbidden strings
		if str_contains_one_of(str, []string{"ab", "cd", "pq", "xy"}) {
			continue
		}

		nice_strings_num++
	}

	return nice_strings_num, nil
}

func day05_p2(input string) (int64, error) {
	var nice_strings_num int64 = 0

	input = strings.TrimSuffix(input, "\n")

	for _, str := range strings.Split(input, "\n") {
		// Check for pair presented at least twice non-overlapping
		if !str_contains_pair_of_2letters([]rune(str)) {
			continue
		}

		// Check for repeated letter split by 1 letter apart
		if !str_contains_same_char_letter_apart([]rune(str)) {
			continue
		}

		nice_strings_num++
	}

	return nice_strings_num, nil
}

func Solution_Day05(part int, input string) (int64, error) {
	if part == 1 {
		return day05_p1(input)
	} else if part == 2 {
		return day05_p2(input)
	} else {
		return 0, errors.New(fmt.Sprintf("Error: Hold on cowboy. No such part (%d) of this day task", part))
	}
}
