/*
 * 说明：flag 绑定自定义数据类型
 * 作者：zhe
 * 时间：2018-09-12 11:13 PM
 * 更新：
 */

package main

import (
	"flag"
	"fmt"
)

// 自定义flag参数值接收类型
type addrs []string

// 实现 String() string 方法
func (a *addrs) String() string {
	return fmt.Sprintf("%v", *a)
}

// 实现 Set(string) error 方法
func (a *addrs) Set(value string) error {
	*a = append(*a, value)
	return nil
}

// DBAdds 声明接收命令行参数的变量
var DBAdds addrs

// 定义 flag 并完成解析
func init() {
	flag.Var(&DBAdds, "db_addr", "database cluster server address")
	flag.Parse()
}

func main() {
	fmt.Println(DBAdds)
}
