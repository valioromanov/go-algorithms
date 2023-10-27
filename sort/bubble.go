package sort

import "fmt"

// this function sort the slice with a simple Bubble sort algorithm
func Bubble(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

// ints := []int{3, 2, 1, 7, 4, 10, 6, 8, 5, 9}
// ints := []int{2, 3, 1, 7, 4, 10, 6, 8, 5, 9}
// ints := []int{2, 1, 3, 7, 4, 10, 6, 8, 5, 9}
func BubbleSort(arr []int) []int {
	swapped := true

	for swapped {
		swapped = false
		fmt.Println("...")
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				swapped = true
				arr[i], arr[i+1] = arr[i+1], arr[i]
				fmt.Println(arr)
				fmt.Println(swapped)
			}
		}
	}

	return arr
}
