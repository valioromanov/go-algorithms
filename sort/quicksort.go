package sort

import "fmt"

func partition(arr []int, low, high int) int {
	fmt.Printf("arr: [%v], low: {%d}, high: {%d} \n", arr, low, high)
	index := low - 1
	pivotElement := arr[high]

	for i := low; i < high; i++ {
		fmt.Printf("arr[%d]: %v - pivodElement: %v \n", i, arr[i], pivotElement)
		if arr[i] < pivotElement {
			fmt.Printf("arr[%d] is smaller than pivotelement \n", i)
			index += 1
			fmt.Printf("index: %d \n", index)
			fmt.Printf("arr[index]: arr[%d] = %v, arr[i]: arr[%d]=%v \n", index, arr[index], i, arr[i])
			arr[index], arr[i] = arr[i], arr[index]
		}
	}
	fmt.Printf("arr[index+1]: arr[%d] = %v, arr[high]: arr[%d]=%v \n", index+1, arr[index+1], high, arr[high])
	arr[index+1], arr[high] = arr[high], arr[index+1]
	return index + 1
}

func quickSortRec(arr []int, low, high int) {
	if len(arr) <= 1 {
		return
	}
	if low < high {
		pivot := partition(arr, low, high)
		quickSortRec(arr, low, pivot-1)
		quickSortRec(arr, pivot+1, high)
	}

}

func QuickSort(arr []int) []int {
	quickSortRec(arr, 0, len(arr)-1)
	return arr
}
