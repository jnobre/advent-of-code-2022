package AOCUTILS

import (
	"bufio"
	"log"
	"os"
	"sort"
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
