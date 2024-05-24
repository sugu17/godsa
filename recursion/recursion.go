package recursion

import (
	"log"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func Factorial(number int) int {
	if number == 1 {
		return 1
	}
	return number * Factorial(number-1)
}
