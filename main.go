// vim : ts=4 : sw=4 :
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	hbery_aoc2015 "github.com/hbery/aoc-2015-go/src"
)

type ArgStruct struct {
	day  int
	part int
	file string
}

var arguments ArgStruct

func parseArgs() {
	flag.IntVar(&arguments.day, "day", 1, "Choose the day to execute the solution for.")
	flag.IntVar(&arguments.part, "part", 1, "Choose the part of the Day task.")
	flag.StringVar(&arguments.file, "file", "", "Specify input file for the puzzle.")

	flag.Parse()
}

func main() {
	// Parse arguments first
	parseArgs()

	var input string
	if arguments.file != "" {
		if file_contents, err := ioutil.ReadFile(arguments.file); err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else {
			input = string(file_contents)
		}
	} else {
		if len(flag.Args()) < 1 {
			fmt.Println("Error: Specify testcase string or file name using -file option.")
			fmt.Printf("Usage of %s:\n", os.Args[0])
			flag.PrintDefaults()
			os.Exit(1)
		}
		input = flag.Arg(0)
	}

	// Run Solution for certain day
	if err := hbery_aoc2015.Solution(arguments.day, arguments.part, input); err != nil {
		fmt.Println(err)
	}
}
