package unit

import "errors"

func main() {
}

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func Mul(a, b int) int {
	return a * b
}

func Div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("divisor can not be 0")
	}
	return a / b, nil
}
