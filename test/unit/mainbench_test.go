package unit

import "testing"

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(1, 2)
	}
}

func BenchmarkSub(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sub(1, 2)
	}
}

func BenchmarkMul(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Mul(1, 2)
	}
}

func BenchmarkDiv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Div(4, 2)
	}
}
