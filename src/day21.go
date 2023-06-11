package hbery_aoc2015

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var BASE_PLAYER_HP int = 100

type ShopItem struct {
	item_type	string
	name		string
	damage		int
	armor		int
	cost		int64
}

type RPGCharacter struct {
	max_hit_points	int
	hit_points	int
	base_damage	int
	base_armor	int
	equipment	[]*ShopItem
}

func make_rpgcharacter(hp, b_damage, b_armor int, items []*ShopItem) (c *RPGCharacter) {
	c = &RPGCharacter{}
	c.max_hit_points = hp
	c.base_damage = b_damage
	c.base_armor = b_armor
	c.equipment = items
	c.reset_hp()
	return
}

func (c *RPGCharacter) repr() (repr string) {
	var items string = "[ "
	for _, it := range c.equipment {
		items = fmt.Sprintf("%s%s:{%s -> d:%d a:%d c:%d} ", items, it.item_type, it.name, it.damage, it.armor, it.cost)
	}
	items = fmt.Sprintf("%s%s", items, "]")

	repr = fmt.Sprintf("RPGCharacter{hp:%d bd:%d ba:%d eq:%s}", c.max_hit_points, c.base_damage, c.base_armor, items)
	return
}

func (c *RPGCharacter) get_hit(by *RPGCharacter) {
	var character_armor int = c.base_armor
	var enemy_damage int = by.base_damage
	var actual_damage int = 0

	if c.equipment != nil {
		for _, i := range c.equipment {
			character_armor += i.armor
		}
	}

	if by.equipment != nil {
		for  _, i := range by.equipment {
			enemy_damage += i.damage
		}
	}

	if actual_damage = (enemy_damage - character_armor); actual_damage < 1 {
		actual_damage = 1
	}

	c.hit_points -= actual_damage
}

func (c *RPGCharacter) reset_hp() {
	c.hit_points = c.max_hit_points
}

func (c *RPGCharacter) sum_damage() (real_dmg int) {
	real_dmg += c.base_damage
	if c.equipment != nil {
		for _, i := range c.equipment {
			real_dmg += i.damage
		}
	}
	return
}

func (c *RPGCharacter) sum_armor() (real_armor int) {
	real_armor += c.base_armor
	if c.equipment != nil {
		for _, i := range c.equipment {
			real_armor += i.armor
		}
	}
	return
}

func (c *RPGCharacter) equipment_cost() (eq_cost_sum int64) {
	for _, item := range c.equipment {
		eq_cost_sum += item.cost
	}

	return eq_cost_sum
}

func arr_copy[T any](arr *[]T) (arrcp []T) {
	arrcp = make([]T, len(*arr))
	copy(arrcp, *arr)
	return
}

func filter[T any](ss *[]T, test func(T) bool) (ret []T) {
	for _, el := range *ss { if test(el) { ret = append(ret, el) } }
	return
}

func parse_boss_stats(input string) (*RPGCharacter, error) {
	var boss_char *RPGCharacter = &RPGCharacter{}
	var err error

	for _, line := range strings.Split(input, "\n") {
		var parsed_num int64
		ld := strings.Split(line, ":")

		if ld[0] == "Hit Points" {
			if parsed_num, err = strconv.ParseInt(strings.TrimSpace(ld[1]), 10, 64); err != nil {
				return nil, err
			}
			boss_char.max_hit_points = int(parsed_num)
		}

		if ld[0] == "Damage" {
			if parsed_num, err = strconv.ParseInt(strings.TrimSpace(ld[1]), 10, 64); err != nil {
				return nil, err
			}
			boss_char.base_damage = int(parsed_num)
		}

		if ld[0] == "Armor" {
			if parsed_num, err = strconv.ParseInt(strings.TrimSpace(ld[1]), 10, 64); err != nil {
				return nil, err
			}
			boss_char.base_armor = int(parsed_num)
		}
	}

	boss_char.reset_hp()
	return boss_char, nil
}

