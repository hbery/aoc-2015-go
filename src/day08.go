package hbery_aoc2015

import (
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
)

func decode_string(str string) (string, error) {
	var new_str string
	var skip int = 0

	for idx, char := range str {
		switch char {
		case '\\':
			if skip == 1 {
				new_str += string(char)
				skip--
			} else {
				skip++
			}
		case '"':
			if skip == 1 {
				new_str += string(char)
				skip--
			}
		case 'x':
			if skip > 0 {
				hex_str := strings.Join(strings.Split(str, "")[idx+1:idx+3], "")
				if ascii_char, err := hex.DecodeString(hex_str); err != nil {
					return "", err
				} else {
					new_str += string(ascii_char)
				}
				skip++
			} else {
				new_str += string(char)
			}
		default:
			if skip > 0 {
				skip--
			} else {
				new_str += string(char)
			}
		}
	}

	return new_str, nil
}

func encode_string(str string) (string, error) {
	var new_str string

	for _, char := range str {
		switch char {
		case '\\':
			new_str += fmt.Sprintf("%s%s", "\\", string(char))
		case '"':
			new_str += fmt.Sprintf("%s%s", "\\", string(char))
		default:
			new_str += string(char)
		}
	}

	return fmt.Sprintf("%s%s%s", "\"", new_str, "\""), nil
}

func day08_p1(input string) (int64, error) {
	var num_chars, num_code_literals int

	input = strings.TrimSuffix(input, "\n")

	for _, line := range strings.Split(input, "\n") {
		num_code_literals += len(line)
		// fmt.Println("+ ", line)

		if str, err := decode_string(line); err != nil {
			panic(err)
		} else {
			num_chars += len(str)
			// fmt.Println("- ", str)
		}
	}

	return int64(num_code_literals - num_chars), nil
}

func day08_p2(input string) (int64, error) {
	var num_escaped_chars, num_code_literals int

	input = strings.TrimSuffix(input, "\n")

	for _, line := range strings.Split(input, "\n") {
		num_code_literals += len(line)
		// fmt.Println("+ ", line)

		if str, err := encode_string(line); err != nil {
			panic(err)
		} else {
			num_escaped_chars += len(str)
			// fmt.Println("- ", str)
		}
	}

	return int64(num_escaped_chars - num_code_literals), nil
}

func Solution_Day08(part int, input string) (int64, error) {
	if part == 1 {
		return day08_p1(input)
	} else if part == 2 {
		return day08_p2(input)
	} else {
		return 0, errors.New(fmt.Sprintf("Error: Hold on cowboy. No such part (%d) of this day task", part))
	}
}
