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
	intCode(formatted)
}

func intCode(temp []int) {
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
			handle3(temp, i+1)
			i += 2
		case 4:
			handle4(temp, i+1, c)
			i += 2
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

func handle3(arr []int, i int) {
	pos := arr[i]
	arr[pos] = 5
}

func handle4(arr []int, i int, c int) {
	var pos int
	if c == 1 {
		pos = arr[i]
	} else {
		pos = arr[i]
	}
	fmt.Println(arr[pos])
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
