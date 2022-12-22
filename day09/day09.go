package main

import (
	"fmt"
	"strings"

	AOCUTILS "github.com/jnobre/advent-of-code-2022"
)

func main() {
	problem1()
	problem2()
}

func problem1() {
	lines := AOCUTILS.ReadLines("input.txt")
	grid := make([][]int, 512)
	for i := range grid {
		grid[i] = make([]int, 512)
	}
	taily, tailx, heady, headx := 300, 256, 300, 256 // :)
	grid[taily][tailx] = 1
	for _, line := range lines {
		parts := strings.Split(line, " ")
		amt := AOCUTILS.ToIntMust(parts[1])
		for i := 0; i < amt; i += 1 {
			switch parts[0] {
			case "R":
				headx++
			case "L":
				headx--
			case "U":
				heady--
			case "D":
				heady++
			}
			if !isTouching(headx, heady, tailx, taily) {
				if heady > taily {
					taily++
				}
				if heady < taily {
					taily--
				}
				if headx > tailx {
					tailx++
				}
				if headx < tailx {
					tailx--
				}
			}
			grid[taily][tailx] = 1
		}
	}
	fmt.Println(countVisited(grid))
}

func problem2() {
	lines := AOCUTILS.ReadLines("input.txt")
	grid := make([][]int, 512)
	for i := range grid {
		grid[i] = make([]int, 512)
	}
	rope := make([][2]int, 10)
	for i := range rope {
		rope[i] = [2]int{300, 256}
	}
	grid[rope[0][0]][rope[0][1]] = 1
	for _, line := range lines {
		parts := strings.Split(line, " ")
		amt := AOCUTILS.ToIntMust(parts[1])
		for i := 0; i < amt; i += 1 {
			switch parts[0] {
			case "R":
				rope[9][1]++
			case "L":
				rope[9][1]--
			case "U":
				rope[9][0]--
			case "D":
				rope[9][0]++
			}
			for i := 0; i < 9; i += 1 {
				if !isTouching(rope[i+1][1], rope[i+1][0], rope[i][1], rope[i][0]) {
					if rope[i+1][0] > rope[i][0] {
						rope[i][0]++
					}
					if rope[i+1][0] < rope[i][0] {
						rope[i][0]--
					}
					if rope[i+1][1] > rope[i][1] {
						rope[i][1]++
					}
					if rope[i+1][1] < rope[i][1] {
						rope[i][1]--
					}
				}
			}
			grid[rope[0][0]][rope[0][1]] = 1
		}
	}
	fmt.Println(countVisited(grid))
}

func countVisited(grid [][]int) int {
	visited := 0
	for _, row := range grid {
		for _, elem := range row {
			if elem == 1 {
				visited += 1
			}
		}
	}
	return visited
}

func isTouching(headx, heady, tailx, taily int) bool {
	xdist := AOCUTILS.Abs(headx - tailx)
	if xdist > 1 {
		return false
	}
	ydist := AOCUTILS.Abs(heady - taily)
	return ydist <= 1
}
