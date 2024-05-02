package main

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
}
