/*
 * 说明：记一次栈溢出错误：`fatal error: stack overflow`
 * 作者：zhe
 * 时间：2018-04-23 15:11
 * 更新：
 */

package main

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
	"gopkg.in/mgo.v2/bson"
)

var data = `{
    "_id": "5ad4030fc7f41c2920eeccd4",
    "account": "mongo_0",
    "age": 0,
    "comments": [],
    "create_at": "2018-04-16 09:57:35",
    "delete_at": "",
    "email": "3061871118@qq.com"
}`

func main() {
	var resp Response

	resp = dataWithBsonMSlice()
	bytes, err := resp.MarshalJSON()
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(string(bytes))

	resp = dataWithBsonM()
	resp.MarshalJSON()
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(string(bytes))
}

// Response 数据库查询结果
type Response struct {
	Total int         `json:"total"`
	Data  interface{} `json:"data"`
}

type Responder interface {
	result() interface{}
}

func (r Response) result() interface{} {
	if r.Total == 0 || r.Data == nil {
		r.Data = []string{}
	}
	return r
}

// dataWithBsonSlice
func dataWithBsonMSlice() Response {
	var m bson.M
	var ms []bson.M
	err := json.Unmarshal([]byte(data), &m)
	if err != nil {
		fmt.Println("Error dataWithBsonMSlice:", err)
		return Response{}
	}
	ms = append(ms, m)

	return Response{
		Total: len(ms),
		Data:  ms,
	}
}

// dataWithBsonM
func dataWithBsonM() Response {
	var m bson.M
	err := json.Unmarshal([]byte(data), &m)
	if err != nil {
		fmt.Println("Error dataWithBsonM:", err)
		return Response{}
	}

	return Response{
		Total: 1,
		Data:  m,
	}
}

// 在序列化或反序列化的时候如何改变一个或多个字段的值: 自己实现 Marshaler | Unmarshaler 接口的方法
//
// Response.MarshalJSON 实现 Marshaler
//
// Marshaler is the interface implemented by types that
// can marshal themselves into valid JSON.
//
// func (r Response) MarshalJSON() ([]byte, error) {
// 	resp := r.result().(Response)
// 	fmt.Printf("type: %T, data:%+v\n", resp.Data, resp.Data)
//
// 	return json.Marshal(resp)
// 	// Fatal: endless loop (json.Marshal的参数不能在传入Response类型的对象)
// 	//              |
// 	//        runtime: goroutine stack exceeds 1000000000-byte limit
// 	//        fatal error: stack overflow
// 	//              |
// 	// call again in: D:/Go/src/encoding/json/encode.go:445
// }

// 优化：MarshalJSON，构造新的匿名结构体，并用Old Data进行填充，或重命名json标签
// 该方式的缺点：不适用于拥有大量字段的 custom types
// func (r Response) MarshalJSON() ([]byte, error) {
// 	resp := r.result().(Response)
// 	// fmt.Printf("type: %T, data:%+v\n", resp.Data, resp.Data)
//
// 	return json.Marshal(struct {
// 		Total int         `json:"total"`
// 		Data  interface{} `json:"data"`
// 	}{
// 		Total: resp.Total,
// 		Data:  resp.Data,
// 	})
// }

// 死循环：当新构造的struct嵌入原始对象时，也会继承其的所有字段以及方法，就会导致程序进入 `infinite loop`
// func (r Response) MarshalJSON() ([]byte, error) {
// 	resp := r.result().(Response)
//
// 	return json.Marshal(&struct {
// 		Data interface{}
// 		Response
// 	}{
// 		Data:     resp.Data,
// 		Response: r,
// 	})
// }

// 优化：alias the original type. This alias will have all the same fields, but none of the methods.
func (r Response) MarshalJSON() ([]byte, error) {
	resp := r.result().(Response)

	type AliasResp Response
	return json.Marshal(&struct {
		*AliasResp
		Data interface{} `json:"data"` // 序列化时，覆盖原始字段data的值
	}{
		AliasResp: (*AliasResp)(&resp),
		Data:      resp.Data,
	})
	// output:
	// {"total":1,"data":[{"_id":"5ad4030fc7f41c2920eeccd4","account":"mongo_0","age":0,"comments":[],"create_at":
	// 	"2018-04-16 09:57:35","delete_at":"","email":"3061871118@qq.com"}]}
	// {"total":1,"data":[{"_id":"5ad4030fc7f41c2920eeccd4","account":"mongo_0","age":0,"comments":[],"create_at":
	// 	"2018-04-16 09:57:35","delete_at":"","email":"3061871118@qq.com"}]}
}

// 实现 Unmarshaler 接口
func (r Response) UnmarshalJSON([]byte) error {
	return errors.New("nil")
}