func get_static_shop_items() []*ShopItem {
/*
	Weapons:   Cost  Damage  Armor
	Dagger        8       4      0
	Shortsword   10       5      0
	Warhammer    25       6      0
	Longsword    40       7      0
	Greataxe     74       8      0
	
	Armor:      Cost  Damage  Armor
	Leather       13       0      1
	Chainmail     31       0      2
	Splintmail    53       0      3
	Bandedmail    75       0      4
	Platemail    102       0      5
	
	Rings:      Cost  Damage  Armor
	Damage +1     25       1      0
	Damage +2     50       2      0
	Damage +3    100       3      0
	Defense +1    20       0      1
	Defense +2    40       0      2
	Defense +3    80       0      3
*/
	return []*ShopItem{
		{item_type: "weapon", name: "Dagger", cost: 8, damage: 4, armor: 0},
		{item_type: "weapon", name: "Shortsword", cost: 10, damage: 5, armor: 0},
		{item_type: "weapon", name: "Warhammer", cost: 25, damage: 6, armor: 0},
		{item_type: "weapon", name: "Longsword", cost: 40, damage: 7, armor: 0},
		{item_type: "weapon", name: "Greataxe", cost: 74, damage: 8, armor: 0},

		// armor empty slot placeholder
		{item_type: "armor", name: "No armor", cost: 0, damage: 0, armor: 0},

		{item_type: "armor", name: "Leather", cost: 13, damage: 0, armor: 1},
		{item_type: "armor", name: "Chainmail", cost: 31, damage: 0, armor: 2},
		{item_type: "armor", name: "Splintmail", cost: 53, damage: 0, armor: 3},
		{item_type: "armor", name: "Bandedmail", cost: 75, damage: 0, armor: 4},
		{item_type: "armor", name: "Platemail", cost: 102, damage: 0, armor: 5},

		// jewelry empty slots placeholders
		{item_type: "jewelry", name: "Empty 1", cost: 0, damage: 0, armor: 0},
		{item_type: "jewelry", name: "Empty 2", cost: 0, damage: 0, armor: 0},

		{item_type: "jewelry", name: "Damage +1", cost: 25, damage: 1, armor: 0},
		{item_type: "jewelry", name: "Damage +2", cost: 50, damage: 2, armor: 0},
		{item_type: "jewelry", name: "Damage +3", cost: 100, damage: 3, armor: 0},
		{item_type: "jewelry", name: "Defense +1", cost: 20, damage: 0, armor: 1},
		{item_type: "jewelry", name: "Defense +2", cost: 40, damage: 0, armor: 2},
		{item_type: "jewelry", name: "Defense +3", cost: 80, damage: 0, armor: 3},
	}
}

func permutate_equipment(available_items *[]*ShopItem) (perms []*RPGCharacter) {
	weapon_type := func(item *ShopItem) bool { return item.item_type == "weapon" }
	armor_type := func(item *ShopItem) bool { return item.item_type == "armor" }
	jewelry_type := func(item *ShopItem) bool { return item.item_type == "jewelry" }

	for _, weapon := range filter(available_items, weapon_type) {
		var items []*ShopItem
		items = append(items, weapon)
		for _, armor := range filter(available_items, armor_type) {
			items = append(items, armor)
			jewels := filter(available_items, jewelry_type)
			for _, j_1 := range jewels {
				items = append(items, j_1)
				for _, j_2 := range jewels {
					if j_1 == j_2 { continue }
					items = append(items, j_2)

					perms = append(perms, make_rpgcharacter(BASE_PLAYER_HP, 0, 0, arr_copy(&items)))
					items = items[:len(items)-1]
				}
				items = items[:len(items)-1]
			}
			items = items[:len(items)-1]
		}
	}
	return
}

func has_player_won(boss, player *RPGCharacter) bool {
	var turn string = "player"

	for boss.hit_points > 0 && player.hit_points > 0 {
		if turn == "player" {
			boss.get_hit(player)
			turn = "boss"
		} else {
			player.get_hit(boss)
			turn = "player"
		}
	}

	if player.hit_points > 0 {
		return true
	}

	return false
}

func choose_cheapest_equipment(winning_characters []*RPGCharacter) (cheapest_set *RPGCharacter) {
	for _, character := range winning_characters {
		if cheapest_set == nil {
			cheapest_set = character
		}

		if character.equipment_cost() < cheapest_set.equipment_cost() {
			cheapest_set = character
		}
	}

	return
}

func choose_most_expensive_equipment(losing_characters []*RPGCharacter) (most_expensive_set *RPGCharacter) {
	for _, character := range losing_characters {
		if most_expensive_set == nil {
			most_expensive_set = character
		}

		if character.equipment_cost() > most_expensive_set.equipment_cost() {
			most_expensive_set = character
		}
	}

	return
}

func day21_p1(input string) (int64, error) {
	var result int64 = 0
	var boss_character *RPGCharacter
	var err error

	if boss_character, err = parse_boss_stats(strings.TrimSuffix(input, "\n")); err != nil {
		return -1, err
	}

	items := get_static_shop_items()

	possible_player_perms := permutate_equipment(&items)

	var winning_character_sets []*RPGCharacter
	for _, player := range possible_player_perms {
		if has_player_won(boss_character, player) {
			winning_character_sets = append(winning_character_sets, player)
		}
		boss_character.reset_hp()
	}

	cheapest_character_set := choose_cheapest_equipment(winning_character_sets)

	result = cheapest_character_set.equipment_cost()

	return result, nil
}

func day21_p2(input string) (int64, error) {
	var result int64 = 0
	var boss_character *RPGCharacter
	var err error

	if boss_character, err = parse_boss_stats(strings.TrimSuffix(input, "\n")); err != nil {
		return -1, err
	}

	items := get_static_shop_items()

	possible_player_perms := permutate_equipment(&items)

	var losing_character_sets []*RPGCharacter
	for _, player := range possible_player_perms {
		if ! has_player_won(boss_character, player) {
			losing_character_sets = append(losing_character_sets, player)
		}
		boss_character.reset_hp()
	}

	most_expensive_character_set := choose_most_expensive_equipment(losing_character_sets)

	result = most_expensive_character_set.equipment_cost()

	return result, nil
}

func Solution_Day21(part int, input string) (int64, error) {
	if part == 1 {
		return day21_p1(input)
	} else if part == 2 {
		return day21_p2(input)
	} else {
		return 0, errors.New(fmt.Sprintf("Error: Hold on cowboy. No such part (%d) of this day task", part))
	}
}
