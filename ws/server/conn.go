package server

/**
* @Author: Jam Wong
* @Date: 2020/6/3 12:32 下午
 */

import (
	"errors"
	"github.com/gorilla/websocket"
	"sync"
	"time"
)

type wsConn struct {
	lock                  sync.Mutex
	connId                uint64
	wsSocket              *websocket.Conn
	inChan                chan *WsMessage
	outChan               chan *WsMessage
	closeChan             chan byte
	isClosed              bool
	lastHeartbeatTime     time.Time
	heartbeatTimeDuration time.Duration
}

type WsMessage struct {
	MsgType int
	MsgData []byte
}

// InitWsConn init a websocket connection
func InitWsConn(wsSocket *websocket.Conn, heartbeatTimeDuration time.Duration) *wsConn {
	w := &wsConn{
		wsSocket:              wsSocket,
		inChan:                make(chan *WsMessage, 1000), // TODO
		outChan:               make(chan *WsMessage, 1000),
		closeChan:             make(chan byte),
		lastHeartbeatTime:     time.Now(),
		heartbeatTimeDuration: heartbeatTimeDuration,
	}

	go w.readLoop()
	go w.writeLoop()
	return w
}

// readLoop
func (w *wsConn) readLoop() {
	var (
		msgType int
		msgData []byte
		message *WsMessage
		err     error
	)
	for {
		if msgType, msgData, err = w.wsSocket.ReadMessage(); err != nil {
			w.Close()
			break
		}

		message = &WsMessage{MsgData: msgData, MsgType: msgType}
		select {
		case w.inChan <- message:
		case <-w.closeChan:
			break
		}
	}
}

// writeLoop
func (w *wsConn) writeLoop() {
	var (
		message *WsMessage
		err     error
	)
	for {
		select {
		case message = <-w.outChan:
			if err = w.wsSocket.WriteMessage(message.MsgType, message.MsgData); err != nil {
				w.Close()
				break
			}
		case <-w.closeChan:
			break
		}
	}
}

// readMsg
func (w *wsConn) readMsg() (msg *WsMessage, err error) {
	select {
	case msg = <-w.inChan:
	case <-w.closeChan:
		err = errors.New("Connection Loss")
	}
	return
}

// sendMsg
func (w *wsConn) sendMsg(msg *WsMessage) (err error) {
	select {
	case w.outChan <- msg:
	case <-w.closeChan:
		err = errors.New("Connection Loss")
	default: // 写操作不会阻塞, 因为channel已经预留给websocket一定的缓冲空间
		err = errors.New("Connection full")
	}
	return
}

// Close
func (w *wsConn) Close() {
	w.wsSocket.Close()
	w.lock.Lock()
	defer w.lock.Unlock()

	if !w.isClosed {
		w.isClosed = true
		close(w.closeChan)
	}
}

// IsAlive
func (w *wsConn) IsAlive() bool {
	w.lock.Lock()
	defer w.lock.Unlock()
	// 连接已关闭 或者 太久没有心跳
	if w.isClosed || time.Now().Sub(w.lastHeartbeatTime) > w.heartbeatTimeDuration {
		return false
	}
	return true
}

// KeepAlive
func (w *wsConn) KeepAlive() {
	w.lock.Lock()
	defer w.lock.Unlock()
	w.lastHeartbeatTime = time.Now()
}
