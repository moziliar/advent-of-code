package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lower := 359282
	upper := 820401

	var count int
	for i := lower; i <= upper; i++ {
		if accept(i) {
			count++
		}
	}
	fmt.Println(count)
}

func accept(num int) bool {
	m := make(map[string]int)
	strArr := strings.Split(strconv.Itoa(num), "")
	dup := false
	for _, v := range strArr {
		_, exist := m[v]
		if !exist {
			m[v] = 1
		} else {
			m[v] = m[v] + 1
		}
	}
	for _, v := range m {
		if v == 2 {
			dup = true
			break
		}
	}
	var prev string
	for i, v := range strArr {
		if i == 0 {
			prev = v
			continue
		}
		prevNum, _ := strconv.Atoi(prev)
		currNum, _ := strconv.Atoi(v)
		if prevNum > currNum {
			dup = false
			break
		}
		prev = v
	}
	return dup
}
