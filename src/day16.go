package hbery_aoc2015

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Aunt struct {
	num			int64
	props		map[string]int64
	commons		int64
}

func parse_aunts(input string) ([]*Aunt, error) {
	var aunts []*Aunt
	var parsed_num int64
	var err error

	for a, line := range strings.Split(input, "\n") {
		var aunt *Aunt = &Aunt{
			num: int64(a+1),
			props: make(map[string]int64),
			commons: 0,
		}
		ld := strings.Split(line, " ")

		// skip Aunt Sue number
		for i := 2; i < len(ld); i++ {
			category := strings.TrimSuffix(ld[i], ":")
			i++
			possible_num := strings.TrimSuffix(ld[i], ",")

			if parsed_num, err = strconv.ParseInt(possible_num, 10, 64); err != nil {
				return nil, errors.New(fmt.Sprintf("Error: Failed to parse number %s for aunt %d.", possible_num, a))
			}

			aunt.props[category] = parsed_num
		}
		aunts = append(aunts, aunt)
	}

	return aunts, nil
}

func day16_p1(input string) (int64, error) {
	var result int64 = 0
	var aunts []*Aunt
	var err error

	var mfcsam_aunt *Aunt = &Aunt{
		props: map[string]int64{
			"children": 3,
			"cats": 7,
			"samoyeds": 2,
			"pomeranians": 3,
			"akitas": 0,
			"vizslas": 0,
			"goldfish": 5,
			"trees": 3,
			"cars": 2,
			"perfumes": 1,
		},
	}

	input = strings.TrimSuffix(input, "\n")
	if aunts, err = parse_aunts(input); err != nil {
		return -1, err
	}

	for _, a := range aunts {
		for p,n := range mfcsam_aunt.props {
			if val, ok := a.props[p]; ok {
				if val == n {
					a.commons++
				}
			}
		}
		
		if a.commons == 3 { result = a.num }
	}

	return result, nil
}

func day16_p2(input string) (int64, error) {
	var result int64 = 0
	var aunts []*Aunt
	var err error

	var mfcsam_aunt *Aunt = &Aunt{
		props: map[string]int64{
			"children": 3,
			"cats": 7,
			"samoyeds": 2,
			"pomeranians": 3,
			"akitas": 0,
			"vizslas": 0,
			"goldfish": 5,
			"trees": 3,
			"cars": 2,
			"perfumes": 1,
		},
	}

	input = strings.TrimSuffix(input, "\n")
	if aunts, err = parse_aunts(input); err != nil {
		return -1, err
	}

	for _, a := range aunts {
		for p,n := range mfcsam_aunt.props {
			if val, ok := a.props[p]; ok {
				if p == "cats" || p == "trees" {
					if val > n {
						a.commons++
					}
				} else if p == "pomeranians" || p == "goldfish" {
					if val < n {
						a.commons++
					}
				} else if val == n {
					a.commons++
				}
			}
		}
		
		if a.commons == 3 { result = a.num }
	}

	return result, nil
}

func Solution_Day16(part int, input string) (int64, error) {
	if part == 1 {
		return day16_p1(input)
	} else if part == 2 {
		return day16_p2(input)
	} else {
		return 0, errors.New(fmt.Sprintf("Error: Hold on cowboy. No such part (%d) of this day task", part))
	}
}
