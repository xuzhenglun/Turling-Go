#图灵机器人
---
Web service的作业

调用了图灵机器人的API，参照了http://github.com/domluna/websocket-golang-chat的代码。

#Useage

```
go build chat.go
openssl genrsa -out key.pem 2048
openssl req -new -x509 -key key.pem -out cert.pem -days 1095
./chat.go
```

#Config

如果需要被别的机器访问，必须设置公开地址，否则客户无法找到服务器。

```
{
  	"key": "abcdefg"
  	"address": "Http://www.tuling123.com/openapi/api"
    "listenAddr":"localhost:4000"
    "publicAddr":"192.168.100.108"
}
```
