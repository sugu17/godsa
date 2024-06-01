package quicksort

import (
	"fmt"
	"log"
	"slices"
	"testing"
)

func TestMedianOfThree(t *testing.T) {
	slice1 := []int{1, 3, 5}
	slice2 := []int{5, 1, 3}
	slice3 := []int{8, 4, 1}
	slice4 := []int{12, 13, 18}
	slcs := []([]int){
		slice1,
		slice2,
		slice3,
		slice4,
	}
	for _, slice := range slcs {
		slices.Sort(slice)
		sliceStart := 0
		sliceEnd := len(slice) - 1
		expectedIdx := sliceStart + (sliceEnd-sliceStart)/2
		actualIdx := MedianOfThree(slice, sliceStart, sliceEnd)
		if actualIdx != expectedIdx {
			t.Fatalf(`MedianOfThree(%v) = %v, Wanted %v`, slice, actualIdx, expectedIdx)
		}
	}
}

func TestDutchFlagSort(t *testing.T) {
	slice := []int{1, 8, 9, 4, 89, 8, 7, 45, 90}
	pivotIdx := MedianOfThree(slice, 0, len(slice)-1)
	pivotValue := slice[pivotIdx]
	_, _, pivotIdx = DutchFlagSort(slice, 0, len(slice)-1, pivotValue)
	for i := 0; i < len(slice); i++ {
		if i < pivotIdx && !(slice[i] <= slice[pivotIdx]) {
			t.Fatalf(`TestDutchFlagSort(%v), value %v is to the left of pivot and is not lesser than or equal to pivot value`, slice, slice[i])
		} else if i > pivotIdx && !(slice[i] >= slice[pivotIdx]) {
			log.Println(pivotIdx, slice[pivotIdx])
			t.Fatalf(`TestDutchFlagSort(%v), value %v is to the right of pivot and is not greater than or equal to pivot value`, slice, slice[i])
		}
	}
}

func TestSort(t *testing.T) {
	slice1 := []int{92, 24, 4, 23, 1, 8, 90}
	slice2 := []int{2, 52, 5, 2, 5, 2, 2, 1}
	slice3 := []int{8, 4, 1}
	slice4 := []int{12, 13, 18, 53, 21, 12}
	slcs := []([]int){
		slice1,
		slice2,
		slice3,
		slice4,
	}
	for _, slice := range slcs {
		original := slice
		quickSortedSlice := make([]int, len(original))
		libSortedSlice := make([]int, len(original))
		copy(quickSortedSlice, original)
		copy(libSortedSlice, original)
		slices.Sort(libSortedSlice)
		Sort(quickSortedSlice, 0, len(quickSortedSlice)-1)
		fmt.Println(quickSortedSlice)
		fmt.Println("-----------------")
		if slices.Compare(libSortedSlice, quickSortedSlice) != 0 {
			t.Fatalf(`Sort(%v) = %v, Wanted %v`, original, quickSortedSlice, libSortedSlice)
		}
	}
}
