package hbery_aoc2015

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func json_split(r rune) bool {
	return r == '{' || r == '}' || r == '[' || r == ']' || r == ',' || r == ':'
}

func dummy_json_nums_counter(input string) int64 {
	var json_elements []string
	var result int64 = 0
	var int_el int64
	var err error

	json_elements = strings.FieldsFunc(input, json_split)

	for _, el := range json_elements {
		if int_el, err = strconv.ParseInt(el, 10, 64); err != nil {
			continue
		}

		result += int_el
	}

	return result
}

func sum_numbers_in_array(input_json []interface{}, skip_red bool) int64 {
	var result int64 = 0

	for _, v := range input_json {
		switch v.(type) {
		case map[string]interface{}:
			result += sum_numbers_in_map(v.(map[string]interface{}), skip_red)
		case []interface{}:
			result += sum_numbers_in_array(v.([]interface{}), skip_red)
		case float64:
			result += int64(v.(float64))
		}
	}

	return result
}

func sum_numbers_in_map(input_json map[string]interface{}, skip_red bool) int64 {
	var result int64 = 0

	for _, v := range input_json {
		switch v.(type) {
		case map[string]interface{}:
			result += sum_numbers_in_map(v.(map[string]interface{}), skip_red)
		case []interface{}:
			result += sum_numbers_in_array(v.([]interface{}), skip_red)
		case string:
			if skip_red && v == "red" {
				return 0
			}
		case float64:
			result += int64(v.(float64))
		}
	}

	return result
}

func json_count(input string, skip_red bool) (int64, error) {
	var json_structure map[string]interface{}
	var result int64 = 0

	if err := json.Unmarshal([]byte(input), &json_structure); err != nil {
		return -1, err
	}

	result = sum_numbers_in_map(json_structure, skip_red)

	return result, nil
}

func day12_p1(input string) (int64, error) {
	var result int64 = 0
	var count_case string = "json"
	var err error

	input = strings.TrimSuffix(input, "\n")

	switch count_case {
	case "dummy":
		result = dummy_json_nums_counter(input)
	case "json":
		if result, err = json_count(input, false); err != nil {
			return result, err
		}
	}

	return result, nil
}

func day12_p2(input string) (int64, error) {
	var result int64 = 0
	var err error

	input = strings.TrimSuffix(input, "\n")

	if result, err = json_count(input, true); err != nil {
		return result, err
	}

	return result, nil
}

func Solution_Day12(part int, input string) (int64, error) {
	if part == 1 {
		return day12_p1(input)
	} else if part == 2 {
		return day12_p2(input)
	} else {
		return 0, errors.New(fmt.Sprintf("Error: Hold on cowboy. No such part (%d) of this day task", part))
	}
}
