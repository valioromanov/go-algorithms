package sort

func goToNextGap(gap int) int {
	gap = (gap * 10) / 13
	if gap < 1 {
		gap = 1
	}

	return gap
}

func CombSort[K int | string](arr []K) []K {

	arrLen := len(arr)
	gap := arrLen
	swapped := true
	for gap != 1 || swapped {
		swapped = false
		gap = goToNextGap(gap)
		for i := 0; i < arrLen-gap; i++ {
			if arr[i] > arr[i+gap] {
				arr[i], arr[i+gap] = arr[i+gap], arr[i]
				swapped = true
			}
		}
	}
	return arr
}
