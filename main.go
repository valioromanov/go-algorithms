package main

import (
	"algorithms/sort"
	"fmt"
)

func main() {
	ints := []int{3, 2, 1}
	stringss := []string{"c", "b", "a"}

	fmt.Println("Sorting with BubbleSort....")
	ints = sort.Bubble(ints)
	stringss = sort.Bubble(stringss)
	fmt.Println(ints)
	fmt.Println(stringss)

	fmt.Println("Sorting with CombSort....")
	ints = sort.CombSort(ints)
	stringss = sort.CombSort(stringss)
	fmt.Println(ints)
	fmt.Println(stringss)
}
