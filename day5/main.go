package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

type Line struct {
	A Point
	B Point
}

func (l *Line) GetPoints() []Point {
	points := make([]Point, 0)
	if l.A.X == l.B.X {
		min, max := l.A, l.B
		if max.Y < min.Y {
			min, max = max, min
		}
		for i := min.Y; i <= max.Y; i++ {
			points = append(points, Point {
				X: l.A.X,
				Y: i,
			})
		}
	} else if l.A.Y == l.B.Y {
		min, max := l.A, l.B
		if max.X < min.X {
			min, max = max, min
		}
		for i := min.X; i <= max.X; i++ {
			points = append(points, Point {
				X: i,
				Y: l.A.Y,
			})
		}

	} else {
		xStep := 1
		yStep := 1
		if l.A.X > l.B.X {
			xStep = -1
		}
		if l.A.Y > l.B.Y {
			yStep = -1
		}
		p := l.A
		for p != l.B {
			points = append(points, p)
			//newX := l.B.X
			//if p.X != newX {
			//	newX = p.X + xStep
			//}
			//newY := l.B.Y
			//if p.Y != newY {
			//	newY = p.Y + yStep
			//}
			p = Point {
				X: p.X + xStep,
				Y: p.Y + yStep,
			}
		}
		points = append(points, l.B)
	}

	return points
}

func run() error {
	lines, err := parseLines("day5/input.txt")
	if err != nil {
		return err
	}

	fmt.Printf("\nparsed %d lines", len(lines))
	sect := make(map[Point]int)
	for _, l := range lines {
		for _, p := range l.GetPoints() {
			sect[p]++
		}
	}

	count := 0
	for _, val := range sect {
		if val >= 2 {
			count++
		}
	}
	fmt.Printf("\n%d intersection with 2 and higher", count)
	return nil
}

func parseLines(filename string) ([]*Line, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]*Line, 0)
	for scanner.Scan() {
		raw := scanner.Text()
		parts := strings.Split(raw," -> ")
		pointA, err := parsePoint(parts[0])
		if err != nil {
			return nil, fmt.Errorf("failed to parse A: %w", err)
		}
		pointB, err := parsePoint(parts[1])
		if err != nil {
			return nil, fmt.Errorf("failed to parse B: %w", err)
		}

		lines = append(lines, &Line{A:pointA, B: pointB})
	}

	return lines, nil
}

func parsePoint(raw string) (Point, error) {
	rawSplit := strings.Split(raw, ",")
	x, err := strconv.Atoi(rawSplit[0])
	if err != nil {
		return Point{}, err
	}
	y, err := strconv.Atoi(rawSplit[1])
	if err != nil {
		return Point{}, err
	}
	return Point{X:x, Y:y}, nil
}

func main() {
	err := run()
	if err != nil {
		panic(err)
	}
}
