package hbery_aoc2015

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// too slow..
func look_and_say_strings(input string) (string, error) {
	var new_string string
	var currnum string = ""
	var numcount int = 0

	for _, num := range strings.Split(input, "") {
		if num == currnum {
			numcount++
		} else {
			if currnum == "" {
				currnum = num
				numcount++
				continue
			}

			new_string = fmt.Sprintf("%s%d%s", new_string, numcount, currnum)
			currnum = num
			numcount = 1
		}
	}

	// last num scenario
	new_string = fmt.Sprintf("%s%d%s", new_string, numcount, currnum)

	return new_string, nil
}

func look_and_say_bytes(input []byte) ([]byte, error) {
	var new_numbers []byte

	for i := 0; i < len(input); i++ {
		start_idx := i
		for i < len(input)-1 && input[i+1] == input[i] {
			i++
		}
		end_idx := i

		numcount := end_idx - start_idx + 1

		new_numbers = append(new_numbers, strconv.Itoa(numcount)[0], input[i])
	}

	return new_numbers, nil
}

func day10_p1(input string) (int64, error) {
	var result []byte
	var repetition int = 40
	var err error

	result = []byte(strings.TrimSuffix(input, "\n"))
	
	for i := 0; i < repetition; i++ {
		if result, err = look_and_say_bytes(result); err != nil {
			return -1, errors.New(fmt.Sprintf("Error: new creation failed at rep %d at input string: %s", i, result))
		}
	}

	return int64(len(result)), nil
}

func day10_p2(input string) (int64, error) {
	var result []byte
	var repetition int = 50
	var err error

	result = []byte(strings.TrimSuffix(input, "\n"))
	
	for i := 0; i < repetition; i++ {
		if result, err = look_and_say_bytes(result); err != nil {
			return -1, errors.New(fmt.Sprintf("Error: new creation failed at rep %d at input string: %s", i, result))
		}
	}

	return int64(len(result)), nil
}

func Solution_Day10(part int, input string) (int64, error) {
	if part == 1 {
		return day10_p1(input)
	} else if part == 2 {
		return day10_p2(input)
	} else {
		return 0, errors.New(fmt.Sprintf("Error: Hold on cowboy. No such part (%d) of this day task", part))
	}
}
