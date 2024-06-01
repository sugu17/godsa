package quicksort

import (
	"slices"
	"testing"
	"log"
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
		log.Println(slice)
		slices.Sort(slice)
		sliceStart := 0
		sliceEnd := len(slice) - 1
		expectedIdx := sliceStart + (sliceEnd-sliceStart)/2
		actualIdx := MedianOfThree(slice)
		log.Printf("Actual %v, Expected %v", slice[actualIdx], slice[expectedIdx])
		if actualIdx != expectedIdx {
			t.Fatalf(`MedianOfThree(%v) = %v, Wanted %v`, slice, actualIdx, expectedIdx)
		}
	}
}
