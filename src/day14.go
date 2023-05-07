package hbery_aoc2015

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Reindeer struct {
	name string

	top_speed int64

	fly_burst_time int64
	rest_burst_time int64

	current_distance int64
	fly_time_left int64
	rest_time_left int64

	points int64
}

func make_reindeer(name string, top_speed int64, fly_time int64, rest_time int64) *Reindeer {
	return &Reindeer{
		name: name,
		top_speed: top_speed,
		fly_burst_time: fly_time,
		rest_burst_time: rest_time,

		current_distance: 0,
		fly_time_left: fly_time,
		rest_time_left: rest_time,

		points: 0,
	}
}

func (r *Reindeer) proceedInRace() {
	if r.fly_time_left == 0 && r.rest_time_left == 0 {
		r.fly_time_left = r.fly_burst_time
		r.rest_time_left = r.rest_burst_time
	}

	if r.fly_time_left > 0 {
		r.current_distance += r.top_speed
		r.fly_time_left -= 1
		return
	}

	if r.rest_time_left > 0 {
		r.rest_time_left -= 1
		return
	}
}

func parse_reindeers(input string) ([]*Reindeer, error) {
	var reindeers []*Reindeer
	var ts, ft, rt int64
	var err error

	for _, line := range strings.Split(input, "\n") {
		ld := strings.Split(line, " ")

		// top speed
		if ts, err = strconv.ParseInt(ld[3], 10, 64); err != nil {
			return nil, errors.New(fmt.Sprintf("Error: Failed parsing top_speed (%s) for reindeer %s.", ld[3], ld[0]))
		}
		
		// fly time
		if ft, err = strconv.ParseInt(ld[6], 10, 64); err != nil {
			return nil, errors.New(fmt.Sprintf("Error: Failed parsing fly_time (%s) for reindeer %s.", ld[6], ld[0]))
		}

		// rest time
		if rt, err = strconv.ParseInt(ld[13], 10, 64); err != nil {
			return nil, errors.New(fmt.Sprintf("Error: Failed parsing rest_time (%s) for reindeer %s.", ld[13], ld[0]))
		}
		
		reindeers = append(reindeers, make_reindeer(ld[0], ts, ft, rt))
	}

	return reindeers, nil
}


func get_head_of_the_race(reindeers []*Reindeer) []*Reindeer {
	var head []*Reindeer

	var leader *Reindeer = nil
	for _, r := range reindeers {
		if leader == nil {
			leader = r
		}
		
		if r.current_distance > leader.current_distance {
			leader = r
		}
	}

	// double check for draw
	for _, r := range reindeers {
		if r.current_distance == leader.current_distance {
			head = append(head, r)
		}
	}

	return head
}

func race_reindeers(reindeers []*Reindeer, duration int) *Reindeer {
	for i := 0; i < duration + 1; i++ {
		for _, r := range reindeers {
			r.proceedInRace()
		}
	}

	var winner *Reindeer = nil
	for _, r := range reindeers {
		if winner == nil {
			winner = r
		}
		
		if r.current_distance > winner.current_distance {
			winner = r
		}
	}

	return winner
}

func race_reindeers_2(reindeers []*Reindeer, duration int) *Reindeer {
	for i := 0; i < duration + 1; i++ {
		for _, r := range reindeers {
			r.proceedInRace()
		}

		head_reindeers := get_head_of_the_race(reindeers)
		for _, r := range head_reindeers {
			r.points += 1
		}
	}

	var winner *Reindeer = nil
	for _, r := range reindeers {
		if winner == nil {
			winner = r
		}
		
		if r.points > winner.points {
			winner = r
		}
	}

	return winner
}

func day14_p1(input string) (int64, error) {
	var race_duration int = 2503
	var reindeers []*Reindeer
	var err error

	input = strings.TrimSuffix(input, "\n")

	if reindeers, err = parse_reindeers(input); err != nil {
		return -1, err
	}
	
	winner := race_reindeers(reindeers, race_duration)

	return winner.current_distance, nil
}

func day14_p2(input string) (int64, error) {
	var race_duration int = 2503
	var reindeers []*Reindeer
	var err error

	input = strings.TrimSuffix(input, "\n")

	if reindeers, err = parse_reindeers(input); err != nil {
		return -1, err
	}
	
	winner := race_reindeers_2(reindeers, race_duration)

	return winner.points, nil
}

func Solution_Day14(part int, input string) (int64, error) {
	if part == 1 {
		return day14_p1(input)
	} else if part == 2 {
		return day14_p2(input)
	} else {
		return 0, errors.New(fmt.Sprintf("Error: Hold on cowboy. No such part (%d) of this day task", part))
	}
}
