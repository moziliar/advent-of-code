package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fuel()
}

func fuel() {
	reader := bufio.NewScanner(os.Stdin)
	fuel := 0;
	for reader.Scan(){
		i, _ := strconv.Atoi(reader.Text())
		fuel += i / 3 - 2
	}

	fmt.Println(fuel)
}
