package main

import (
	"fmt"
	"sort"
	"strings"

	AOCUTILS "github.com/jnobre/advent-of-code-2022"
)

type Monkey struct {
	items       []int
	operation   string
	with        int
	divisibleBy int
	ifTrue      int
	ifFalse     int
	inspected   int
}

func main() {
	pt1()
	pt2()
}

func pt2() {
	input := AOCUTILS.ReadString("input.txt")
	monkes := make([]Monkey, 0)
	k := 1
	for _, mo := range strings.Split(input, "\n\n") {
		fields := AOCUTILS.SplitAndTrim(mo)
		items := make([]int, 0)
		items = append(items, AOCUTILS.ToInts(strings.Split(fields[1][16:], ", "))...)
		operation := ""
		with := 0
		if strings.Contains(fields[2], "* old") {
			operation = "mult"
		} else if strings.Contains(fields[2], "*") {
			operation = "mult"
			with = AOCUTILS.ToIntMust(fields[2][23:])
		} else {
			operation = "add"
			with = AOCUTILS.ToIntMust(fields[2][23:])
		}
		divisibleBy := AOCUTILS.ToIntMust(fields[3][19:])
		k *= divisibleBy
		ifTrue := AOCUTILS.ToIntMust(fields[4][25:])
		ifFalse := AOCUTILS.ToIntMust(fields[5][26:])
		monkes = append(monkes, Monkey{
			items:       items,
			operation:   operation,
			with:        with,
			divisibleBy: divisibleBy,
			ifTrue:      ifTrue,
			ifFalse:     ifFalse,
			inspected:   0,
		})
	}
	for i := 0; i < 10000; i += 1 {
		for j := 0; j < len(monkes); j += 1 {
			for {
				var item int
				var err error
				monkes[j].items, item, err = AOCUTILS.PopFront(monkes[j].items)
				if err != nil {
					break
				}
				monkes[j].inspected += 1
				switch monkes[j].operation {
				case "add":
					item += monkes[j].with
				case "mult":
					if monkes[j].with == 0 {
						item *= item
					} else {
						item *= monkes[j].with % monkes[j].divisibleBy
					}
				}
				// item %= monkes[j].divisibleBy
				if item%monkes[j].divisibleBy == 0 {
					monkes[monkes[j].ifTrue].items = AOCUTILS.PushBack(monkes[monkes[j].ifTrue].items, item%k)
				} else {
					monkes[monkes[j].ifFalse].items = AOCUTILS.PushBack(monkes[monkes[j].ifFalse].items, item%k)
				}
			}
		}
	}
	sort.Slice(monkes, func(i, j int) bool {
		return monkes[i].inspected > monkes[j].inspected
	})
	fmt.Println(monkes[0].inspected * monkes[1].inspected)
}

func pt1() {
	input := AOCUTILS.ReadString("input.txt")
	monkes := make([]Monkey, 0)
	for _, mo := range strings.Split(input, "\n\n") {
		fields := AOCUTILS.SplitAndTrim(mo)
		items := make([]int, 0)
		items = append(items, AOCUTILS.ToInts(strings.Split(fields[1][16:], ", "))...)
		operation := ""
		with := 0
		if strings.Contains(fields[2], "* old") {
			operation = "mult"
		} else if strings.Contains(fields[2], "*") {
			operation = "mult"
			with = AOCUTILS.ToIntMust(fields[2][23:])
		} else {
			operation = "add"
			with = AOCUTILS.ToIntMust(fields[2][23:])
		}
		divisibleBy := AOCUTILS.ToIntMust(fields[3][19:])
		ifTrue := AOCUTILS.ToIntMust(fields[4][25:])
		ifFalse := AOCUTILS.ToIntMust(fields[5][26:])
		monkes = append(monkes, Monkey{
			items:       items,
			operation:   operation,
			with:        with,
			divisibleBy: divisibleBy,
			ifTrue:      ifTrue,
			ifFalse:     ifFalse,
			inspected:   0,
		})
	}
	// fmt.Println(monkes)
	for i := 0; i < 20; i += 1 {
		for j := 0; j < len(monkes); j += 1 {
			for {
				var item int
				var err error
				monkes[j].items, item, err = AOCUTILS.PopFront(monkes[j].items)
				if err != nil {
					break
				}
				monkes[j].inspected += 1
				switch monkes[j].operation {
				case "add":
					item += monkes[j].with
				case "mult":
					if monkes[j].with == 0 {
						item *= item
					} else {
						item *= monkes[j].with
					}
				}
				item /= 3
				if item%monkes[j].divisibleBy == 0 {
					monkes[monkes[j].ifTrue].items = AOCUTILS.PushBack(monkes[monkes[j].ifTrue].items, item)
				} else {
					monkes[monkes[j].ifFalse].items = AOCUTILS.PushBack(monkes[monkes[j].ifFalse].items, item)
				}
			}
		}
	}
	sort.Slice(monkes, func(i, j int) bool {
		return monkes[i].inspected > monkes[j].inspected
	})
	fmt.Println(monkes[0].inspected * monkes[1].inspected)
}
