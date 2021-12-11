package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func run() error {
	lines, err := parse("day10/input.txt")
	if err != nil {
		return err
	}

	pointsMap := map[string]int{
		")": 1,
		"]": 2,
		"}": 3,
		">": 4,
	}

	closePair := map[string]string {
		")": "(",
		"]": "[",
		"}":"{",
		">":"<",
	}

	startPair := map[string]string {
		"(": ")",
		"[": "]",
		"{":"}",
		"<":">",
	}

	scores := make([]int, 0)
	for k, line := range lines {
		stack := make([]string, 0)
		hasCorrupted := false
		for i, runeEl := range line {
			el := string(runeEl)
			if el == "{" || el == "(" || el == "[" || el == "<" {
				stack = append([]string{el}, stack...)
			} else {
				if len(stack) == 0 {
					panic(fmt.Sprintf("dont expect start from bad char on line %d", i))
				}
				startChar := closePair[el]
				var lastClose string
				lastClose, stack = stack[0], stack[1:]
				if startChar != lastClose {
					fmt.Printf("\n skip corrupted line: %d", i)
					// p := pointsMap[el]
					// fmt.Printf("\n error on line %d: el %s, expected %s not equal last closing %s, points: %d", i, el, startChar, lastClose, p)
					// score += p
					hasCorrupted = true
					break
				}
			}
		}

		if hasCorrupted {
			continue
		}

		lineScore := 0
		lineAdd := ""
		for _, el := range stack {
			closeChar := startPair[el]
			lineAdd = closeChar
			p := pointsMap[closeChar]
			lineScore = 5 * lineScore + p
		}
		scores = append(scores, lineScore)
		fmt.Printf("\n add %s closing to line %d for points, score: %d", lineAdd, k, lineScore)
	}

	sort.Ints(scores)
	score := scores[len(scores) / 2]
	fmt.Printf("\n total score: %d", score)
	return nil
}

func parse(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	lines := make([]string, 0)
	scan := bufio.NewScanner(f)
	for scan.Scan() {
		lines = append(lines, scan.Text())
	}
	return lines, nil
}

func main() {
	err := run()
	if err != nil {
		panic(err)
	}
}
