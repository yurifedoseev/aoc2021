package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func run() error {
	lines, err := parse("day8/input.txt")
	if err != nil {
		return err
	}

	fmt.Printf("\n count: %d", countAll(lines))

	return nil
}

func countAll(lines []*Line) int {
	sum := 0

	for _, l := range lines {
		lineSum := countLine(l)
		fmt.Printf("\nline %v sum is %d", l.Digits, lineSum)
		sum += lineSum
	}
	return sum
}

func countLine(line *Line) int {
	one, four, seven, eight := "", "", "", ""

	mapping := &Mapping{
		A:  "",
		CF: make(map[string]bool),
		BD: make(map[string]bool),
	}

	allInput := append(line.Examples, line.Digits...)
	for _, d := range allInput {
		switch len(d) {
		case 2:
			//fmt.Printf("\n%s is 1 with len 2", d)
			one = d
		case 3:
			//fmt.Printf("\n%s is 7 with len 3", d)
			seven = d
		case 4:
			// fmt.Printf("\n%s is 4 with len 4", d)
			four = d
		case 7:
			//fmt.Printf("\n%s is 8 with len 7", d)
			eight = d
		}
	}

	fmt.Printf("\none: %s, four: %s, seven: %s, eight: %s", one, four, seven, eight)

	if one == "" || four == "" || seven == "" || eight == "" {
		panic("упс!")
	}

	// 7 - acf, 1 = cf
	for _, el := range seven {
		if !strings.Contains(one, string(el)) {
			fmt.Printf("\n find %s is a element in 7", string(el))
			mapping.A = string(el)
			break
		}
	}

	for _, el := range one {
		mapping.CF[string(el)] = true
	}

	for _, el := range four {
		if _, ok := mapping.CF[string(el)]; !ok {
			mapping.BD[string(el)] = true
		}
	}

	sum := 0
	counter := 1000
	for _, val := range line.Digits {
		dig := countDigit(mapping, val)
		sum += dig * counter
		counter = counter / 10
	}

	return sum
}

func countDigit(m *Mapping, digit string) int {
	// try simple
	switch len(digit) {
	case 2:
		return 1
	case 3:
		return 7
		//fmt.Printf("\n%s is 7 with len 3", d)
	case 4:
		return 4
	case 7:
		return 8
	}

	cfCount, bdCount := 0, 0
	aCount := 0
	for _, el := range digit {
		val := string(el)
		if val == m.A {
			aCount++
			continue
		}
		if m.BD[val] {
			bdCount++
		}
		if m.CF[val] {
			cfCount++
		}
	}

	if len(digit) == 6 {
		// 0 - len 6
		// 6 - len 6
		// 9 - len 6
		if cfCount == 1 && bdCount == 2 	{
			return 6
		}

		if bdCount == 1 && cfCount == 2 {
			return 0
		}
		return 9
	}

	if len(digit) == 5 {
		// 2 - len 5
		// 3 - len 5
		// 5 - len 5
		if cfCount == 2 && bdCount == 1 {
			return 3
		}
		if cfCount == 1 && bdCount == 1 {
			return 2
		}
		return 5
	}

	panic("unknown digit " + digit)
}

type Mapping struct {
	A string
	CF map[string]bool
	BD map[string] bool
}

func countSimple(lines []*Line) int {
	count := 0
	for _, l := range lines {
		for _, d := range l.Digits {
			switch len(d) {
			case 2:
				fmt.Printf("\n%s is 1 with len 2", d)
				count++
			case 3:
				fmt.Printf("\n%s is 7 with len 3", d)
				count++
			case 4:
				fmt.Printf("\n%s is 4 with len 4", d)
				count++
			case 7:
				fmt.Printf("\n%s is 8 with len 7", d)
				count++
			}

		}
	}
	return count
}

type Line struct {
	Examples []string
	Digits []string
}

func parse(filename string) ([]*Line, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	lines := make([]*Line, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := &Line{
			Examples: make([]string, 0),
			Digits:   make([]string, 0),
		}

		raw := scanner.Text()
		parts := strings.Split(raw, " | ")
		for _, val := range strings.Split(parts[0], " ") {
			trimmed := strings.TrimSpace(val)
			if trimmed == "" {
				continue
			}
			line.Examples = append(line.Examples, trimmed)
		}
		for _, val := range strings.Split(parts[1], " ") {
			trimmed := strings.TrimSpace(val)
			if trimmed == "" {
				continue
			}
			line.Digits = append(line.Digits, trimmed)
		}
		lines = append(lines, line)
	}

	return lines, nil
}

func main() {
	err := run()
	if err != nil {
		panic(err)
	}
}
