package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)
	orbits := make(map[string][]string)
	for reader.Scan() {
		input := strings.Split(reader.Text(), ")")
		prevName := input[0]
		currName := input[1]

		_, exists := orbits[prevName]
		if !exists {
			orbits[prevName] = []string{}
		}

		orbits[prevName] = append(orbits[prevName], currName)
		orbits[currName] = append(orbits[currName], prevName)

	}

	fmt.Println(findSSSP(orbits, "", "YOU", 0))
}

func findSSSP(m map[string][]string, from string, orbit string, count int) int {
	min := 10000000
	for _, value := range m[orbit] {
		if value == "SAN" {
			return count
		} else if value == from {
			continue
		}
		min = int(math.Min(float64(findSSSP(m, orbit, value, count+1)), float64(min)))
	}
	return min
}
