package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

// x1 < x2 && y1 < y2
type obstacleGrid struct {
	x1, y1 int
	x2, y2 int
}

func (o obstacleGrid) collides(v vec2) bool {
	return v.x >= o.x1 && v.x <= o.x2 && v.y >= o.y1 && v.y <= o.y2
}

type game struct {
	dims          vec2
	obstacles     []obstacleGrid
	start, finish vec2
}

func (g *game) collides(v vec2) bool {
	if v.x < 0 || v.y < 0 || v.x >= g.dims.x || v.y >= g.dims.y {
		return true
	}

	for _, obs := range g.obstacles {
		if obs.collides(v) {
			return true
		}
	}

	return false
}

func parseGames(r io.Reader) ([]*game, error) {
	scanner := bufio.NewScanner(r)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	if len(lines) == 0 {
		log.Fatal("empty input")
	}

	gamescnt := str2int(lines[0])
	lines = lines[1:]

	var games []*game

	idx := 0
	for i := 0; i < gamescnt; i++ {
		dimsstr := strings.SplitN(lines[idx], " ", 2)
		dimw, dimh := str2int(dimsstr[0]), str2int(dimsstr[1])

		idx++
		startfinishstr := strings.SplitN(lines[idx], " ", 4)

		sx, sy, fx, fy := str2int(startfinishstr[0]), str2int(startfinishstr[1]), str2int(startfinishstr[2]), str2int(startfinishstr[3])
		idx++
		obstaclescnt := str2int(lines[idx])

		var obs []obstacleGrid

		idx++
		for j := 0; j < obstaclescnt; j++ {
			obstacles := strings.SplitN(lines[idx], " ", 4)

			x1, x2, y1, y2 := str2int(obstacles[0]), str2int(obstacles[1]), str2int(obstacles[2]), str2int(obstacles[3])
			obs = append(obs, obstacleGrid{
				x1: x1,
				y1: y1,
				x2: x2,
				y2: y2,
			})

			idx++
		}

		g := &game{
			dims:      vec2{x: dimw, y: dimh},
			start:     vec2{x: sx, y: sy},
			finish:    vec2{x: fx, y: fy},
			obstacles: obs,
		}

		games = append(games, g)
	}

	return games, nil
}

func main() {
	data, err := ioutil.ReadFile("games.txt")
	if err != nil {
		log.Fatal(err)
	}

	games, err := parseGames(bytes.NewReader(data))

	if err != nil {
		log.Fatalf("parsing games: %v", err)
	}

	for _, game := range games {
		hops := solve(game)
		switch hops {
		case -1:
			fmt.Println("No solution.")
		default:
			fmt.Printf("Optimal solution takes %d hops.\n", hops)
		}
	}
}
