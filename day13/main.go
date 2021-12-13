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

type FoldCommand struct {
	Name string
	Value int
}

type Task struct {
	Points map[Point]bool
	Commands []FoldCommand
	MaxY int
	MaxX int
}

func run() error {
	task, err := parse("day13/input.txt")
	if err != nil {
		return err
	}

	fmt.Println("initial state")
	print(task)

	for _, cmd := range task.Commands {
		if cmd.Name == "y" {
			applyYCommand(task, cmd)
		} else if cmd.Name == "x" {
			applyXCommand(task, cmd)
		}
	}

	print(task)

	return nil
}

func applyYCommand(task *Task, cmd FoldCommand) {
	points := make([]Point, 0)
	for p := range task.Points {
		points = append(points, p)
	}

	for _, p := range points {
		if p.Y > cmd.Value {
			newY := p.Y- 2 * (p.Y - cmd.Value)
			task.Points[Point{X:p.X, Y:newY}] = true
			delete(task.Points, p)
		}
	}
	task.MaxY = cmd.Value - 1
}

func applyXCommand(task *Task, cmd FoldCommand) {
	points := make([]Point, 0)
	for p := range task.Points {
		points = append(points, p)
	}

	for _, p := range points {
		if p.X > cmd.Value {
			newX := p.X - 2 * (p.X - cmd.Value)
			task.Points[Point{X:newX, Y:p.Y}] = true
			delete(task.Points, p)
		}
	}
	task.MaxX = cmd.Value - 1
}

func print(task *Task) {
	fmt.Printf("\n\n")
	visible := 0
	for y := 0; y <= task.MaxY; y++ {
		fmt.Printf("\n")
		for x :=0; x <=task.MaxX; x++ {
			p := Point{
				X: x,
				Y: y,
			}
			if task.Points[p] {
				fmt.Printf("#")
				visible++
			} else {
				fmt.Printf(".")
			}
		}
	}

	fmt.Printf("\n visible %d dots", visible)
}

func parse(filename string) (*Task, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	task := &Task{
		Points:   make(map[Point]bool),
		Commands: make([]FoldCommand, 0),
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		raw := scanner.Text()
		raw = strings.TrimSpace(raw)
		if raw == "" {
			continue
		}

		if strings.Contains(raw, "fold") {
			splitted := strings.Split(raw, " ")
			rawVals:= strings.Split(splitted[len(splitted) - 1], "=")
			val, err := strconv.Atoi(rawVals[1])
			if err != nil {
				return nil, err
			}
			name := rawVals[0]
			task.Commands = append(task.Commands, FoldCommand{
				Name:  name,
				Value: val,
			})
		} else {
			splitted := strings.Split(raw, ",")
			x, err := strconv.Atoi(splitted[0])
			if err != nil {
				return nil, err
			}
			y, err := strconv.Atoi(splitted[1])
			if err != nil {
				return nil, err
			}
			task.Points[Point{X:x, Y:y}] = true
		}
	}

	maxY, maxX := 0, 0
	for p := range task.Points {
		if p.X > maxX {
			maxX = p.X
		}
		if p.Y > maxY {
			maxY = p.Y
		}
	}

	task.MaxY = maxY
	task.MaxX = maxX

	return task, nil
}

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}
