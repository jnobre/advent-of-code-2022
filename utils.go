package AOCUTILS

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ReadLines(filename string) []string {
	file, err := os.Open(filename)
	Check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	lines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	Check(scanner.Err())
	return lines
}

func SortDescending(a []int) []int {
	b := a[:]
	sort.Slice(b, func(i, j int) bool { return b[i] > b[j] })
	return b
}

func Sum(a []int) int {
	sum := 0
	for _, e := range a {
		sum += e
	}
	return sum
}

func ToInts(a []string) []int {
	return Map(a, func(x string) int {
		n, _ := strconv.Atoi(x)
		return n
	})
}

func Map[T any, S any](a []T, f func(x T) S) []S {
	r := make([]S, 0)
	for _, e := range a {
		r = append(r, f(e))
	}
	return r
}
