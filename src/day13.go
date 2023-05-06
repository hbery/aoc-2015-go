package hbery_aoc2015

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Person struct {
	name		string
	guests 		map[string]int64
}

type Table struct {
	happiness	int64
	seats		[]*Person
}

func make_person(name string) *Person {
	return &Person{
		name: name,
		guests: make(map[string]int64),
	}
}

func make_table() *Table {
	return &Table{
		happiness: 0,
		seats: nil,
	}
}

func (t *Table) addSeat(p *Person) {
	t.seats = append(t.seats, p)
}

func (t *Table) print() {
	var repr string
	for _, p := range t.seats {
		repr = fmt.Sprintf("%s %s,", repr, p.name)
	}
	repr = fmt.Sprintf("%s, H:%d", repr, t.happiness)
	fmt.Println(repr)
}

func (t *Table) calculateHapiness() {
	for i := 0; i < len(t.seats); i++ {
		if i != len(t.seats) - 1 {
			t.happiness += t.seats[i].guests[t.seats[i+1].name] + t.seats[i+1].guests[t.seats[i].name]
		} else {
			t.happiness += t.seats[i].guests[t.seats[0].name] + t.seats[0].guests[t.seats[i].name]
		}
	}
}

func get_highest_happiness_table_idx(tables *[]*Table) *Table {
	var best_table *Table
	
	for _, t := range *tables {
		if best_table == nil {
			best_table = t
			continue
		}

		if t.happiness > best_table.happiness {
			best_table = t
		}
	}

	return best_table
}

func get_or_create_person(people *[]*Person, name string) *Person {
	for _, p := range *people {
		if name == p.name {
			return p
		}
	}
	n_p := make_person(name)
	*people = append(*people, n_p)
	return n_p
}

func permutate_tables(num_people int, people []*Person, tables *[]*Table) {
	if num_people == 1 {
		var table *Table = make_table()

		for _, p := range people {
			table.addSeat(p)
		}

		*tables = append(*tables, table)
	} else {
		for i := 0; i < num_people - 1; i++ {
			permutate_tables(num_people - 1, people, tables)

			if num_people % 2 == 0 {
				people[i], people[num_people - 1] = people[num_people - 1], people[i]
			} else {
				people[0], people[num_people - 1] = people[num_people - 1], people[0]
			}
		}

		permutate_tables(num_people - 1, people, tables)
	}
}

func day13_p1(input string) (int64, error) {
	var people []*Person
	var tables []*Table
	var err error

	input = strings.TrimSuffix(input, "\n")

	for i, line := range strings.Split(input, "\n") {
		line_divided := strings.Split(strings.TrimSuffix(line, "."), " ")
		var happiness_differentiator int64

		if line_divided[2] == "gain" {
			if happiness_differentiator, err = strconv.ParseInt(fmt.Sprintf("%s", line_divided[3]), 10, 64); err != nil {
				return -1, errors.New(fmt.Sprintf("Error: Error during parsing number for line %d.", i+1))
			}
		} else if line_divided[2] == "lose" {
			if happiness_differentiator, err = strconv.ParseInt(fmt.Sprintf("-%s", line_divided[3]), 10, 64); err != nil {
				return -1, errors.New(fmt.Sprintf("Error: Error during parsing number for line %d.", i+1))
			}
		}

		get_or_create_person(&people, line_divided[0]).guests[line_divided[10]] = happiness_differentiator
	}

	// fmt.Printf("%T{\n", people)
	// for _, p := range people {
	// 	fmt.Printf("%+#v,\n", *p)
	// }
	// fmt.Printf("}\n")

	permutate_tables(len(people), people, &tables)
	
	for _, t := range tables {
		t.calculateHapiness()
	}

	return get_highest_happiness_table_idx(&tables).happiness, nil
}

func day13_p2(input string) (int64, error) {
	var people []*Person
	var tables []*Table
	var err error

	input = strings.TrimSuffix(input, "\n")

	for i, line := range strings.Split(input, "\n") {
		line_divided := strings.Split(strings.TrimSuffix(line, "."), " ")
		var happiness_differentiator int64

		if line_divided[2] == "gain" {
			if happiness_differentiator, err = strconv.ParseInt(fmt.Sprintf("%s", line_divided[3]), 10, 64); err != nil {
				return -1, errors.New(fmt.Sprintf("Error: Error during parsing number for line %d.", i+1))
			}
		} else if line_divided[2] == "lose" {
			if happiness_differentiator, err = strconv.ParseInt(fmt.Sprintf("-%s", line_divided[3]), 10, 64); err != nil {
				return -1, errors.New(fmt.Sprintf("Error: Error during parsing number for line %d.", i+1))
			}
		}

		get_or_create_person(&people, line_divided[0]).guests[line_divided[10]] = happiness_differentiator
	}

	// Adding 'Me' to the equation
	me := make_person("Me")
	for _, p := range people {
		p.guests[me.name] = 0
		me.guests[p.name] = 0
	}
	people = append(people, me)

	permutate_tables(len(people), people, &tables)
	
	for _, t := range tables {
		t.calculateHapiness()
	}

	return get_highest_happiness_table_idx(&tables).happiness, nil
}

func Solution_Day13(part int, input string) (int64, error) {
	if part == 1 {
		return day13_p1(input)
	} else if part == 2 {
		return day13_p2(input)
	} else {
		return 0, errors.New(fmt.Sprintf("Error: Hold on cowboy. No such part (%d) of this day task", part))
	}
}
