/*
 * 说明：
 * 作者：zhe
 * 时间：2018-12-14 3:32 PM
 * 更新：
 */

package main

import (
	"bytes"
	"testing"
)

func BenchmarkAlgorithmOne(b *testing.B) {
	var output bytes.Buffer
	in := assembleInputStream()
	find := []byte("elvis")
	repl := []byte("Elvis")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		output.Reset()
		algOne(in, find, repl, &output)
	}
}
