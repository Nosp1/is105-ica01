package sum

import(
	"fmt"
	"math"
)

func SumInt8(a, b int) int8 {
	var s int8
	if a+b >= math.MinInt8 && a+b <= math.MaxInt8 {
		s = int8(a + b)
		return s
	} else {
		fmt.Print("Sum of parameters a: ", a, "and b: ", b, "out of range: ", a+b)
		return s
	}
}

func SumUint32(a, b uint32) uint32 {
	return a + b
}

func SumInt32(a, b int32) int32 {
	return a + b
}

func SumInt64(a, b int64) int64 {
	return a + b
}

func SumFloat64(a, b float64) float64 {
	return a + b
}


