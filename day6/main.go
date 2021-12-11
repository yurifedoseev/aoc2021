package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)



func run() error {
	DAY_COUNT := 256

	sea, err := parseInput("day6/input.txt")
	if err != nil {
		return err
	}

	count := 0
	days := make([]int, DAY_COUNT)
	for _, fish := range sea {
		count++
		next := fish
		for next < DAY_COUNT  {
			days[next]++
			next += 7
		}
	}

	for i := 0; i < DAY_COUNT; i++ {
		fmt.Printf("\n day %d", i+1)
		fishCount := days[i]
		count += fishCount
		next := i + 9
		for next < DAY_COUNT  {
			days[next] += fishCount
			next += 7
		}
	}

	//for i, day := range days {
	//	fmt.Printf("\n after %d day: %v", i, day)
	//}

	fmt.Printf("\n %d fish", count)

	return nil
}


func parseInput(filename string) ([]int, error) {
	rawBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read input: %w", err)
	}

	fishes := make([]int, 0)
	for _, raw := range strings.Split(string(rawBytes), ",") {
		val, err := strconv.Atoi(raw)
		if err != nil {
			return nil, fmt.Errorf("failed to parse %s val to int: %v", raw, err)
		}
		fishes = append(fishes, val)
	}
	return fishes, nil
}

func main() {
	err := run()
	if err != nil {
		panic(err)
	}
}
