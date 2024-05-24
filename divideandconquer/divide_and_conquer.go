package divideandconquer

import (
	"log"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func Sum(numbers []int) int {
	length := len(numbers)
	if length == 0 {
		return 0
	}
	if len(numbers) == 1 {
		return numbers[0]
	}
	return numbers[0] + Sum(numbers[1:])
}

func Length(numbers []int) int {
	log.Println(numbers)
	if len(numbers) == 0 {
		return 0
	}
	return 1 + Length(numbers[1:])
}

func Max(numbers []int) int {
	if len(numbers) == 2 {
		if numbers[0] > numbers[1] {
			return numbers[0]
		} else {
			return numbers[1]
		}
	}
	subMax := Max(numbers[1:])
	if numbers[0] > subMax {
		return numbers[0]
	} else {
		return subMax
	}
}

func BinarySearch(numbers []int, target int, low int, high int) int {
	log.Println(low, high)
	if (high - low) == 1 {
		if numbers[low] == target {
			return low
		} else if numbers[high] == target {
			return high
		} else {
			return -1
		}
	}
	var mid int = low + high/2
	if numbers[mid] == target {
		return mid
	} else if numbers[mid] > target {
		return BinarySearch(numbers, target, low, mid-1)
	} else {
		return BinarySearch(numbers, target, mid+1, high)
	}
}
