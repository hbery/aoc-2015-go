// vim : ts=4 : sw=4 :
package main

import (
	"flag"

	hbery_aoc2015 "github.com/hbery/aoc-2015-go/src"
)

type Args struct {
	day int
}

var arguments Args

func parseArgs() {
	flag.IntVar(&arguments.day, "day", 1, "Choose the day to execute the solution for.")

	flag.Parse()
}

func main() {
	// Parse arguments first
	parseArgs()

	// Run Solution for certain day
	hbery_aoc2015.Solution(arguments.day)
}
