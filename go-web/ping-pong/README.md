Gorilla WebSocket
=================

## 连接建立过程

### http->websocket

1. client: websocket dial: ws://ip:port/url
2. server: upgrade http -> websocket
    - handle http upgrade
    - call `h.Hijack()` got net.Conn
    - call `newConn()` got websocket.Conn
        - will SetPingHandle(fn) // if fn=nil, with default PingHandleFunc
        - will SetPongHandle(fn) // if fn=nil, with default PongHandleFunc
    - some setting and return websocket.Conn
3. server: start goroutine: read & write
    - go conn.Reader()
    - go conn.Writer()

##  Keep-Alive

### Ping Pong

Client: ping

> 客户端在与服务器建立连接后，按一定时间间隔发 ping 消息给服务器

```go
    c.Ws.SetWriteDeadline(time.Now().Add(WriteWait))
    err := c.Ws.WriteMessage(websocket.PingMessage, []byte{})
    if err != nil {
        fmt.Printf("Error: ping %v failed.%v\n", c.Ws.RemoteAddr(), err)
        return
    }
```

Server: pong

> 在服务端 Reader() 函数设置回调函数：SetPingHandler 对客户端发来的Ping消息回应Pong；如果回调函数为空，
则Gorilla Websocket会在底层初始化一个PingHandler匿名函数

```go
c.Ws.SetPingHandler(func(string) error {
    c.Ws.SetReadDeadline(time.Now().Add(PingWait))
    if err := c.Ws.WriteControl(websocket.PongMessage, []byte{}, time.Now().Add(WriteWait)); err != nil {
        return err
    }
    fmt.Printf("pong %v msg: %v\n", c.Ws.RemoteAddr(), websocket.PongMessage)
    return nil
})
```

> Note: 实际测试中发现，如果没有设置Ping-Pong相关的函数(或说C&S端都未发送Ping-Pong消息)，连接也不会中断

## TODO

- 中断重连
