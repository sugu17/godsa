package selection

import (
	"errors"
	"log"
)

type SortOrder int

const (
	Ascending SortOrder = iota
	Descending
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func pop(slice []int, index int) ([]int, bool) {
	if index < 0 || index > len(slice)-1 {
		return nil, false
	}
	return append(slice[:index], slice[index+1:]...), true
}

func findLowest(slice []int) int {
	index := 0
	minVal := slice[0]
	for idx, val := range slice {
		if val < minVal {
			minVal = val
			index = idx
		}
	}
	return index
}

func findHighest(slice []int) int {
	index := 0
	maxVal := slice[0]
	for idx, val := range slice {
		if val > maxVal {
			maxVal = val
			index = idx
		}
	}
	return index
}

func Sort(slice []int, sortOrder SortOrder) ([]int, error) {
	temp := append([]int{}, slice...)
	output := []int{}
	var test func([]int) int
	if sortOrder == Ascending {
		test = findLowest
	} else if sortOrder == Descending {
		test = findHighest
	} else {
		return nil, errors.New("Invalid sort order flag.")
	}
	for len(temp) > 0 {
		index := test(temp)
		output = append(output, temp[index])
		temp, _ = pop(temp, index)
	}
	return output, nil
}
