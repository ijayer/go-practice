/*
 * 说明：Go 语言机制之逃逸分析
 * 作者：zhe
 * 时间：2018-12-12 4:18 PM
 * 更新：https://zhezh09.github.io//post/tech/code/golang/20181212-go语言机制-02-逃逸分析/
 */

package main

type user struct {
	name  string
	email string
}

func main() {
	u1 := createUserV1()
	u2 := createUserV2()
	u3 := createUserV3()

	println("u1", &u1, " u2", &u2, " u3", &u3)
}

//go:noinline
func createUserV1() user {
	u := user{
		name:  "Bill",
		email: "bill@gmail.com",
	}
	println("V1", &u)
	return u
}

//go:noinline
func createUserV2() *user {
	u := user{
		name:  "Bill",
		email: "bill@gmail.com",
	}
	println("V2", &u)
	return &u
}

/*
	# 逃逸分析
	$ go tool compile -m demo1.go
	demo1.go:28:16: createUserV1 &u does not escape
	demo1.go:39:9: &u escapes to heap
	demo1.go:34:2: moved to heap: u
	demo1.go:38:16: createUserVw &u does not escape
	demo1.go:19:16: main &u1 does not escape
	demo1.go:19:28: main &u2 does not escape
*/

// Updated: func createUserV3
//go:noinline
func createUserV3() *user {
	u := &user{
		name:  "Bill",
		email: "bill@gmail.com",
	}
	println("V3", u)
	return u
}
