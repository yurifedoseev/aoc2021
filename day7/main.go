package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func run() error {
	positions, err := parse("day7/example.txt")
	if err != nil {
		return err
	}

	sort.Ints(positions)
	fmt.Printf("\nprinted input: %v", positions)

	pin, pinCount := positions[0], 1
	currMult := 1
	for i := 1; i < len(positions); i++ {
		val := positions[i]
		if val == pin {
			currMult = 1
			pinCount++
			continue
		}

		if positions[i] == positions[i-1] {
			currMult++
			continue
		}
	}

	fmt.Printf("\npin is %d", pin)

	return nil
}

func parse(filename string) ([]int, error) {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}
	result := make([]int, 0)
	for _, t := range strings.Split(string(raw), ",") {
		val, err := strconv.Atoi(t)
		if err != nil {
			return nil, err
		}
		result = append(result, val)
	}

	return result, nil
}


func main(){
	err := run()
	if err != nil {
		panic(err)
	}
}
