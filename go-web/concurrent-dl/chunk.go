/*
 * 说明：
 * 作者：zhe
 * 时间：2018-11-19 2:43 PM
 * 更新：
 */

package main

const chunk = 4 * 1024 * 1024

func main() {
	fileSize := int64(40960000)
	println(fileSize%chunk == 0)
	println()

	shardByChunk(fileSize)
	println()
	shardByGoroutines(fileSize)
}

func shardByChunk(fileSize int64) {
	n := fileSize / chunk
	println(n)

	if fileSize%chunk != 0 {
		n += 1
	}
	println(n)

	var startPos int64
	var endPos int64
	for i := 0; i < int(n); i++ {
		startPos = int64(i * chunk)
		endPos = int64(i+1)*chunk - 1
		if i == (int(n) - 1) {
			endPos = fileSize
		}
		println("#", i, "  ", startPos, "-", endPos)
	}
}

var goN = 5

func shardByGoroutines(fileSize int64) {
	var rangeS int64
	var remain int64

	remain = fileSize % int64(goN)
	println(remain)

	rangeS = fileSize / int64(goN)
	println(rangeS)

	var startPos int64
	var endPos int64
	for i := 0; i < goN; i++ {
		startPos = int64(i) * rangeS
		endPos = int64(i+1)*rangeS - 1

		if i == (goN-1) && remain != 0 {
			endPos += remain
		}
		println("#", i, "  ", startPos, "-", endPos)
	}
}
