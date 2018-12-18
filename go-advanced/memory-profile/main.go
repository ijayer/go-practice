/*
 * 说明：Golang | 语言机制之内存剖析
 * 作者：zhe
 * 时间：2018-12-14 2:56 PM
 * 更新：https://zhezh09.github.io//post/tech/code/golang/20181212-go语言机制-03-内存剖析/
 */

package main

import (
	"bytes"
	"fmt"
)

// data represents a table of input and expected output.
var data = []struct {
	input  []byte
	output []byte
}{
	{[]byte("abc"), []byte("abc")},
	{[]byte("elvis"), []byte("Elvis")},
	{[]byte("aElvis"), []byte("aElvis")},
	{[]byte("abcelvis"), []byte("abcElvis")},
	{[]byte("eelvis"), []byte("eElvis")},
	{[]byte("aelvis"), []byte("aElvis")},
	{[]byte("aabeeeelvis"), []byte("aabeeeElvis")},
	{[]byte("e l v i s"), []byte("e l v i s")},
	{[]byte("aa bb e l v i saa"), []byte("aa bb e l v i saa")},
	{[]byte(" elvi s"), []byte(" elvi s")},
	{[]byte("elvielvis"), []byte("elviElvis")},
	{[]byte("elvielvielviselvi1"), []byte("elvielviElviselvi1")},
	{[]byte("elvielviselvis"), []byte("elviElvisElvis")},
}

// assembleInputStream combines all the input into a
// single stream for processing.
func assembleInputStream() []byte {
	var in []byte
	for _, d := range data {
		in = append(in, d.input...)
	}
	return in
}

// assembleOutputStream combines all the output into a
// single stream for comparing.
func assembleOutputStream() []byte {
	var out []byte
	for _, d := range data {
		out = append(out, d.output...)
	}
	return out
}

func main() {
	var output bytes.Buffer
	in := assembleInputStream()
	out := assembleOutputStream()

	find := []byte("elvis")
	repl := []byte("Elvis")

	fmt.Println("=======================================\nRunning Algorithm One")
	output.Reset()
	algOne(in, find, repl, &output)
	matched := bytes.Compare(out, output.Bytes())
	fmt.Printf("Matched: %v\nInp: [%s]\nExp: [%s]\nGot: [%s]\n", matched == 0, in, out, output.Bytes())

	fmt.Println("=======================================\nRunning Algorithm Two")
	output.Reset()
	algTwo(in, find, repl, &output)
	matched = bytes.Compare(out, output.Bytes())
	fmt.Printf("Matched: %v\nInp: [%s]\nExp: [%s]\nGot: [%s]\n", matched == 0, in, out, output.Bytes())
}

// algOne is one way to solve the problem.
//
// 	    # Bench Test for Old Code
// 	    [zhe@zhe memory-profile](master)$ go test -run none -bench AlgorithmOne -benchtime 3s -benchmem  -memprofile mem.out
// 	    goos: windows
// 	    goarch: amd64
// 	    pkg: github.com/zhezh09/go-practice/go-advanced/memory-profile
// 	    BenchmarkAlgorithmOne-4          1000000              3136 ns/op             117 B/op          2 allocs/op
// 	    PASS
// 	    ok      github.com/zhezh09/go-practice/go-advanced/memory-profile       3.682s
//
//      # Bench Test for Old Code
//
func algOne(data []byte, find []byte, repl []byte, output *bytes.Buffer) {

	// Use a bytes Buffer to provide a stream to process.
	input := bytes.NewBuffer(data)

	// The number of bytes we are looking for.
	size := len(find)

	// Declare the buffers we need to process the stream.
	buf := make([]byte, size)
	end := size - 1

	// Read in an initial number of bytes we need to get started.
	//
	// ----------------Old Code(To be optimized)
	// if n, err := io.ReadFull(input, buf[:end]); err != nil {
	// 	output.Write(buf[:n])
	// 	return
	// }

	// ---------------- New Code(Optimized)
	if n, err := input.Read(buf[:end]); err != nil || n < end {
		output.Write(buf[:n])
		return
	}

	for {

		// Read in one byte from the input stream.
		//
		// ----------------Old Code(To be optimized)
		// if _, err := io.ReadFull(input, buf[end:]); err != nil {
		//
		// 	// Flush the reset of the bytes we have.
		// 	output.Write(buf[:end])
		// 	return
		// }

		// ---------------- New Code(Optimized)
		if _, err := input.Read(buf[end:]); err != nil {

			// Flush the reset of the bytes we have.
			output.Write(buf[:end])
			return
		}

		// If we have a match, replace the bytes.
		if bytes.Compare(buf, find) == 0 {
			output.Write(repl)

			// Read a new initial number of bytes.

			// ----------------Old Code(To be optimized)
			// if n, err := io.ReadFull(input, buf[:end]); err != nil {
			// 	output.Write(buf[:n])
			// 	return
			// }

			// ----------------New Code(Optimized)
			if n, err := input.Read(buf[:end]); err != nil {
				output.Write(buf[:n])
				return
			}

			continue
		}

		// Write the front byte since it has been compared.
		output.WriteByte(buf[0])

		// Slice that front byte out.
		copy(buf, buf[1:])
	}
}

// algTwo is a second way to solve the problem.
// Provided by Tyler Bunnell https://twitter.com/TylerJBunnell
func algTwo(data []byte, find []byte, repl []byte, output *bytes.Buffer) {

	// Use the bytes Reader to provide a stream to process.
	input := bytes.NewReader(data)

	// The number of bytes we are looking for.
	size := len(find)

	// Create an index variable to match bytes.
	idx := 0

	for {

		// Read a single byte from our input.
		b, err := input.ReadByte()
		if err != nil {
			break
		}

		// Does this byte match the byte at this offset?
		if b == find[idx] {

			// It matches so increment the index position.
			idx++

			// If every byte has been matched, write
			// out the replacement.
			if idx == size {
				output.Write(repl)
				idx = 0
			}

			continue
		}

		// Did we have any sort of match on any given byte?
		if idx != 0 {

			// Write what we've matched up to this point.
			output.Write(find[:idx])

			// Unread the unmatched byte so it can be processed again.
			input.UnreadByte()

			// Reset the offset to start matching from the beginning.
			idx = 0

			continue
		}

		// There was no previous match. Write byte and reset.
		output.WriteByte(b)
		idx = 0
	}
}
