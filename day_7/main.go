package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const initSig int = 0

func main() {
	// Reads the intCode program
	file, _ := os.Open("input.in")
	defer file.Close()
	reader := bufio.NewScanner(file)
	var input []string
	for reader.Scan() {
		input = strings.Split(reader.Text(), ",")
	}
	formatted := make([]int, len(input))
	for i, v := range input {
		formatted[i], _ = strconv.Atoi(v)
	}

	// Gets all permutations of control codes
	var sig, max int
	temp := make([]int, len(formatted))
	perm := Permute([]int{0, 1, 2, 3, 4})
	fmt.Println(perm)
	for i := 0; i < len(perm); i++ {
		conSeq := perm[i]
		for j := 0; j < 5; j++ {
			copy(temp, formatted)
			con := conSeq[j]
			if i == 0 {
				sig = intCode(temp, con, initSig)
			} else {
				sig = intCode(temp, con, sig)
			}
		}
		max = int(math.Max(float64(sig), float64(max)))
		sig = 0
	}
	fmt.Println("Output:", max)
}

func intCode(temp []int, con int, sig int) int {
	isCon := true
Intcodeloop:
	for i := 0; i < len(temp); {
		_, b, c, opcode := readOpcode(temp[i])
		if opcode == 99 {
			break
		}
		switch opcode {
		case 1:
			handle1(temp, i+1, b, c)
			i += 4
		case 2:
			handle2(temp, i+1, b, c)
			i += 4
		case 3:
			if isCon {
				handle3(temp, i+1, con)
			} else {
				handle3(temp, i+1, sig)
			}
			isCon = !isCon
			i += 2
		case 4:
			return handle4(temp, i+1, c)
		case 5:
			i = handle5(temp, i+1, b, c)
		case 6:
			i = handle6(temp, i+1, b, c)
		case 7:
			handle7(temp, i+1, b, c)
			i += 4
		case 8:
			handle8(temp, i+1, b, c)
			i += 4
		default:
			fmt.Println("default", opcode)
			break Intcodeloop
		}
	}
	return 0
}

func readOpcode(num int) (a, b, c, op int) {
	var cycle int
	for num > 0 {
		switch cycle {
		case 0:
			op = num % 100
			num = num / 100
		case 1:
			c = num % 10
			num = num / 10
		case 2:
			b = num % 10
			num = num / 10
		case 3:
			a = num % 10
			num = num / 10
		}
		cycle++
	}
	return
}

func handleArith(op func(int, int) int, arr []int, i int, b int, c int) {
	var (
		first  int
		second int
		pos    int
	)
	if c == 1 {
		first = arr[i]
	} else {
		first = arr[arr[i]]
	}
	if b == 1 {
		second = arr[i+1]
	} else {
		second = arr[arr[i+1]]
	}
	pos = arr[i+2]

	arr[pos] = op(first, second)
}

func handle1(arr []int, i int, b int, c int) {
	handleArith(func(x, y int) int { return x + y }, arr, i, b, c)
}

func handle2(arr []int, i int, b int, c int) {
	handleArith(func(x, y int) int { return x * y }, arr, i, b, c)
}

func handle3(arr []int, i int, sig int) {
	pos := arr[i]
	arr[pos] = sig
}

func handle4(arr []int, i int, c int) int {
	var pos int
	if c == 1 {
		pos = arr[i]
	} else {
		pos = arr[i]
	}
	fmt.Println(arr[pos])
	return arr[pos]
}

func handle5(arr []int, i int, b int, c int) int {
	var (
		first  int
		second int
	)
	if c == 1 {
		first = arr[i]
	} else {
		first = arr[arr[i]]
	}
	if b == 1 {
		second = arr[i+1]
	} else {
		second = arr[arr[i+1]]
	}
	if first != 0 {
		return second
	} else {
		return i + 2
	}
}

func handle6(arr []int, i int, b int, c int) int {
	var (
		first  int
		second int
	)
	if c == 1 {
		first = arr[i]
	} else {
		first = arr[arr[i]]
	}
	if b == 1 {
		second = arr[i+1]
	} else {
		second = arr[arr[i+1]]
	}
	if first == 0 {
		return second
	} else {
		return i + 2
	}
}

func handle7(arr []int, i int, b int, c int) {
	var (
		first  int
		second int
		third  int
	)
	if c == 1 {
		first = arr[i]
	} else {
		first = arr[arr[i]]
	}
	if b == 1 {
		second = arr[i+1]
	} else {
		second = arr[arr[i+1]]
	}
	third = arr[i+2]
	if first < second {
		arr[third] = 1
	} else {
		arr[third] = 0
	}
}

func handle8(arr []int, i int, b int, c int) {
	var (
		first  int
		second int
		third  int
	)
	if c == 1 {
		first = arr[i]
	} else {
		first = arr[arr[i]]
	}
	if b == 1 {
		second = arr[i+1]
	} else {
		second = arr[arr[i+1]]
	}
	third = arr[i+2]
	if first == second {
		arr[third] = 1
	} else {
		arr[third] = 0
	}
}
