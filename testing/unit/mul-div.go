package unit

import "errors"

func Mul(a, b int) int {
	return a * b
}

func Div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("divisor can not be 0")
	}
	return a / b, nil
}
