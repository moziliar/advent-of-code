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
	noun := 0
	verb := 0
	for intCode(formatted, noun, verb) != 19690720 {
		fmt.Println(intCode(formatted, noun, verb), noun, verb)
		if noun > 99 {
			break
		}
		if verb >= 99 {
			noun++
			verb = 0
		}
		verb++
	}
	fmt.Println(intCode(formatted, noun, verb), noun, verb)
}

func intCode(input []int, noun int, verb int) int {
	temp := make([]int, len(input))
	copy(temp, input)
	temp[1] = noun
	temp[2] = verb
	for i := 0; i < len(temp); i += 4 {
		opcode := temp[i]
		if opcode == 99 {
			return temp[0]
		}
		first := temp[temp[i+1]]
		second := temp[temp[i+2]]
		pos := temp[i+3]
		if pos > len(temp)+1 {
			return temp[0]
		}
		if opcode == 1 {
			temp[pos] = first + second
		} else if opcode == 2 {
			temp[pos] = first * second
		}
	}
	return temp[0]
}
