package main

import (
	"fmt"
	"strings"

	AOCUTILS "github.com/jnobre/advent-of-code-2022"
)

func main() {
	lines := AOCUTILS.ReadLines("input.txt")
	contained := 0
	overlapped := 0

	for _, line := range lines {
		sections := strings.Split(line, ",")
		e1, e2 := AOCUTILS.ToInts(strings.Split(sections[0], "-")), AOCUTILS.ToInts(strings.Split(sections[1], "-"))
		if (e1[0] >= e2[0] && e1[1] <= e2[1]) || (e1[0] <= e2[0] && e1[1] >= e2[1]) {
			contained += 1
		}

		if !(e1[1] < e2[0] || e2[1] < e1[0]) {
			overlapped += 1
		}
	}
	fmt.Println(contained)
	fmt.Println(overlapped)
}
