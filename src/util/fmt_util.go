package hbery_aoc2015

import (
	"fmt"
	"strings"
)

func PrintlnCenter(text string, width int, border string) {
	text_width := len(text)
	var to_pad int
	if (text_width%2 == 0 && width%2 == 1) || (text_width%2 == 1 && width%2 == 0) {
		text = fmt.Sprintf("%s ", text)
		to_pad = (width - text_width - 2*len(border) - 2 - 1) / 2
	} else {
		to_pad = (width - text_width - 2*len(border) - 2) / 2
	}

	fmt.Printf("%s%s %s %s%s\n",
		border, fmt.Sprintf(fmt.Sprintf("%%%ds", to_pad), " "),
		text,
		fmt.Sprintf(fmt.Sprintf("%%%ds", to_pad), " "), border,
	)
}

func PrintlnCenterAndPad(text string, width int, pad_char string) {
	text_width := len(text)
	var to_pad int
	if (text_width%2 == 0 && width%2 == 1) || (text_width%2 == 1 && width%2 == 0) {
		text = fmt.Sprintf("%s ", text)
		to_pad = (width - text_width - 2 - 1) / 2
	} else {
		to_pad = (width - text_width - 2) / 2
	}

	fmt.Printf("%s %s %s\n",
		strings.Repeat(pad_char, to_pad),
		text,
		strings.Repeat(pad_char, to_pad),
	)
}
