package AOCUTILS

import (
	"bufio"
	"errors"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
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

func ReadString(filename string) string {
	lines := ReadLines(filename)
	return strings.Join(lines, "\n")
}

func SplitAndTrim(s string) []string {
	return Map(strings.Split(s, "\n"), func(x string) string {
		return strings.TrimSpace(x)
	})
}

func SortDescending(a []int) []int {
	b := a[:]
	sort.Slice(b, func(i, j int) bool { return b[i] > b[j] })
	return b
}

func SortAscending(a []int) []int {
	b := a[:]
	sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
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

func Set[T comparable](a []T) []T {
	hash := make(map[T]uint8)
	for _, elem := range a {
		hash[elem] = 1
	}
	set := make([]T, len(hash))
	i := 0
	for k := range hash {
		set[i] = k
		i += 1
	}
	return set
}

func ToIntMust(a string) int {
	n, _ := strconv.Atoi(a)
	return n
}

func Max[T int | int8 | int16 | int32 | int64 | float32 | float64](xs []T) T {
	max := xs[0]
	for _, s := range xs {
		if s >= max {
			max = s
		}
	}
	return max
}

func Min[T int | int8 | int16 | int32 | int64 | float32 | float64](xs []T) T {
	min := xs[0]
	for _, s := range xs {
		if s <= min {
			min = s
		}
	}
	return min
}

func StrToDigits(a string) []int {
	l := len(a)
	res := make([]int, l)
	for i := range a {
		res[i] = ToIntMust(string(a[i]))
	}
	return res
}

func Abs(x int) int {
	if x < 0 {
		return (-x)
	}
	return x
}

func PushBack[T any](xs []T, s T) []T {
	return append(xs, s)
}

func PushFront[T any](xs []T, s T) []T {
	return append([]T{s}, xs...)
}

func PopBack[T any](xs []T) ([]T, T, error) {
	if len(xs) == 0 {
		var zero T
		return nil, zero, errors.New("array is empty")
	}
	return xs[:len(xs)-1], xs[len(xs)-1], nil
}

func PopFront[T any](xs []T) ([]T, T, error) {
	if len(xs) == 0 {
		var zero T
		return nil, zero, errors.New("array is empty")
	}
	return xs[1:], xs[0], nil
}
