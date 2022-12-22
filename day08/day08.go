package main

import (
	"fmt"

	AOCUTILS "github.com/jnobre/advent-of-code-2022"
)

func main() {
	lines := AOCUTILS.ReadLines("input.txt")
	l := len(lines)
	seen := len(lines[0])*4 - 4
	trees := make([][]int, len(lines))
	for i := range trees {
		trees[i] = AOCUTILS.StrToDigits(lines[i])
	}
	scenicScores := make([]int, l*l)
	for i := 1; i < l-1; i += 1 {
		for j := 1; j < l-1; j += 1 {
			left, up, right, down := true, true, true, true
			for m := 0; m < j; m += 1 {
				if trees[i][m] >= trees[i][j] {
					left = false
					break
				}
			}
			for m := l - 1; m > j; m -= 1 {
				if trees[i][m] >= trees[i][j] {
					right = false
					break
				}
			}
			for m := 0; m < i; m += 1 {
				if trees[m][j] >= trees[i][j] {
					down = false
					break
				}
			}
			for m := l - 1; m > i; m -= 1 {
				if trees[m][j] >= trees[i][j] {
					up = false
					break
				}
			}
			if left || down || right || up {
				seen += 1
			}
		}
	}
	for i := 1; i < l-1; i += 1 {
		for j := 1; j < l-1; j += 1 {
			left, up, right, down := 0, 0, 0, 0
			for m := i - 1; m >= 0; m -= 1 {
				up += 1
				if trees[m][j] >= trees[i][j] {
					break
				}
			}
			for m := j - 1; m >= 0; m -= 1 {
				left += 1
				if trees[i][m] >= trees[i][j] {
					break
				}
			}

			for m := i + 1; m < l; m += 1 {
				down += 1
				if trees[m][j] >= trees[i][j] {
					break
				}
			}
			for m := j + 1; m < l; m += 1 {
				right += 1
				if trees[i][m] >= trees[i][j] {
					break
				}
			}
			scenicScores = append(scenicScores, left*up*right*down)
		}
	}
	fmt.Println(seen)
	fmt.Println(AOCUTILS.Max(scenicScores))
}
