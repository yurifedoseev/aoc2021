package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Point struct {
	I int
	J int
}

func run() error {
	lines, err := parseLines("day11/input.txt")
	if err != nil {
		return err
	}

	err = gold(lines)
	if err != nil {
		return err
	}

	return nil
}

type Travers struct {
	CurrentLevel int
	Highlights int
	LevelHigh map[Point]bool
	LevelAdjusted  map[Point]bool
}


func silver(days int, lines [][]int) error {
	t := &Travers{
		CurrentLevel: 0,
		Highlights: 0,
	}

	simulateDays := days
	for s := 0; s < simulateDays; s++ {
		t.CurrentLevel++
		t.LevelHigh = make(map[Point]bool)
		t.LevelAdjusted = make(map[Point]bool)
		hights := make([]Point, 0)
		// 1. update all points
		for i, line := range lines {
			for j := range line {
				p := Point{I: i, J: j}
				lines[p.I][p.J]++
				if lines[p.I][p.J] > 9 {
					lines[p.I][p.J] = 0
					t.LevelHigh[p] = true
					hights = append(hights, p)
				}
			}
		}

		for _, p := range hights {
			highlightNeighbors(t, p, lines)
		}

		fmt.Printf("\n\n\nafter day: %d", s+1)
		for _, line := range lines {
			fmt.Printf("\n")
			for _, val := range line {
				fmt.Printf("%d", val)
			}
		}
		fmt.Printf("\n\n")
	}

	fmt.Printf("\ntotal highlights: %d", t.Highlights)
	return nil
}

func gold(lines [][]int) error {
	t := &Travers{
		CurrentLevel: 0,
		Highlights: 0,
	}

	simulateDays := 1000
	for s := 0; s < simulateDays; s++ {
		t.CurrentLevel++
		t.LevelHigh = make(map[Point]bool)
		t.LevelAdjusted = make(map[Point]bool)
		hights := make([]Point, 0)
		// 1. update all points
		for i, line := range lines {
			for j := range line {
				p := Point{I: i, J: j}
				lines[p.I][p.J]++
				if lines[p.I][p.J] > 9 {
					lines[p.I][p.J] = 0
					t.LevelHigh[p] = true
					hights = append(hights, p)
				}
			}
		}

		for _, p := range hights {
			highlightNeighbors(t, p, lines)
		}

		fmt.Printf("\n\n\nafter day: %d", s+1)
		for _, line := range lines {
			fmt.Printf("\n")
			for _, val := range line {
				fmt.Printf("%d", val)
			}
		}
		fmt.Printf("\n\n")

		if len(t.LevelHigh) == len(lines) * len(lines[0]) {
			fmt.Printf("\nflash day is %d!!!", s+1)
			break
		}
	}

	return nil
}

func highlightNeighbors(t *Travers, lightPoint Point, lines [][]int){
	t.LevelHigh[lightPoint] = true
	t.Highlights++

	points := make([]Point, 0)
	i,j := lightPoint.I, lightPoint.J
	jLen := len(lines[i])
	// i, j+1
	if j < jLen - 1 {
		points = append(points, Point{I:i, J: j+1})
	}
	// i, j-1
	if j > 0 {
		points = append(points, Point{I:i, J: j-1})
	}
	// i-1, j
	if i > 0 {
		points = append(points, Point{I:i-1, J: j})
	}
	// i+1, j
	if i < len(lines) - 1 {
		points = append(points, Point{I:i+1, J: j})
	}
	// i-1, j-1
	if i > 0 && j > 0 {
		points = append(points, Point{I:i-1, J: j-1})
	}
	// i-1, j+1
	if i > 0 && j < jLen - 1 {
		points = append(points, Point{I:i-1, J: j+1})
	}
	// i+1, j-1
	if i < len(lines) - 1 && j > 0 {
		points = append(points, Point{I:i+1, J: j-1})
	}
	// i+1, j+1
	if i < len(lines) - 1 && j < jLen - 1 {
		points = append(points, Point{I:i+1, J: j+1})
	}

	for _, p := range points {
		//if t.LevelAdjusted[p]
		if t.LevelHigh[p] {
			continue
		}

		t.LevelAdjusted[p] = true
		lines[p.I][p.J]++
		if lines[p.I][p.J] > 9 {
			lines[p.I][p.J] = 0
			highlightNeighbors(t, p, lines)
		}
	}
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
