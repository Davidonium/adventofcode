package util

import "math"

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := range integers {
		result = LCM(result, integers[i])
	}

	return result
}


func DigitCount(n int) int {
	return int(math.Floor(math.Log10(float64(n))+1))
}
