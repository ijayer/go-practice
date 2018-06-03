##### Gorilla Websocket Test

> **Note**
> 
> 0. client登陆，成功进入step-1
> 1. client 发送握手连接(携带用户信息：id, platform)
> 2. 注册client到connection中
> 3. 广播消息
> 4. 换平台登陆，服务器给已登陆终端发送下线通知