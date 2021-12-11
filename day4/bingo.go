package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Board struct {
	Numbers [][]int
	Marked [][]bool
	WinIndex int
	WinPoints int
	Index int
}

func (b *Board) Mark(i,j int) bool {
	b.Marked[i][j] = true
	allTrue := true

	// check row
	for _, el := range b.Marked[i] {
		allTrue = allTrue && el
	}
	if allTrue {
		return true
	}

	// check column
	allTrue = true
	for _, line := range b.Marked {
		allTrue = allTrue && line[j]
	}

	return allTrue
}


func (b *Board) UnmarkedSum() int {
	sum := 0
	for i, line := range b.Numbers {
		for j, val := range line {
			if !b.Marked[i][j] {
				sum += val
			}
		}
	}
	return sum
}

type DigitPoint struct {
	Board *Board
	I int
	J int
}

type Game struct {
	Digits map[int][]*DigitPoint
	Boards []*Board
	Inputs []int
	WinIndex int
}

func (g *Game) Play() int {
	for _, val := range g.Inputs {
		g.ProcessInput(val)
	}

	lastBoard := g.Boards[0]
	for _, board := range g.Boards {
		if board.WinIndex > lastBoard.WinIndex {
			lastBoard = board
		}
	}

	fmt.Printf("\n win index %d on board %d", lastBoard.WinIndex, lastBoard.Index)
	return lastBoard.WinPoints
}

func (g *Game) ProcessInput(val int)  {
	for _, p := range g.Digits[val] {
		if p.Board.WinIndex > 0 {
			continue
		}

		if p.Board.Mark(p.I, p.J) {
			fmt.Printf("\nsum %d won on board %d", val, p.Board.Index)
			g.WinIndex++
			p.Board.WinIndex = g.WinIndex
			p.Board.WinPoints = val * p.Board.UnmarkedSum()
		}
	}
}

func run() error {
	game, err := readLines("day4/input.txt")
	if err != nil {
		return err
	}

	fmt.Println("parsed!")
	final := game.Play()
	fmt.Printf("\nwon is %d", final)
	return nil
}

func readLines(filename string) (*Game, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	game := &Game{
		Inputs: make([]int, 0),
		Boards: make([]*Board, 0),
		Digits: make(map[int][]*DigitPoint),
	}

	scanner := bufio.NewScanner(f)
	var board *Board
	for scanner.Scan() {
		raw := strings.TrimSpace(scanner.Text())
		if len(game.Inputs) == 0 {
			for _, r := range strings.Split(raw, ",") {
				val, err := strconv.Atoi(r)
				if err != nil {
					return nil, fmt.Errorf("failed to convert %s to int: %w", r, err)
				}
				game.Inputs = append(game.Inputs, val)
			}
			continue
		}
		if raw == "" {
			board = &Board{
				Index : len(game.Boards),
				Numbers: make([][]int, 0),
				Marked: make([][]bool, 0),
			}
			game.Boards = append(game.Boards, board)
			continue
		}

		nums := make([]int, 0)
		marked := make([]bool, 0)
		i := len(board.Numbers)
		j := 0
		for _, r := range strings.Split(raw, " ") {
			trimmed := strings.TrimSpace(r)
			if trimmed == "" {
				continue
			}
			val, err := strconv.Atoi(trimmed)
			if err != nil {
				return nil, fmt.Errorf("failed to convert %s to int: %w", r, err)
			}

			nums = append(nums, val)
			marked = append(marked, false)
			game.Digits[val] = append(game.Digits[val], &DigitPoint{
				Board: board,
				I: i,
				J: j,
			})
			j++
		}
		board.Numbers = append(board.Numbers, nums)
		board.Marked = append(board.Marked, marked)
	}

	return game, nil
}

func main() {
	err := run()
	if err != nil {
		panic(err)
	}
}
