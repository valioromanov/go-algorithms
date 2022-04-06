package sort

func InsertionSort(arr []int) []int {
	i := 1

	for i < len(arr) {
		j := i
		for j >= 1 && arr[j] < arr[j-1] {
			arr[j], arr[j-1] = arr[j-1], arr[j]
			j--
		}
		i++
	}

	return arr
}
