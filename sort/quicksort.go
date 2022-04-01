package sort

func partition[T int | string](arr []T, low, high int) int {
	index := low - 1
	pivotElement := arr[high]

	for i := low; i < high; i++ {
		if arr[i] < pivotElement {
			index += 1
			arr[index], arr[i] = arr[i], arr[index]
		}
	}

	arr[index+1], arr[high] = arr[high], arr[index+1]
	return index + 1
}

func quickSortRec[T int | string](arr []T, low, high int) {
	if len(arr) <= 1 {
		return
	}
	if low < high {
		pivot := partition(arr, low, high)
		quickSortRec(arr, low, pivot-1)
		quickSortRec(arr, pivot+1, high)
	}

}

func QuickSort[T int | string](arr []T) []T {
	quickSortRec(arr, 0, len(arr)-1)
	return arr
}
