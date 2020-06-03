package server

/**
* @Author: Jam Wong
* @Date: 2020/6/3 12:32 下午
 */

import (
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

type WsHandleFunc func([]byte) []byte

func (w *wsConn) heartbeat() {
	timer := time.NewTimer(w.heartbeatTimeDuration)
	for {
		select {
		case <-timer.C:
			if !w.IsAlive() {
				w.Close()
				break
			}
			timer.Reset(w.heartbeatTimeDuration)
		case <-w.closeChan:
			timer.Stop()
			break
		}
	}
}

func (w *wsConn) Handle(f WsHandleFunc) {
	var (
		err     error
		message *WsMessage
		resp    *WsMessage
	)
	go w.heartbeat()
	for {
		if message, err = w.readMsg(); err != nil {
			w.Close()
			break
		}

		// 只处理文本消息
		//if message.MsgType != websocket.TextMessage {
		//	continue
		//}
		w.KeepAlive()
		switch message.MsgType {
		// 处理ping请求
		case websocket.PingMessage:
			resp = &WsMessage{MsgType: websocket.PongMessage, MsgData: nil}
		case websocket.TextMessage:
			resp = &WsMessage{MsgType: websocket.TextMessage, MsgData: f(message.MsgData)}
		default:
			continue
		}

		// TODO
		if err = w.sendMsg(resp); err != nil {
			w.Close()
			break
		}
	}
	return
}

func InitWsUpGrader() websocket.Upgrader {
	var upGrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		// allow CORS
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	return upGrader
}

// wsHandler net/http handler
func WsHandlerHttp(f WsHandleFunc) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		upGrader := InitWsUpGrader()
		wsSocket, err := upGrader.Upgrade(w, r, nil)
		if err != nil {
			panic(err)
		}

		wsConn := InitWsConn(wsSocket, time.Second*60)
		wsConn.Handle(f)
	}
}
