package sort

func SelectionSort(arr []int) []int {

	for i := 0; i < len(arr); i++ {
		minEl, minInd := arr[i], i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < minEl {
				minEl, minInd = arr[j], j
			}
		}
		arr[i], arr[minInd] = arr[minInd], arr[i]
	}

	return arr
}
