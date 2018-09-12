/*
 * 说明：切片引用
 * 作者：zhe
 * 时间：2018-04-19 22:32
 * 更新：
 */

package main

func main() {
	println(basename("a/b/c.go"))
	// output: go
	//
	// analysis: basename
	//      i=1 str: (str[2:]) str = b/c/.go
	//      i=3 str: (str[4:]) str = go
}

func basename(str string) string {
	for i, v := range str {
		println("len", len(str))
		println(i, string(v))
		if v == '/' {
			str = str[i+1:]
		}
	}
	return str
}
