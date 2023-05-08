package hbery_aoc2015

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Ingredient struct {
	name		string

	capacity	int64
	durability	int64
	flavor		int64
	texture		int64
	calories	int64
}

type IngredientsUsed struct {
	ingredients	[]*Ingredient
	used		[]bool
}

func make_ingredient(name string, capacity, durability, flavor, texture, calories int64) *Ingredient {
	return &Ingredient{
		name: name,

		capacity: capacity,
		durability: durability,
		flavor: flavor,
		texture: texture,
		calories: calories,
	}
}

func parse_ingredients(input string) ([]*Ingredient, error) {
	var ingredients []*Ingredient
	var cp, dr, fl, tx, cal int64
	var err error

	for _, line := range strings.Split(input, "\n") {
		ld := strings.Split(line, ": ")
		props := strings.Split(ld[1], ", ")

		// capacity
		if cp, err = strconv.ParseInt(strings.Split(props[0], " ")[1], 10, 64); err != nil {
			return nil, errors.New(fmt.Sprintf("Error: Failed parsing property (%s) for ingredient %s.", props[0], ld[0]))
		}
		
		// durability
		if dr, err = strconv.ParseInt(strings.Split(props[1], " ")[1], 10, 64); err != nil {
			return nil, errors.New(fmt.Sprintf("Error: Failed parsing property (%s) for ingredient %s.", props[1], ld[0]))
		}

		// flavor
		if fl, err = strconv.ParseInt(strings.Split(props[2], " ")[1], 10, 64); err != nil {
			return nil, errors.New(fmt.Sprintf("Error: Failed parsing property (%s) for ingredient %s.", props[2], ld[0]))
		}

		// texture
		if tx, err = strconv.ParseInt(strings.Split(props[3], " ")[1], 10, 64); err != nil {
			return nil, errors.New(fmt.Sprintf("Error: Failed parsing property (%s) for ingredient %s.", props[3], ld[0]))
		}

		// calories
		if cal, err = strconv.ParseInt(strings.Split(props[4], " ")[1], 10, 64); err != nil {
			return nil, errors.New(fmt.Sprintf("Error: Failed parsing property (%s) for ingredient %s.", props[4], ld[0]))
		}

		ingredients = append(ingredients, make_ingredient(ld[0], cp, dr, fl, tx, cal))
	}
	
	return ingredients, nil
}

func calulate_score(i_a map[*Ingredient]int64) (int64, int64) {
	var score int64 = 1
	var props []int64 = []int64{0, 0, 0, 0}
	var calories int64 = 0

	for i, a := range i_a {
		props[0] += i.capacity * a
		props[1] += i.durability * a
		props[2] += i.flavor * a
		props[3] += i.texture * a
		calories += i.calories * a
	}

	for _, p := range props {
		if p < 0 {
			score = 0
			break
		}

		score *= p
	}

	return score, calories
}

func ingredients_left_to_supply(ingredients_prepared *IngredientsUsed) int {
	var left int = 0

	for _, l := range (*ingredients_prepared).used {
		if !l {
			left += 1
		}
	}

	return left
}

func get_not_supplied_ingredient(ingredients_prepared *IngredientsUsed) int {
	for i, l := range (*ingredients_prepared).used {
		if !l {
			return i
		}
	}

	return -1
}

func permutate_ingredients(teaspoons int, ingredients_used *IngredientsUsed, ingredients_amount map[*Ingredient]int64, perms *[]*map[*Ingredient]int64) {
	if ingredients_left_to_supply(ingredients_used) == 1 {
		idx := get_not_supplied_ingredient(ingredients_used)
		in := ingredients_used.ingredients[idx]
		ingredients_amount[in] = int64(teaspoons)

		// stuuuupid map copy
		var new_ing_amnt map[*Ingredient]int64 = make(map[*Ingredient]int64, len(ingredients_amount))
		for k, v := range ingredients_amount {
			new_ing_amnt[k] = v
		}

		*perms = append(*perms, &new_ing_amnt)
	} else {
		for i := 0; i < teaspoons; i++ {
			idx := get_not_supplied_ingredient(ingredients_used)
			in := ingredients_used.ingredients[idx]
			ingredients_amount[in] = int64(i+1)

			ingredients_used.used[idx] = true
			permutate_ingredients(teaspoons - (i+1), ingredients_used, ingredients_amount, perms)
			ingredients_used.used[idx] = false
		}
	}
}

func choose_highest_score_recipe(ingredients_amounts_permutations *[]*map[*Ingredient]int64) int64 {
	var highest_score int64 = 0

	for _, p := range *ingredients_amounts_permutations {
		current_score, _ := calulate_score(*p)

		if current_score > highest_score {
			highest_score = current_score
		}
	}

	return highest_score
}

func choose_highest_score_recipe2(ingredients_amounts_permutations *[]*map[*Ingredient]int64) int64 {
	var highest_score int64 = 0

	for _, p := range *ingredients_amounts_permutations {
		current_score, calories := calulate_score(*p)

		if current_score > highest_score && calories == 500 {
			highest_score = current_score
		}
	}

	return highest_score
}

func day15_p1(input string) (int64, error) {
	var result int64 = 0
	var ingredients []*Ingredient
	var teaspoons int = 100
	var err error

	input = strings.TrimSuffix(input, "\n")
	if ingredients, err = parse_ingredients(input); err != nil {
		return -1, err
	}

	var ingredient_amounts map[*Ingredient]int64 = make(map[*Ingredient]int64)
	var ingredient_supplied IngredientsUsed = IngredientsUsed{
		ingredients: ingredients,
		used: make([]bool, len(ingredients)),
	}
	for i, in := range ingredients {
		ingredient_amounts[in] = 0
		ingredient_supplied.used[i] = false
	}
	
	var ing_amnt_scores []*map[*Ingredient]int64
	permutate_ingredients(teaspoons, &ingredient_supplied, ingredient_amounts, &ing_amnt_scores)

	result = choose_highest_score_recipe(&ing_amnt_scores)

	return result, nil
}

func day15_p2(input string) (int64, error) {
	var result int64 = 0
	var ingredients []*Ingredient
	var teaspoons int = 100
	var err error

	input = strings.TrimSuffix(input, "\n")
	if ingredients, err = parse_ingredients(input); err != nil {
		return -1, err
	}

	var ingredient_amounts map[*Ingredient]int64 = make(map[*Ingredient]int64)
	var ingredient_supplied IngredientsUsed = IngredientsUsed{
		ingredients: ingredients,
		used: make([]bool, len(ingredients)),
	}
	for i, in := range ingredients {
		ingredient_amounts[in] = 0
		ingredient_supplied.used[i] = false
	}
	
	var ing_amnt_scores []*map[*Ingredient]int64
	permutate_ingredients(teaspoons, &ingredient_supplied, ingredient_amounts, &ing_amnt_scores)

	result = choose_highest_score_recipe2(&ing_amnt_scores)

	return result, nil
}

func Solution_Day15(part int, input string) (int64, error) {
	if part == 1 {
		return day15_p1(input)
	} else if part == 2 {
		return day15_p2(input)
	} else {
		return 0, errors.New(fmt.Sprintf("Error: Hold on cowboy. No such part (%d) of this day task", part))
	}
}
