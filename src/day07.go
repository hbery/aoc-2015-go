package hbery_aoc2015

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Wire struct {
	op     string
	signal uint16
	in     []string
	arg    uint16
}

type Circuit struct {
	resolved bool
	scheme   map[string]*Wire
}

func make_circuit() *Circuit {
	new_circuit := new(Circuit)
	new_circuit.scheme = make(map[string]*Wire)
	new_circuit.resolved = false
	return new_circuit
}

func (c *Circuit) resolveCircuit() {
	if c.resolved {
		return
	}

	var nodes []string
	for k := range c.scheme {
		nodes = append(nodes, k)
	}

	for !c.allNodesDone() {
		for _, w := range nodes {
			// Check if node already isolated
			if c.scheme[w].in == nil {
				continue
			}

			var args []uint16
			// Check if sources already isolated from their sources
			good2op := true
			for _, s := range c.scheme[w].in {
				if val, err := strconv.ParseUint(s, 10, 16); err == nil {
					args = append(args, uint16(val))
					continue
				}

				if c.scheme[s].in != nil {
					good2op = false
					break
				} else {
					args = append(args, c.scheme[s].signal)
				}
			}

			if !good2op {
				continue
			}

			if c.scheme[w].arg != 0 {
				args = append(args, c.scheme[w].arg)
			}

			// Perform operation if it is set
			c.doOp(w, args)
		}
	}
	c.resolved = true
}

func (c *Circuit) allNodesDone() bool {
	// Check if all nodes isolated
	for _, w := range c.scheme {
		if w.in != nil {
			return false
		}
	}

	return true
}

func (c *Circuit) doOp(w string, args []uint16) {
	// Perform operation
	switch c.scheme[w].op {
	case "NOT":
		c.scheme[w].signal = ^args[0]
	case "AND":
		c.scheme[w].signal = args[0] & args[1]
	case "OR":
		c.scheme[w].signal = args[0] | args[1]
	case "LSHIFT":
		c.scheme[w].signal = args[0] << args[1]
	case "RSHIFT":
		c.scheme[w].signal = args[0] >> args[1]
	default:
		c.scheme[w].signal = args[0]
	}

	// fmt.Printf("%s: %+v -> %s :: %d\n", c.scheme[w].op, c.scheme[w].in, w, c.scheme[w].signal)

	// Isolate node from sources
	c.scheme[w].in = nil
}

func parse_input(c *Circuit, input string) error {
	// --- REGEX PHRASES COMPILATION --- //
	re_wire := regexp.MustCompile(`^(\d{1,5}|[a-z]{1,2}) -> (\d{1,5}|[a-z]{1,2})`)
	re_not := regexp.MustCompile(`^(NOT) (\d{1,5}|[a-z]{1,2}) -> (\d{1,5}|[a-z]{1,2})`)
	re_andor := regexp.MustCompile(`^(\d{1,5}|[a-z]{1,2}) (AND|OR) (\d{1,5}|[a-z]{1,2}) -> (\d{1,5}|[a-z]{1,2})`)
	re_lrshift := regexp.MustCompile(`^(\d{1,5}|[a-z]{1,2}) (LSHIFT|RSHIFT) (\d{1,2}) -> (\d{1,5}|[a-z]{1,2})`)

	for _, line := range strings.Split(input, "\n") {
		wire := new(Wire)
		if match := re_wire.FindStringSubmatch(line); match != nil {
			// fmt.Printf("op: , sources: %s, destination: %s\n", match[1], match[2])
			if val, err := strconv.ParseUint(match[1], 10, 16); err != nil {
				wire.in = append(wire.in, match[1])
			} else {
				wire.signal = uint16(val)
				wire.in = nil
			}

			fmt.Printf("%s :: %+v\n", match[2], wire)
			c.scheme[match[2]] = wire
			continue
		}

		if match := re_not.FindStringSubmatch(line); match != nil {
			// fmt.Printf("op: %s, sources: %s, destination: %s\n", match[1], match[2], match[3])
			wire.op = match[1]
			wire.in = append(wire.in, match[2])

			fmt.Printf("%s :: %+v\n", match[3], wire)
			c.scheme[match[3]] = wire
			continue
		}

		if match := re_andor.FindStringSubmatch(line); match != nil {
			// fmt.Printf("op: %s, sources: %s %s, destination: %s\n", match[2], match[1], match[3], match[4])
			wire.op = match[2]
			wire.in = append(wire.in, match[1], match[3])

			fmt.Printf("%s :: %+v\n", match[4], wire)
			c.scheme[match[4]] = wire
			continue
		}

		if match := re_lrshift.FindStringSubmatch(line); match != nil {
			// fmt.Printf("op: %s, sources: %s, destination: %s, arg: %s\n", match[2], match[1], match[4], match[3])
			wire.op = match[2]
			wire.in = append(wire.in, match[1])
			if val, err := strconv.ParseUint(match[3], 10, 16); err != nil {
				return errors.New("Error: Cannot parse int from bits string.")
			} else {
				wire.arg = uint16(val)
			}

			fmt.Printf("%s :: %+v\n", match[4], wire)
			c.scheme[match[4]] = wire
			continue
		}
	}

	return nil
}

func day07_p1(input string) (int64, error) {
	this_circuit := make_circuit()

	input = strings.TrimSuffix(input, "\n")

	if err := parse_input(this_circuit, input); err != nil {
		return -1, err
	}

	this_circuit.resolveCircuit()

	for k, v := range this_circuit.scheme {
		fmt.Println(k, " => ", v.signal)
	}

	if val, exists := this_circuit.scheme["a"]; !exists {
		return -1, errors.New("Error: There is no key `a` in the sources map.")
	} else {
		return int64(val.signal), nil
	}

}

func day07_p2(input string) (int64, error) {
	return -1, nil
}

func Solution_Day07(part int, input string) (int64, error) {
	if part == 1 {
		return day07_p1(input)
	} else if part == 2 {
		return day07_p2(input)
	} else {
		return 0, errors.New(fmt.Sprintf("Error: Hold on cowboy. No such part (%d) of this day task", part))
	}
}
