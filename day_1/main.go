package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// fuel()
	advFuel()
}

// fuel calculates the amount of fuel required for part 1
func fuel() {
	reader := bufio.NewScanner(os.Stdin)
	fuel := 0
	for reader.Scan() {
		i, _ := strconv.Atoi(reader.Text())
		fuel += i/3 - 2
	}

	fmt.Println(fuel)
}

// advFuel calculates the amount of fuel required for part 2
func advFuel() {
	reader := bufio.NewScanner(os.Stdin)
	fuel := 0
	for reader.Scan() {
		i, _ := strconv.Atoi(reader.Text())
		fuel += recFuel(i)
	}
	fmt.Println(fuel)
}

// recFuel recursively calculates the amount of fuel required by the fuel
func recFuel(fuel int) int {
	fuelReq := fuel/3 - 2
	if fuelReq <= 0 {
		return 0
	}
	return fuelReq + recFuel(fuelReq)
}
