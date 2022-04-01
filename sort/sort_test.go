package sort

import (
	"testing"
)

func TestBubleSort(t *testing.T) {
	arrForSorting := []int{5, 7, 2, 4, 1, 3, 9, 8, 6}
	expectingOutput := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	arrForSorting = Bubble(arrForSorting)
	for i, el := range arrForSorting {
		if el != expectingOutput[i] {
			t.Errorf("BubbleSort: The sorted array and expectiong array are different: sorted (%d), excpected (%d)", arrForSorting, expectingOutput)
		}
	}
}

func TestCombSort(t *testing.T) {
	arrForSorting := []int{5, 7, 2, 4, 1, 3, 9, 8, 6}
	expectingOutput := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	arrForSorting = CombSort(arrForSorting)
	for i, el := range arrForSorting {
		if el != expectingOutput[i] {
			t.Errorf("CombSort: The sorted array and expectiong array are different: sorted (%d), excpected (%d)", arrForSorting, expectingOutput)
		}
	}
}
