package main

import (
	"fmt"
)

func BinarySearch(list []int, target int) (int, bool) {
	var low int = 0
	var high int = len(list) - 1
	for low <= high {
		// Go does floor division by default for integers
		var mid int = (low + high) / 2
		if list[mid] == target {
			return mid, true
		} else if list[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1, false
}

func main() {
	var test []int = []int{1, 3, 5, 7, 9}
	target := 80
	index, ok := BinarySearch(test, target)
	if ok {
		fmt.Printf("Value %v found at index %v\n", target, index)
	} else {
		fmt.Printf("Value %v could not be found in the array\n", target)
	}
}
