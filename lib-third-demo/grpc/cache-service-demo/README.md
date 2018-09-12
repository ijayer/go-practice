gRPC Practice Demo
==================
```
AuthorBy: zher
CreateAt: 2018-09-11
ModifyAt: 2018-09-12
```

## Overview

- [原文][#1]
- [译文][#2]

### server 实现 gRPC 相关服务：

- 完成 k-v 的存放和管理
- 支持服务限流
- 并发访问控制
- 带有日志统计中间件
- 以流的方式发送数据

### client

- 完成 gRPC 的相关业务调用
- 带有日志统计中间件
- 以流(stream) 的方式取回数据

### interceptor

- 中间件模块实现

## Generate Code

    $ protoc -I . --go_out=plugins=grpc:. proto/app.proto

## 译文勘误：

1. client.go: Server 服务运行在 `5051` 端口

        grpc.Dial("localhost:5053")

   changed to:

        grpc.Dial("localhost:5051")`

2. server.go: CacheService.store：map 未初始化就分配值会导致 panic

        s.store[req.Key] = req.Val

    changed to:

    	if s.store == nil {
    		s.store = make(map[string][]byte)
    	}
    	s.store[req.Key] = req.Val

[#1]:https://about.sourcegraph.com/go/grpc-in-production-alan-shreve/
[#2]:https://blog.lab99.org/post/golang-2017-09-27-video-grpc-from-tutorial-to-production.html?utm_campaign=studygolang.com&utm_medium=studygolang.com&utm_source=studygolang.com#shi-pin-xin-xi