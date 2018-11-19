/*
 * 说明：『Go核心36讲』| 02 & 03 | *源码文件 Demo
 * 作者：zhe
 * 时间：2018-09-13 8:59 PM
 * 更新：
 */

package main

// 导入 lib:
// 可以看到这里的导入路径为："go36/01/q1/lib"
// 但我们在调用 lib 下的程序实体时，却使用的限定符为 `say`

func main() {
	say.Hi("joe")
}
