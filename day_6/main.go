package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)
	orbits := make(map[string]string)
	for reader.Scan() {
		input := strings.Split(reader.Text(), ")")
		prevName := input[0]
		currName := input[1]

		_, exists := orbits[prevName]
		if !exists {
			orbits[prevName] = "nil"
		}

		orbits[currName] = prevName
	}

	fmt.Println(orbits)
	var count int
	for key, _ := range orbits {
		count += recCount(orbits, key)
	}
	fmt.Println(count)
}

func recCount(m map[string]string, orbit string) int {
	if orbit == "nil" || m[orbit] == "nil" {
		return 0
	}
	return 1 + recCount(m, m[orbit])
}
