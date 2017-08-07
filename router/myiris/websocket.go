package myiris

import (
	"fmt"
	"github.com/kataras/iris"
	"log"
)

type clientPage struct {
	Title string
	Host  string
}

func WebSocketTest() {
	iris.Static("/js", "./static/js", 1)

	iris.Get("/", func(ctx *iris.Context) {
		ctx.Render("index.html", clientPage{"Client Page", ctx.HostString()})
	})

	// the path which the websocket client should listen/registed to ->
	// websocket 客户端应该 监听或注册到的路径 ->
	iris.Config.Websocket.Endpoint = "/my_endpoint"
	// for Allow origin you can make use of the middleware
	// 你可以使用中间件来 Allow 源
	//myiris.Config.Websocket.Headers["Access-Control-Allow-Origin"] = "*"

	var myChatRoom = "room1"
	iris.Websocket.OnConnection(func(c iris.WebsocketConnection) {

		c.Join(myChatRoom)
		c.On("chat", func(message string) {
			//	// to all except this connection ->
			//	// 给除了此连接的其它连接发送消息 ->
			//	//c.To(myiris.Broadcast).Emit("chat", "Message from: "+c.ID()+"-> "+message)
			//
			//	// to the client ->
			//	// 发送给客户端 ->
			//	//c.Emit("chat", "Message from myself: "+message)
			//
			//	//send the message to the whole room,
			//	// 向整个房间发送消息
			//	//all connections are inside this room will receive this message
			//	// 所有房间内的连接都会接收到这个消息
			c.To(myChatRoom).Emit("chat", "From: "+c.ID()+": "+message)

			log.Println("RecvMSG:", message)
		})
		c.OnMessage(func(message []byte) {
			log.Println("RecvMSG:", string(message))

		})

		c.OnDisconnect(func() {
			fmt.Printf("\nConnection with ID: %s has been disconnected!", c.ID())
		})
	})
}
