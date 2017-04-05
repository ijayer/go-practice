package myiris

import (
	"github.com/kataras/iris"
	"log"
)

type Test struct {
	*iris.Context
}

func testGet(c *iris.Context) {
	log.Println("@_@_____testGet")
}

func testPost(c *iris.Context) {
	log.Println("@_@_____testPost")
}

func testPut(c *iris.Context) {
	iris.Logger.Println("testPut")
}

type DealWithJson struct {
	*iris.Context
}

func (c DealWithJson) Post() {
	respData := dealWithJsonData(c)
	// response to data
	c.Write(string(iris.StatusOK), respData)
	c.JSON(iris.StatusOK, respData)
}

func dealWithJsonData(c DealWithJson) (respData interface{}) {
	// read request paras
	c.ReadJSON(&respData)
	c.ReadForm(&respData)
	return
}