package hbery_aoc2015

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type City struct {
	name            string
	neigh_distances map[string]int64
}

func newCity(name string) *City {
	c := new(City)
	c.name = name
	c.neigh_distances = make(map[string]int64)
	return c
}

func getCity(cities *[]*City, name string) *City {
	for _, c := range *cities {
		if c.name == name {
			return c
		}
	}

	new_c := newCity(name)
	*cities = append((*cities), new_c)
	return new_c
}

func permutateOverCities(num_cities int, cities []*City, perms *[][]string) {
	if num_cities == 1 {
		var city_names []string

		for _, c := range cities {
			city_names = append(city_names, c.name)
		}

		*perms = append(*perms, city_names)
	} else {
		for i := 0; i < num_cities - 1; i++ {
			permutateOverCities(num_cities - 1, cities, perms)

			if num_cities % 2 == 0 {
				cities[i], cities[num_cities - 1] = cities[num_cities - 1], cities[i]
			} else {
				cities[0], cities[num_cities - 1] = cities[num_cities - 1], cities[0]
			}
		}

		permutateOverCities(num_cities - 1, cities, perms)
	}
}

func calculatePath(cities []*City, path []string) int64 {
	var distance int64 = 0

	for i := 0; i < len(path) - 1; i++ {
		distance += getCity(&cities, path[i]).neigh_distances[path[i+1]]
	}

	return distance
}

func gluePathToString(path []string) string {
	return strings.Join(path, " -> ")
}

func assignDistancesToPaths(cities []*City, permutations [][]string) map[string]int64 {
	paths_with_distances := make(map[string]int64)
	
	for _, path := range permutations {
		paths_with_distances[gluePathToString(path)] = calculatePath(cities, path)
	}

	return paths_with_distances
}

func chooseShortestPath(path_to_distance map[string]int64) string {
	var shortest_distance int64 = int64(^uint64(0) >> 1)
	var chosen_path string

	for path, distance := range path_to_distance {
		if distance < shortest_distance {
			shortest_distance = distance
			chosen_path = path
		}
	}
	return chosen_path
}

func chooseLongestPath(path_to_distance map[string]int64) string {
	var longest_distance int64 = 0
	var chosen_path string

	for path, distance := range path_to_distance {
		if distance > longest_distance {
			longest_distance = distance
			chosen_path = path
		}
	}
	return chosen_path
}

func day09_p1(input string) (int64, error) {
	var cities []*City

	input = strings.TrimSuffix(input, "\n")

	for _, line := range strings.Split(input, "\n") {
		split_strings := strings.Split(line, " = ")
		towns := strings.Split(split_strings[0], " to ")

		if distance, err := strconv.ParseInt(split_strings[1], 10, 64); err != nil {
			return -1, errors.New(fmt.Sprintf("%s: Cannot parse distance string \"%s\" to int64", err, split_strings[1]))
		} else {
			getCity(&cities, towns[0]).neigh_distances[towns[1]] = distance
			getCity(&cities, towns[1]).neigh_distances[towns[0]] = distance
		}
	}

	var routes [][]string
	permutateOverCities(len(cities), cities, &routes)
	path_distances := assignDistancesToPaths(cities, routes)

	return path_distances[chooseShortestPath(path_distances)], nil
}

func day09_p2(input string) (int64, error) {
	var cities []*City

	input = strings.TrimSuffix(input, "\n")

	for _, line := range strings.Split(input, "\n") {
		split_strings := strings.Split(line, " = ")
		towns := strings.Split(split_strings[0], " to ")

		if distance, err := strconv.ParseInt(split_strings[1], 10, 64); err != nil {
			return -1, errors.New(fmt.Sprintf("%s: Cannot parse distance string \"%s\" to int64", err, split_strings[1]))
		} else {
			getCity(&cities, towns[0]).neigh_distances[towns[1]] = distance
			getCity(&cities, towns[1]).neigh_distances[towns[0]] = distance
		}
	}

	var routes [][]string
	permutateOverCities(len(cities), cities, &routes)
	path_distances := assignDistancesToPaths(cities, routes)

	return path_distances[chooseLongestPath(path_distances)], nil
}

func Solution_Day09(part int, input string) (int64, error) {
	if part == 1 {
		return day09_p1(input)
	} else if part == 2 {
		return day09_p2(input)
	} else {
		return 0, errors.New(fmt.Sprintf("Error: Hold on cowboy. No such part (%d) of this day task", part))
	}
}
