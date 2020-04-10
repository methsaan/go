package main

import (
	"fmt"
)

func bubblesort(list [100]int, start int, end int) []int {
	var s int = start
	var e int = end
	var temp int
	for x := s; x < end; x++ {
		for y := s; y < e-1; y++ {
			if list[y] > list[y+1] {
				temp = list[y]
				list[y] = list[y+1]
				list[y+1] = temp
			}
		}
		e--
	}
	return list[start:end]
}

func main() {
	var list [100]int
	var listitem int
	var exit string
	var s int
	var e int
	for x := 0; exit != "q"; x++ {
		fmt.Print("Enter next item: ")
		fmt.Scanf("%d %s", &listitem, &exit)
		list[x] = listitem
	}
	fmt.Print("Enter range to sort: ")
	fmt.Scanf("%d, %d", &s, &e)
	fmt.Println("Original list: \t\t ", list[s:e])
	fmt.Println("Sorted list: \t\t ", bubblesort(list, s, e))
}
