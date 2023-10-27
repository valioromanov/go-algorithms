package main

import (
	"alhorithms/tree"
	"os"
)

func main() {
	//ints := []int{3, 2, 1, 7, 4, 10, 6, 8, 5, 9}

	// fmt.Println("Sorting with BubbleSort....")
	// ints = sort.BubbleSort(ints)
	// fmt.Println(ints)

	/*fmt.Println("Sorting with CombSort....")
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

	// fmt.Println("Sorting with Selection sort second variant")
	// ints = sort.SelectionSort(ints)
	// fmt.Println(ints)

	// fmt.Println(search.BinarySearch(ints, 11))

	treeOne := &tree.BinaryTree{}
	treeOne.InsertEl(100).
		InsertEl(-20).
		InsertEl(-50).
		InsertEl(-15).
		InsertEl(-60).
		InsertEl(50).
		InsertEl(60).
		InsertEl(55).
		InsertEl(85).
		InsertEl(15).
		InsertEl(5).
		InsertEl(-10).
		InsertEl(150).
		InsertEl(145).
		InsertEl(160).
		InsertEl(165)
	tree.PrintTree(os.Stdout, treeOne.Root, 0, 'M')
}
