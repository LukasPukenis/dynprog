/*
	Dynamic programming experiment using hopper game which allows picking up
	speed on a 2d grid
*/

package main

import (
	"fmt"
	"math"
	"strconv"
)

const (
	// because we can visit the same quares, we need a fuse value to terminate the checking
	// as we can't tell that we're jumping in loops
	max_score = 1000

	// give ourselves a velocity limit
	velocity_limit = 4
)

// helper to reduce verbose code - panics in case of failure to parse
func str2int(i string) int {
	a, err := strconv.Atoi(i)
	if err != nil {
		panic(fmt.Sprintf("cant parse: %s -> %v", i, err))
	}

	return a
}

type state struct {
	g        *game
	velocity vec2
	pos      vec2
	score    int
	path     []vec2
}

func (s state) completed() bool {
	return s.pos.equals(s.g.finish)
}

func (s *state) collides(v vec2) bool {
	return s.g.collides(v)
}

func solve(g *game) int {
	var states []state
	winning_score := math.MaxInt64

	// push initial state. We emulate recursion via a slice
	// I find it easier to debug
	states = append(states, state{
		g:   g,
		pos: g.start,
	})

	for len(states) > 0 {
		s := states[0]
		states = states[1:]

		// there's already a better result, skip
		if s.score >= winning_score {
			continue
		}

		// too much of going back and forth, skip
		if s.score >= max_score {
			continue
		}

		if s.completed() {
			if s.score < winning_score {
				winning_score = s.score
			}

			continue
		}

		veloffsets := []vec2{
			{0, 0},
			{0, -1},
			{0, 1},
			{1, -1},
			{1, 0},
			{1, 1},
			{-1, -1},
			{-1, 0},
			{-1, 1},
		}

		for _, off := range veloffsets {
			newvel := s.velocity.add(off)
			newvel = newvel.clamp(
				vec2{
					x: -velocity_limit,
					y: -velocity_limit,
				},
				vec2{
					x: velocity_limit,
					y: velocity_limit,
				})

			newpos := s.pos.add(newvel)

			if !s.collides(newpos) {
				states = append(states, state{
					g:        g,
					pos:      newpos,
					velocity: newvel,
					score:    s.score + 1,
					path:     append(s.path, newpos),
				})
			}
		}
	}

	// don't leak implementation details, return -1 when not found
	if winning_score == math.MaxInt64 {
		return -1
	}

	return winning_score
}
