package main

import (
	"fmt"

	AOCUTILS "github.com/jnobre/advent-of-code-2022"
)

func main() {
	lines := AOCUTILS.ReadLines("input.txt")[0]
	marker := 0
	for i := 4; i < len(lines); i += 1 {
		if len(AOCUTILS.Set([]byte(lines[i-4:i]))) == 4 {
			marker = i
			break
		}
	}
	fmt.Println(marker)
}
