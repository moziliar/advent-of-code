package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

type edge struct {
	start point
	end   point
}

func vertical(e edge) bool {
	return e.start.x == e.end.x
}

func parallel(e1 edge, e2 edge) bool {
	return vertical(e1) == vertical(e2)
}

func (e edge) intersect(e2 edge) (bool, point) {
	if parallel(e, e2) {
		return false, point{}
	}
	if vertical(e) {
		b := within(e.start.x, e2.start.x, e2.end.x) && within(e2.start.y, e.start.y, e.end.y)
		return b, point{e.start.x, e2.start.y}
	} else {
		b := within(e2.start.x, e.start.x, e.end.x) && within(e.start.y, e2.start.y, e2.end.y)
		return b, point{e2.start.x, e.start.y}
	}
}

func within(guess int, x1 int, x2 int) bool {
	return guess >= x1 && guess <= x2 || guess >= x2 && guess <= x1
}

func main() {
	sc := bufio.NewReader(os.Stdin)
	line1, _ := sc.ReadString('\n')
	line2, _ := sc.ReadString('\n')
	wire1 := readWire(line1)
	wire2 := readWire(line2)
	fmt.Println(nearestInt(wire1, wire2))
}

func readWire(input string) []edge {
	in := strings.Split(input, ",")
	prev := point{x: 0, y: 0}
	points := []point{prev}
	for _, v := range in {
		op := v[0]
		length, _ := strconv.Atoi(strings.Trim(v, string(op)))
		var pt point
		switch op {
		case 'U':
			pt = point{x: prev.x, y: prev.y + length}
		case 'D':
			pt = point{x: prev.x, y: prev.y - length}
		case 'L':
			pt = point{x: prev.x - length, y: prev.y}
		case 'R':
			pt = point{x: prev.x + length, y: prev.y}
		}
		points = append(points, pt)
		prev = pt
	}
	var wire []edge
	start := point{x: 0, y: 0}
	for i, v := range points {
		if i == 0 {
			continue
		}
		e := edge{start, v}
		wire = append(wire, e)
		start = v
	}
	return wire
}

func nearestInt(w1 []edge, w2 []edge) int {
	min := math.MaxFloat64
	for _, e1 := range w1 {
		for _, e2 := range w2 {
			if inter, p := e1.intersect(e2); inter{
				if p.x == 0 && p.y == 0 {
					continue
				}
				min = math.Min(min, math.Abs(float64(p.x))+math.Abs(float64(p.y)))
			}
		}
	}
	return int(min)
}
