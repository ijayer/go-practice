package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"unicode"
	"unicode/utf8"
)

type Vertex struct {
	Lat, Long float64
}

type Addr struct {
	doMainName string
}

var m map[string]Vertex

func main() {
	//__________________________________map
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	v := make(map[string]Addr)
	v["google"] = Addr{"www.google.com"}

	fmt.Println("##___________", m["Bell Labs"])
	fmt.Println("##___________", v["google"])

	//__________________________________map的文法
	var mm = map[string]Vertex{
		"Bell Labs": {
			40.68433, -74.39967,
		},
		"Google": {
			37.42202, -122.08408,
		},
	}
	fmt.Println("##___________", mm)

	//__________________________________修改map
	mmm := make(map[string]int)
	mmm["Answer"] = 42
	fmt.Println("##___________The value: ", mmm["Answer"])
	mmm["Answer"] = 48
	fmt.Println("##___________The value: ", mmm["Answer"])
	delete(mmm, "Answer")
	fmt.Println("##___________The value: ", mmm["Answer"])
	vv, ok := mmm["Answer"]
	fmt.Println("##___________The value: ", vv, "Present?", ok)

	//__________________________________函数值
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println("##___________hypot(5, 12) =", hypot(5, 12))

	fmt.Println("##___________compute(hypot) =", compute(hypot))
	fmt.Println("##___________compute(math.Pow) =", compute(math.Pow))

	//__________________________________统计utf-8字符
	counts := make(map[rune]int)
	var utflen [utf8.MaxRune + 1]int
	invalid := 0

	input := bufio.NewReader(os.Stdin)
	for {
		rune, n, err := input.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("##___________fatal error:", err.Error())
			os.Exit(1)
		}
		if rune == unicode.ReplacementChar {
			invalid++
			continue
		}
		counts[rune]++
		utflen[n]++
	}
	fmt.Println("##___________\rune\tcount")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}

	fmt.Println("##___________\nlen\tcount")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("##___________%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("##__________\n%d invalid utf-8 characters\n", invalid)
	}
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}
