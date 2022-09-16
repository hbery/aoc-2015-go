// vim : ts=4 : sw=4 :
package main

import (
	"flag"
	"fmt"
)

type Args struct {
	day int
}

var arguments Args

func parseArgs() {
	flag.IntVar(&arguments.day, "day", 1, "Choose the day to execute.")

	flag.Parse()
}

func main() {
	fmt.Println("========== Advent of Code 2015 ==========")
	parseArgs()
	fmt.Printf("Day: %d\n", arguments.day)
}
