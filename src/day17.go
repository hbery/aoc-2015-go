package hbery_aoc2015

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Container struct {
	capacity int64
}

// container lists can be different size, so need a wrapper for incositent matrix size
type ContainerSet struct {
	containers *[]*Container
}

func test_container_set(total_capacity int64, test_set []*Container, combinations *[]*ContainerSet) {
	var sum int64 = 0
	for _, con := range test_set {
		sum += con.capacity
	}
	
	if sum == total_capacity {
		*combinations = append(*combinations, &ContainerSet{containers: &test_set})
	}
}

func get_next_container_set(pos, from int, total_capacity int64, containers *[]*Container, working_container_set []*Container, combinations *[]*ContainerSet) {
	for i := from; int(i) < len(*containers); i++ {
		working_container_set[pos] = (*containers)[i]
		
		if pos == (cap(working_container_set) - 1) {
			test_container_set(total_capacity, working_container_set, combinations)
		} else {
			get_next_container_set(pos + 1, i + 1, total_capacity, containers, working_container_set, combinations)
		}
	}
}

func choose_containers_sets(total_capacity int64, containers *[]*Container, combinations *[]*ContainerSet) {
	for arrlen := 1; arrlen <= len(*containers); arrlen++ {
		con_set := make([]*Container, arrlen)
		get_next_container_set(0, 0, total_capacity, containers, con_set, combinations)
	}
}

func count_smallest_combinations_of_containers(combinations *[]*ContainerSet) int64 {
	var smallest_set int = 999

	for _, cs := range *combinations {
		csn := len(*cs.containers)
		if csn < smallest_set {
			smallest_set = csn
		}
	}

	var smallest_sets_combinations_number int64 = 0
	for _, cs := range *combinations {
		csn := len(*cs.containers)
		if csn == smallest_set {
			smallest_sets_combinations_number++
		}
	}

	return smallest_sets_combinations_number
}

func day17_p1(input string) (int64, error) {
	var result int64 = 0
	var total_capacity int64 = 150
	var containters []*Container
	var val int64
	var err error

	for _, c := range strings.Split(strings.TrimSuffix(input, "\n"), "\n") {
		if val, err = strconv.ParseInt(c, 10, 64); err != nil {
			return -1, errors.New(fmt.Sprintf("Error: Cannot parse number %s.", c))
		}

		containters = append(containters, &Container{capacity: val})
	}

	var possible_container_sets []*ContainerSet
	choose_containers_sets(total_capacity, &containters, &possible_container_sets)

	result = int64(len(possible_container_sets))

	return result, nil
}

func day17_p2(input string) (int64, error) {
	var result int64 = 0
	var total_capacity int64 = 150
	var containters []*Container
	var val int64
	var err error

	for _, c := range strings.Split(strings.TrimSuffix(input, "\n"), "\n") {
		if val, err = strconv.ParseInt(c, 10, 64); err != nil {
			return -1, errors.New(fmt.Sprintf("Error: Cannot parse number %s.", c))
		}

		containters = append(containters, &Container{capacity: val})
	}

	var possible_container_sets []*ContainerSet
	choose_containers_sets(total_capacity, &containters, &possible_container_sets)

	result = count_smallest_combinations_of_containers(&possible_container_sets)

	return result, nil
}

func Solution_Day17(part int, input string) (int64, error) {
	if part == 1 {
		return day17_p1(input)
	} else if part == 2 {
		return day17_p2(input)
	} else {
		return 0, errors.New(fmt.Sprintf("Error: Hold on cowboy. No such part (%d) of this day task", part))
	}
}
