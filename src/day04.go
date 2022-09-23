package hbery_aoc2015

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"strings"
)

func day04_p1(input string) (int64, error) {
	var number int64 = 1

	input = strings.TrimSuffix(input, "\n")

	for {
		digest := md5.New()
		io.WriteString(digest, fmt.Sprintf("%s%d", input, number))
		hash := digest.Sum(nil)

		if hash[0] == 0 && hash[1] == 0 && hash[2] < 16 {
			break
		}

		number++
	}

	return number, nil
}

func day04_p2(input string) (int64, error) {
	return -1, nil
}

func Solution_Day04(part int, input string) (int64, error) {
	if part == 1 {
		return day04_p1(input)
	} else if part == 2 {
		return day04_p2(input)
	} else {
		return 0, errors.New(fmt.Sprintf("Error: Hold on cowboy. No such part (%d) of this day task", part))
	}
}
