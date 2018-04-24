/*
 * 说明：
 * 作者：zhe
 * 时间：2018-04-23 14:51
 * 更新：
 */

package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Name string `json:"name"`
}

func main() {
	var u user

	data := `{"name":"you"}`

	// json unmarshal 必须传入一个地址参数用来存储最后的结果
	//
	// 底层函数：
	// rv := reflect.ValueOf(v)
	// if rv.Kind() != reflect.Ptr || rv.IsNil() {
	// 	return &InvalidUnmarshalError{reflect.TypeOf(v)}
	// }
	err := json.Unmarshal([]byte(data), &u)
	if err != nil {
		fmt.Println("Error:", err.Error())
	}
	fmt.Println(u)

	// 当参数为指针时，Unmarshal的参数即为该指针地地址
	var up *user
	if up == nil {
		fmt.Println("up", &up)
	}
	err = json.Unmarshal([]byte(data), &up)
	if err != nil {
		fmt.Println("Error:", err.Error())
	}
	fmt.Println(up)

	// json marshal 传入值类型和指针类型都🆗
	u = user{Name: "biu°"}
	bytes, err := json.Marshal(u)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(string(bytes))

	bytes, err = json.Marshal(&u)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(string(bytes))
}
