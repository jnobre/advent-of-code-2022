package main

import (
	"fmt"

	AOCUTILS "github.com/jnobre/advent-of-code-2022"
)

func main() {
	lines := AOCUTILS.ReadLines("input.txt")[0]
	marker, message := 0, 0
	for i := 4; i < len(lines); i += 1 {
		if len(AOCUTILS.Set([]byte(lines[i-4:i]))) == 4 {
			marker = i
			break
		}
	}
	for i := 14; i < len(lines); i += 1 {
		if len(AOCUTILS.Set([]byte(lines[i-14:i]))) == 14 {
			message = i
			break
		}
	}
	fmt.Println(marker)
	fmt.Println(message)
}
