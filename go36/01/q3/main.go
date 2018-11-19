/*
 * 说明：『Go核心36讲』| 02 & 03 | *源码文件 Demo
 * 作者：zhe
 * 时间：2018-09-13 10:00 PM
 * 更新：
 */

package main

/*
// Case: 代码包冲突
//
// 描述：下面代码中，import 后路径最后一级相同，且产生了冲突的解决方法测试，分 4 中：
//
import (
	"go36/01/q3/lib1/say"
	"go36/01/q3/lib2/say"
)

func main() {
	say.SayHi()
	say.SayHello()
}
*/

// --------------------------------------------Resolve::A -> 设置别名
// import (
// 	"go36/01/q3/lib1/say"
// 	say2 "go36/01/q3/lib2/say"
// )
//
// func main() {
// 	say.SayHi()
// 	say2.SayHello()
// }

// --------------------------------------------Resolve::B -> 导入的点操作
// import (
// 	. "go36/01/q3/lib1/say"
// 	"go36/01/q3/lib2/say"
// )
//
// func main() {
// 	SayHi()
// 	say.SayHello()
// }

// --------------------------------------------Resolve::C -> _ 只导入包，无其内部程序实体的引用
// import (
// 	_ "go36/01/q3/lib1/say"
// 	"go36/01/q3/lib2/say"
// )
//
// func main() {
// 	say.SayHello()
// }

// --------------------------------------------Resolve::D -> 修改包声明语句：即重命名包名
import (
	"go36/01/q3/lib1/say"
	"go36/01/q3/lib2/say" // 包名声明为：package sayhello
)

func main() {
	say.SayHi()
	sayhello.SayHello()
}
