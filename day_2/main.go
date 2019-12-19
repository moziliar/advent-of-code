package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)
	reader.Scan()
	input := strings.Split(reader.Text(), ",")
	formatted := make([]int, len(input))
	for i, v := range input {
		formatted[i], _ = strconv.Atoi(v)
	}
	fmt.Println(intCode(formatted))
}

func intCode(input []int) []int {
	for i := 0; i < len(input); i += 4 {
		opcode := input[i]
		if opcode == 99 {
			return input
		}
		first := input[input[i + 1]]
		second := input[input[i + 2]]
		pos := input[i + 3]
		if opcode == 1 {
			input[pos] = first + second
		} else if opcode == 2 {
			input[pos] = first * second
		}
	}
	return input
}


