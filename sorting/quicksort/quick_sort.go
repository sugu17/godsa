package quicksort

import (
	"cmp"
	"fmt"
	"log"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func Digits(number int) int {
	count := 0
	for ; number > 0; number = number / 10 {
		count++
	}
	return count
}

func MedianOfThree[T cmp.Ordered](arr []T, start int, end int) int {
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

func PrintPointerPos(arr []int, start int, curr int, end int) {
	fmt.Printf("[")
	for i := range arr {
		digits := Digits(arr[i])
		for _ = range digits - 1 {
			fmt.Printf(" ")
		}
		if i == start {
			fmt.Printf("s")
		} else if i == curr {
			fmt.Printf("c")
		} else if i == end {
			fmt.Printf("e")
		} else {
			fmt.Printf(".")
		}
		if i != len(arr)-1 {
			fmt.Printf(" ")
		}
	}
	fmt.Printf("]\n")
	fmt.Printf("----------------------\n\n")
}

func DutchFlagSort[T cmp.Ordered](arr []T, start int, end int, pivotVal T) (int, int, int) {
	pivotIdx := -1
	for curr := start; curr <= end; {
		if arr[curr] < pivotVal {
			if arr[curr] != arr[start] {
				arr[curr], arr[start] = arr[start], arr[curr]
			}
			if arr[curr] == pivotVal {
				pivotIdx = curr
			}
			start++
			curr++
		} else if arr[curr] > pivotVal {
			if arr[curr] != arr[end] {
				arr[curr], arr[end] = arr[end], arr[curr]
			}
			if arr[curr] == pivotVal {
				pivotIdx = curr
			}
			end--
		} else {
			pivotIdx = curr
			curr++
		}
	}
	return start, end, pivotIdx
}

func Sort[T cmp.Ordered](arr []T, start int, end int) {
	if end-start <= 0 {
		return
	}
	medianIdx := MedianOfThree(arr, start, end)
	partStart, partEnd, _ := DutchFlagSort(arr, start, end, arr[medianIdx])
	fmt.Println(arr[partStart: partEnd+1])
	Sort(arr, start, partStart-1)
	Sort(arr, partEnd+1, end)
}
