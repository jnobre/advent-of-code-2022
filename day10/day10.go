package main

import (
	"fmt"

	AOCUTILS "github.com/jnobre/advent-of-code-2022"
)

func main() {
	lines := AOCUTILS.ReadLines("input.txt")
	x, signal_strength, important, amt, cc, blocked := 1, 0, 20, 0, 1, false
	crt := make([][]string, 0)
	for i := 0; i < 6; i += 1 {
		crt = append(crt, make([]string, 40))
	}
	for i := 0; i < len(lines); {
		if cc == important {
			signal_strength += x * cc
			important += 40
		}
		crt_x := (cc - 1) % 40
		crt_y := (cc - 1) / 40
		if inSprite(crt_x, x) {
			crt[crt_y][crt_x] = "#"
		} else {
			crt[crt_y][crt_x] = "."
		}
		if blocked {
			blocked = false
			x += amt
			i += 1
		} else {
			if lines[i] == "noop" {
				i += 1
			} else {
				amt = AOCUTILS.ToIntMust(lines[i][5:])
				blocked = true
			}
		}
		cc += 1
	}
	fmt.Println(signal_strength)
	for i := 0; i < 6; i += 1 {
		for j := 0; j < 40; j += 1 {
			fmt.Print(crt[i][j])
		}
		fmt.Println()
	}
}

func inSprite(x int, sprite int) bool {
	return x == (sprite-1) || x == sprite || x == (sprite+1)
}
