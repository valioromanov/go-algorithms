package main

import (
	"alhorithms/sort"
	"fmt"
)

func main() {
	ints := []int{3, 2, 1, 7, 4, 10, 6, 5}

	/*fmt.Println("Sorting with BubbleSort....")
	ints = sort.Bubble(ints)
	stringss = sort.Bubble(stringss)
	fmt.Println(ints)
	fmt.Println(stringss)

	fmt.Println("Sorting with CombSort....")
	ints = sort.CombSort(ints)
	stringss = sort.CombSort(stringss)
	fmt.Println(ints)
	fmt.Println(stringss)*/

	/*fmt.Println("Sorting with QuickSort....")
	ints = sort.QuickSort(ints)
	fmt.Println(ints)*/

	/*fmt.Println("Sorting with insertion sort")
	ints = sort.InsertionSort(ints)
	fmt.Println(ints)*/

	/*fmt.Println("Sorting with Bubble sort second variant")
	ints = sort.BubbleSort(ints)
	fmt.Println(ints)*/

	fmt.Println("Sorting with Selection sort second variant")
	ints = sort.SelectionSort(ints)
	fmt.Println(ints)
}
