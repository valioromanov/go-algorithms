package sort

import "fmt"

// []int{3, 2, 1, 7, 4, 10, 6, 8, 5, 9}
func SelectionSort(arr []int) []int {

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr)
		minEl, minInd := arr[i], i
		for y := i + 1; y < len(arr); y++ {
			if arr[y] < minEl {
				minEl, minInd = arr[y], y
			}
		}
		arr[i], arr[minInd] = arr[minInd], arr[i]
	}

	return arr
}
