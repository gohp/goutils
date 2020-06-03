# Websocket

base on "github.com/gorilla/websocket"


## Client

[官方demo](https://github.com/gorilla/websocket/tree/master/examples/echo)

## Server

### Usage

```go

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wzyonggege/goutils/ws/server"
	"net/http"
)

// ws text msg 规则
func wsRule(msg []byte) []byte {
	// 编写 规则
	return nil
}

// http server
func InitHttpServer() {
	http.HandleFunc("/ws", server.WsHandlerHttp(func(msg []byte) []byte {
		return append([]byte("sssss"), msg...)
	}))

	http.HandleFunc("/ws2", server.WsHandlerHttp(wsRule))

	if err := http.ListenAndServe(":7778", nil); err != nil {
		panic(err)
	}
}

// gin server
func InitGinServer() (err error) {
	g := gin.New()
	g.Use(gin.Recovery())

	g.GET("/ws", func(c *gin.Context) {
		f := server.WsHandlerHttp(func(msg []byte) []byte {
			return append([]byte("sssss"), msg...)
		})
		f(c.Writer, c.Request, )
	})
	err = http.ListenAndServe(":7777", g)
	return
}

```