package main

import (
	"fmt"
	//"math/rand"
)

func main() {
	var p1ships [10][10]string
	var p2ships [10][10]string
	var p1shot [10][10]string
	var p2shot [10][10]string
	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			p1ships[x][y] = "."
			p2ships[x][y] = "."
			p1shot[x][y] = "."
			p2shot[x][y] = "."
		}
	}
	var carrier1 [5]int
	var battleship1 [4]int
	var cruiser1 [3]int
	var submarine1 [3]int
	var destroyer1 [2]int
	var carrier2 [5]int
	var battleship2 [4]int
	var cruiser2 [3]int
	var submarine2 [3]int
	var destroyer2 [2]int
	// randInt
	var p1shipcoords [2]int
}
