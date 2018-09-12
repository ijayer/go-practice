//从控制台去读输入
//fmt包中Scan和Scanf开头函数的使用
package main

import (
	"fmt"
)

var (
	firstName, lastName, s string
	www, google, com       string
	i                      int
	ff                     float32
	str                    = "24 / 36.2 / Go"
	format                 = "%d / %f / %s"
)

func main() {
	fmt.Scan(&firstName, &lastName, &s) //www baidu com
	fmt.Println(firstName, lastName, s) //www baidu com

	fmt.Scanf("%s %s %s %d", &www, &google, &com, &i) //www google com
	fmt.Println(www, google, com, i)                  //www google com

	//input from str to i, f, s
	fmt.Sscanf(str, format, &i, &ff, &s)
	fmt.Println("Read from str: ", i, ff, s)
}
