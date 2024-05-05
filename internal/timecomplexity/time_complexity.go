package timecomplexity

import (
	"fmt"
	"math"
	"os"
	"text/tabwriter"
)

type ComplexityType int

const (
	_                       = iota
	BigOLogn ComplexityType = iota
	BigOn
	BigOnLogn
	BigOnSquare
)

func (c ComplexityType) String() string {
	switch c {
	case BigOLogn:
		return "O(log n)"
	case BigOn:
		return "O(n)"
	case BigOnLogn:
		return "O(n * log n)"
	case BigOnSquare:
		return "O(n2)"
	default:
		return "Unknown complexity type."
	}
}

func CalculateTimeComplexity(inputSize int, complexity ComplexityType) int {
	switch complexity {
	case BigOLogn:
		return int(math.Floor(math.Log2(float64(inputSize))))
	case BigOn:
		return inputSize
	case BigOnLogn:
		return int(math.Floor(float64(inputSize) * (math.Log2(float64(inputSize)))))
	case BigOnSquare:
		return inputSize * inputSize
	}
	return -1
}

func CompareTimeComplexity(sizes ...int) {
	if len(sizes) == 0 {
		fmt.Println("No input size provided. Exiting...")
		return
	}
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', 0)
	fmt.Fprintf(w, "%v\t%v\t%v\t%v\n", BigOLogn, BigOn, BigOnLogn, BigOnSquare)
	for _, n := range sizes {
		fmt.Fprintf(w, "%v\t%v\t%v\t%v\n", CalculateTimeComplexity(n, BigOLogn), CalculateTimeComplexity(n, BigOn), CalculateTimeComplexity(n, BigOnLogn), CalculateTimeComplexity(n, BigOnSquare))
	}
	w.Flush()
}
