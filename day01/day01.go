package main

import (
	"fmt"
	"strconv"

	AOCUTILS "github.com/jnobre/advent-of-code-2022"
)

func main() {
	lines := AOCUTILS.ReadLines("input.txt")
	groups := make([]int, 0)
	curr := 0
	for _, el := range lines {
		num, err := strconv.Atoi(el)
		if err != nil {
			groups = append(groups, curr)
			curr = 0
		} else {
			curr += num
		}
	}
	groups = append(groups, curr)
	groups = AOCUTILS.SortDescending(groups)
	fmt.Println(groups[0])
	fmt.Println(AOCUTILS.Sum(groups[:3]))
}
