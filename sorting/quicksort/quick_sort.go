package quicksort

import (
	"cmp"
	"log"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func MedianOfThree[T cmp.Ordered](arr []T) int {
	start := 0
	end := len(arr) - 1
	mid := start + ((end - start) / 2)
	if arr[start] <= arr[mid] {
		if arr[start] >= arr[end] {
			return start
		} else if arr[mid] <= arr[end] {
			return mid
		} else {
			return end
		}
	} else {
		if arr[start] <= arr[end] {
			return start
		} else if arr[mid] >= arr[end] {
			return mid
		} else {
			return end
		}
	}
}
