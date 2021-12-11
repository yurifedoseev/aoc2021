package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func run() error {
	lines, err := parseLines("day9/input.txt")
	if err != nil {
		return err
	}

	fmt.Printf("\nparse %d lines", len(lines))
	lowestPoints := countLowestPoints(lines)
	basins := make([]*Basin, 0, len(lowestPoints))
	for _, start := range lowestPoints {
		basin := &Basin{
			Lowest:  start,
			Visited: make(map[Point]bool),
		}
		basins = append(basins, basin)
		Visit(basin, lines, start)
	}

	sort.Slice(basins, func(i, j int) bool {
		return len(basins[i].Visited) > len(basins[j].Visited)
	})

	sum := basins[0].Sum() * basins[1].Sum() * basins[2].Sum()
	fmt.Printf("\n: sum is %d, 0: %d, 1: %d, 2: %d", sum, basins[0].Sum(), basins[1].Sum(), basins[2].Sum())
	return nil
}

func Visit(b *Basin, lines [][]int, currP Point) {
	if _, ok := b.Visited[currP]; ok {
		return
	}
	b.Visited[currP] = true
	nextPoints := make([]Point, 0)

	i,j := currP.I, currP.J
	// val := lines[i][j]

	// up
	if i > 0 {
		// abs(lines[i-1][j], val) == 1
		if lines[i-1][j] != 9 {
			nextPoints = append(nextPoints, Point{I: i-1, J: j})
		}
	}
	// down
	if i < len(lines) - 1 {
		// abs(lines[i+1][j], val) == 1
		if lines[i+1][j] != 9 {
			nextPoints = append(nextPoints, Point{I: i + 1, J: j})
		}
	}
	// left
	if j > 0 {
		//abs(lines[i][j-1], val) == 1 &&
		if lines[i][j-1] != 9 {
			nextPoints = append(nextPoints, Point{I: i, J: j-1})
		}
	}
	// right
	if j < len(lines[i]) - 1 {
		// if abs(lines[i][j+1], val) == 1 && lines[i][j+1] != 9 {
		if lines[i][j+1] != 9 {
			nextPoints = append(nextPoints, Point{I: i, J: j+1})
		}
	}
	for _, p := range nextPoints {
		Visit(b, lines, p)
	}
}

type Basin struct {
	Lowest Point
	Visited map[Point]bool
}

func (b *Basin) Sum() int {
	return len(b.Visited)
}

func abs(a,b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

type Point struct {
	I int
	J int
}

func countLowestPoints(lines [][]int) []Point {
	points := make([]Point, 0)

	for i, line := range lines {
		for j, el := range line {
			// сбоку
			isLowest := true
			// up
			if i > 0 {
				isLowest = isLowest && lines[i-1][j] > el
			}

			// down
			if i < len(lines) - 1 {
				isLowest = isLowest && lines[i+1][j] > el
			}

			// left
			if j > 0 {
				isLowest = isLowest && lines[i][j-1] > el
			}

			// right
			if j < len(line) - 1 {
				isLowest = isLowest && lines[i][j+1] > el
			}

			if isLowest {
				fmt.Printf("\n element %d: %d is low: %d", i, j, el)
				points = append(points, Point{I:i, J:j})
			}

		}
	}
	return points
}

func parseLines(filename string) ([][]int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	lines := make([][]int, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		raw := scanner.Text()
		line := make([]int, 0, len(raw))
		for _, el := range raw {
			val, err := strconv.Atoi(string(el))
			if err != nil {
				return nil, fmt.Errorf("failed convert %s to int: %w", string(el), err)
			}
			line = append(line, val)
		}
		lines = append(lines, line)
	}

	return lines, nil
}

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}
