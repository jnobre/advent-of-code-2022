package main

import (
	"fmt"
	"strings"

	AOCUTILS "github.com/jnobre/advent-of-code-2022"
)

func main() {
	lines := AOCUTILS.ReadLines("input.txt")
	problem1(lines)
	problem2(lines)
}

func problem1(lines []string) {
	crates := make([][]string, 9)
	cargos := make([]string, 0)
	opstart := 0
	for i, line := range lines {
		if line == " 1   2   3   4   5   6   7   8   9 " {
			opstart = i + 2
			break
		}
		cargos = append(cargos, line)
	}
	for i := len(cargos) - 1; i >= 0; i -= 1 {
		line := cargos[i]
		items := splitBy(line, 3)
		for j := range items {
			if items[j] != "   " {
				crates[j] = append(crates[j], string(items[j][1]))
			}
		}
	}
	for i := opstart; i < len(lines); i += 1 {
		amt, from, to := parseLine(lines[i])
		for k := 0; k < amt; k += 1 {
			t, c := pop(crates[from])
			crates[from] = t
			crates[to] = push(crates[to], c)
		}
	}
	tops := make([]string, 0)
	for _, crate := range crates {
		if len(crate) > 0 {
			tops = append(tops, string(crate[len(crate)-1]))

		}
	}
	fmt.Println(strings.Join(tops, ""))

}

func problem2(lines []string) {
	crates := make([]string, 9)
	cargos := make([]string, 0)
	opstart := 0
	for i, line := range lines {
		if line == " 1   2   3   4   5   6   7   8   9 " {
			opstart = i + 2
			break
		}
		cargos = append(cargos, line)
	}
	for i := len(cargos) - 1; i >= 0; i -= 1 {
		line := cargos[i]
		items := splitBy(line, 3)
		for j := range items {
			if items[j] != "   " {
				crates[j] += string(items[j][1])
			}
		}
	}
	for i := opstart; i < len(lines); i += 1 {
		amt, from, to := parseLine(lines[i])
		l := len(crates[from])
		part := crates[from][l-amt:]
		crates[from] = crates[from][:l-amt]
		crates[to] = crates[to] + part
	}
	tops := make([]string, 0)
	for _, crate := range crates {
		if len(crate) > 0 {
			tops = append(tops, string(crate[len(crate)-1]))

		}
	}
	fmt.Println(strings.Join(tops, ""))
}

func splitBy(s string, steps int) []string {
	res := make([]string, 0)
	for i := 0; i < len(s); i += steps + 1 {
		res = append(res, s[i:i+steps])
	}
	return res
}

func push(s []string, t string) []string {
	return append(s, t)
}

func pop(s []string) ([]string, string) {
	if len(s) == 0 {
		return []string{}, ""
	}
	item := s[len(s)-1]
	return s[:len(s)-1], item
}

func parseLine(line string) (amt, from, to int) {
	parts := strings.Split(line, " ")
	parsedParts := AOCUTILS.ToInts([]string{parts[1], parts[3], parts[5]})
	amt, from, to = parsedParts[0], parsedParts[1]-1, parsedParts[2]-1
	return
}
