package hbery_aoc2015

import (
	"errors"
	"fmt"
	"strings"

	"github.com/dlclark/regexp2"
)

func re_contains_3chars_of(s string, chars string) bool {
	isMatch, _ := regexp2.MustCompile(fmt.Sprintf("[%s].*[%s].*[%s]", chars, chars, chars), 0).MatchString(s)
	return isMatch
}

func re_contains_any_letter_twice_in_row(s string) bool {
	isMatch, _ := regexp2.MustCompile(`([a-z])\1`, 0).MatchString(s)
	return isMatch
}

func re_contains_one_of(s string, phrases []string) bool {
	isMatch, _ := regexp2.MustCompile(fmt.Sprintf("%s", strings.Join(phrases, "|")), 0).MatchString(s)
	return isMatch
}

func re_contains_pair_of_2letters(s string) bool {
	isMatch, _ := regexp2.MustCompile(`([a-z][a-z]).*\1`, 0).MatchString(s)
	return isMatch
}

func re_contains_same_char_letter_apart(s string) bool {
	isMatch, _ := regexp2.MustCompile(`([a-z])[a-z]\1`, 0).MatchString(s)
	return isMatch
}

func day05_p1_regex(input string) (int64, error) {
	var nice_strings_num int64 = 0

	input = strings.TrimSuffix(input, "\n")

	for _, str := range strings.Split(input, "\n") {
		// Check if contains at least 3 vowels
		if !re_contains_3chars_of(str, "aeiou") {
			continue
		}

		// Check if contains any letter that appears twice in a row
		if !re_contains_any_letter_twice_in_row(str) {
			continue
		}

		// Check for forbidden strings
		if re_contains_one_of(str, []string{"ab", "cd", "pq", "xy"}) {
			continue
		}

		nice_strings_num++
	}

	return nice_strings_num, nil
}

func day05_p2_regex(input string) (int64, error) {
	var nice_strings_num int64 = 0

	input = strings.TrimSuffix(input, "\n")

	for _, str := range strings.Split(input, "\n") {
		// Check for pair presented at least twice non-overlapping
		if !re_contains_pair_of_2letters(str) {
			continue
		}

		// Check for repeated letter split by 1 letter apart
		if !re_contains_same_char_letter_apart(str) {
			continue
		}

		nice_strings_num++
	}

	return nice_strings_num, nil
}

func Solution_Day05_regex(part int, input string) (int64, error) {
	if part == 1 {
		return day05_p1_regex(input)
	} else if part == 2 {
		return day05_p2_regex(input)
	} else {
		return 0, errors.New(fmt.Sprintf("Error: Hold on cowboy. No such part (%d) of this day task", part))
	}
}
