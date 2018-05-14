/*
 * 说明：
 * 作者：zhe
 * 时间：2018-05-14 15:07
 * 更新：
 */

package unit

import "testing"

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
