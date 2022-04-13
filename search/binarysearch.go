package search

/*
	This method needs an ordered slice of elements
	I.E. Search for 7
	M = (0 + 9) / 2  = 4
	[1 2 3 4 5 6 7 8 9 10]
	 L       M          R
	arr[M] != 7 arr[M] < 7 =>
	M = 7
	[1 2 3 4 5 6 7 8 9 10]
			   L   M    R
	arr[M] != 7 && arr[M] > 7=>
	M = 6
	[1 2 3 4 5 6 7 8 9 10]
			    LM
				 R
	arr[M] = 7
*/

func BinarySearch(arr []int, target int) int {
	leftPoint := 0
	rightPoint := len(arr) - 1

	for leftPoint <= rightPoint {
		midPoint := int((leftPoint + rightPoint) / 2)

		if arr[midPoint] == target {
			return midPoint
		} else if arr[midPoint] > target {
			rightPoint = midPoint - 1
		} else if arr[midPoint] < target {
			leftPoint = midPoint + 1
		}

	}
	return -1
}
