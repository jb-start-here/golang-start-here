//go:build ignore 

package main

import "fmt"

// SumInts adds together the values of m.
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
			s += v
	}
	return s
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
			s += v
	}
	return s
}

// The above code is repetitive and can be replaced with a generic function.
// The generic function can be defined as follows:
// T is a constraint that can be int64 or float64.
// You can also have a generic without the constraint
func Sum[T int64 | float64](m map[string]T) T {
	var s T
	for _, v := range m {
		s += v
	}
	return s
}


func main() {
	x := map[string]int64{"a": 1, "b": 2, "c": 3}
	y := map[string]float64{"a": 1.1, "b": 2.2, "c": 3.3}

	fmt.Println(SumInts(x))   // 6
	fmt.Println(SumFloats(y)) // 6.6

	fmt.Println(Sum(x)) // 6
	fmt.Println(Sum(y)) // 6.6
}
