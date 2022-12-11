package main

import (
	"fmt"

	AOCUTILS "github.com/jnobre/advent-of-code-2022"
)

func main() {
	lines := AOCUTILS.ReadLines("input.txt")
	scores1 := map[string]int{
		"A Y": 8,
		"A X": 4,
		"A Z": 3,
		"B Y": 5,
		"B X": 1,
		"B Z": 9,
		"C Y": 2,
		"C X": 7,
		"C Z": 6,
	}
	scores2 := map[string]int{
		"A Y": 4,
		"A X": 3,
		"A Z": 8,
		"B Y": 5,
		"B X": 1,
		"B Z": 9,
		"C Y": 6,
		"C X": 2,
		"C Z": 7,
	}
	score1 := 0
	score2 := 0
	for _, line := range lines {
		score1 += scores1[line]
		score2 += scores2[line]
	}

	fmt.Println(score1, score2)

}
