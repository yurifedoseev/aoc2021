package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strings"
)

type Task struct {
	Input string
	Combo map[string]string
	ComboRune map[RunePair]rune
}

type Progress struct {
	Combo map[string]string
	MaxComboLen int
}

type RunePair struct {
	First rune
	Second rune
}

func run() error {
	task, err := parse("day14/example.txt")
	if err != nil {
		return err
	}
	fmt.Printf("\n input %s, %d combos, %d rune combos", task.Input, len(task.Combo), len(task.ComboRune))

	silver(task)
	return nil
}

func parse(filename string)  (*Task, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	scan := bufio.NewScanner(f)
	task := &Task{
		Input: "",
		Combo: make(map[string]string),
		ComboRune: make(map[RunePair]rune),
	}

	scan.Scan()
	task.Input = scan.Text()
	scan.Scan()
	for scan.Scan() {
		splitted := strings.Split(scan.Text(), " -> ")
		task.Combo[splitted[0]] = splitted[1]
		runePairs := []rune(splitted[0])
		pair := RunePair{
			First: runePairs[0],
			Second: runePairs[1],
		}
		task.ComboRune[pair] = []rune(splitted[1])[0]
	}
	return task, nil
}

func silver(task *Task) {
	days := 15
	polymer := task.Input
	for d := 1; d <= days; d++ {
		polymer = growPolymer(polymer, task.Combo)
		fmt.Printf("\n day %d, len %d", d, len(polymer))
		charCount := make(map[string]int)
		for _, ch := range polymer {
			charCount[string(ch)]++
		}
		countMinMax(charCount)
	}
}

func goldV4(task *Task) {
	days := 40
	polyLen := len(task.Input)
	for d := 1; d <= days; d++ {
		nextLen := polyLen * 2 - 1
		fmt.Printf("\n day %d len %d", d, nextLen)
		polyLen = nextLen
	}
}

func countMinMax(charCount map[string]int) {
	minChar, maxChar := "N", "N"
	for k,v := range charCount {
		if v < charCount[minChar] {
			minChar = k
		}
		if v > charCount[maxChar] {
			maxChar = k
		}
	}

	fmt.Printf("\n max char %s - %d, min char %s - %d, sub: %d",
		maxChar, charCount[maxChar], minChar, charCount[minChar], charCount[maxChar] - charCount[minChar])
}

func growPolymer(input string, combos map[string]string) string {
	poly := ""
	runeInput := []rune(input)
	for i, val := range runeInput {
		ch := string(val)
		if i == len(runeInput) - 1 {
			poly += ch
			continue
		}
		poly += ch
		pair := ch +  string(runeInput[i+1])
		if insertCh, ok := combos[pair]; ok {
			poly += insertCh
		}
	}
	return poly
}

func goldV2(task *Task) {
	combo10 := make(map[string]string)
	for startPair := range task.Combo {
		polymer := startPair
		for d := 1; d <= 10; d++ {
			polymer = growPolymer(polymer, task.Combo)
		}
		combo10[startPair] = polymer[1:len(polymer)-1]
	}
	fmt.Printf("\n calculated combo 10")

	combo20 := make(map[string]string)
	for startPair := range combo10 {
		polymer := startPair
		for d := 10; d <= 20; d+=10 {
			polymer = growPolymer(polymer, combo10)
		}
		combo20[startPair] = polymer[1:len(polymer)-1]
	}

	fmt.Printf("\n calculated combo 20")

	days := 40

	polymer := task.Input
	for d := 20; d <= days; d += 20 {
		polymer = growPolymer(polymer, combo20)
		fmt.Printf("\n day %d, len %d", d, len(polymer))
	}

	charCount := make(map[string]int)
	for _, ch := range polymer {
		charCount[string(ch)]++
	}

	countMinMax(charCount)
}

func goldV3(task *Task) {
	days := 10
	
	polymer := task.Input
	for d := 1; d <= days; d++ {
		polymer = growPolymer(polymer,  task.Combo)
		fmt.Printf("\n day %d, len %d", d, len(polymer))
	}

	charCount := make(map[string]int)
	for _, ch := range polymer {
		charCount[string(ch)]++
	}

	countMinMax(charCount)
	//charCount := make(map[string]int)
	//for _, ch := range polymer {
	//	charCount[string(ch)]++
	//}
	//
	//countMinMax(charCount)
}

func gold(task *Task) {
	days := 40
	polymer := list.New()
	for _, val := range task.Input {
		polymer.PushBack(val)
	}

	for d := 1; d <= days; d++ {
		growPolymerList(polymer, task.ComboRune)
		fmt.Printf("\n day %d, poly len %d", d, polymer.Len())
	}

	fmt.Printf("\n poly:")
	charCount := make(map[rune]int)
	for e := polymer.Front(); e != nil; e = e.Next() {
		val :=  e.Value.(rune)
		charCount[val]++
	}

	countMinMaxRune(charCount)
}

func countMinMaxRune(charCount map[rune]int) {
	minChar, maxChar := []rune("N")[0], []rune("N")[0]
	for k,v := range charCount {
		if v < charCount[minChar] {
			minChar = k
		}
		if v > charCount[maxChar] {
			maxChar = k
		}
	}

	fmt.Printf("\n max char %s - %d, min char %s - %d, sub: %d",
		string(maxChar), charCount[maxChar], string(minChar), charCount[minChar], charCount[maxChar] - charCount[minChar])
}

func growPolymerList(poly *list.List, combos map[RunePair]rune) {
	e := poly.Front()
	for e != nil {
		nextE := e.Next()
		if nextE == nil {
			break
		}

		pair := RunePair{
			First:  e.Value.(rune),
			Second: nextE.Value.(rune),
		}
		if insertCh, ok := combos[pair]; ok {
			poly.InsertAfter(insertCh, e)
		}
		e = nextE
	}
}



func main() {
	if err := run(); err != nil {
		panic(err)
	}
}
