package socks5

import (
	"errors"
	"io"
	"log"
	"net"
	"sync"
)

/**
* @Author: Jam Wong
* @Date: 2020/7/21
 */

type ProxyClient struct {
	listenAddr string
	serverAddr string
}

func NewClient(listenAddr string, serverAddr string) *ProxyClient {
	c := &ProxyClient{
		listenAddr: listenAddr,
		serverAddr: serverAddr,
	}
	return c
}

func (client *ProxyClient) Run() {
	// proxy地址
	serverAddr, err := net.ResolveTCPAddr("tcp", client.serverAddr)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("connect to remote server: %s ....", client.serverAddr)

	listenAddr, err := net.ResolveTCPAddr("tcp", client.listenAddr)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("listen on: %s ", client.listenAddr)

	listener, err := net.ListenTCP("tcp", listenAddr)
	if err != nil {
		log.Fatal(err)
	}

	for {
		localClient, err := listener.AcceptTCP()
		if err != nil {
			log.Println(err)
			continue
		}
		go client.handleRequest(localClient, serverAddr)
	}
}

func (client *ProxyClient) replay(conn net.Conn, data []byte) error {
	_, err := conn.Write(data)
	if err != nil {
		log.Printf("replay bytes err: %v\n", err)
	}
	return err
}

func (client *ProxyClient) replayAuth(conn net.Conn) error {
	// not auth
	// reply
	//    +----+--------+
	//    |VER | METHOD |
	//    +----+--------+
	//    | 1  |   1    |
	//    +----+--------+
	return client.replay(conn, []byte{0x05, 0x00})
}

func (client *ProxyClient) replayConnect(conn net.Conn) error {
	// not auth
	return client.replay(conn, []byte{0x05, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
}

func (client *ProxyClient) unpackAuth(conn net.Conn) error {
	buf := make([]byte, 263)
	n, err := io.ReadAtLeast(conn, buf, 2)
	if err != nil {
		return err
	}

	// |VER | NMETHODS | METHODS  |
	// +----+----------+----------+
	// | 5  |    1     | 1 to 255 |
	// +----+----------+----------+
	if buf[0] != 0x05 {
		return errors.New("only support socks5")
	}

	nMethod := int(buf[1])
	msgLen := nMethod + 2
	if n < msgLen {
		if _, err = io.ReadFull(conn, buf[n:msgLen]); err != nil {
			return err
		}
	} else if n > msgLen {
		return errors.New("todo")
	}

	return nil
}

func (client *ProxyClient) unpackConnecting(conn net.Conn) ([]byte, int, error) {
	// +----+-----+-------+------+----------+----------+
	// |VER | CMD |  RSV  | ATYP | DST.ADDR | DST.PORT |
	// +----+-----+-------+------+----------+----------+
	// | 1  |  1  | X'00' |  1   | Variable |    2     |
	// +----+-----+-------+------+----------+----------+
	var min int
	buf := make([]byte, 263)
	min, err := io.ReadAtLeast(conn, buf, 5)
	if err != nil {
		return nil, min, err
	}
	if buf[0] != 0x05 || buf[1] != 0x01 {
		return nil, min, errors.New("only support socks5")
	}
	return buf, min, nil
}

func (client *ProxyClient) handleRequest(localClient *net.TCPConn, serverAddr *net.TCPAddr) {
	defer localClient.Close()
	log.Println("client -> server")

	// 1. 认证阶段
	// client -> server: 0x05 0x01 0x00
	// server -> client: 0x05 0x00
	if err := client.unpackAuth(localClient); err != nil {
		log.Printf("unpackAuth error: %v\n", err)
		return
	}
	client.replayAuth(localClient)

	// 2. 连接阶段
	buf, n, err := client.unpackConnecting(localClient)
	if err != nil {
		log.Printf("unpackConnecting error: %v\n", err)
		return
	}

	// 3. 传输阶段
	dstServer, err := net.DialTCP("tcp", nil, serverAddr)
	if err != nil {
		log.Printf("remote addr error: %v\n", err)
		return
	}
	defer dstServer.Close()

	dstServer.Write(buf[3:n])
	_ = client.replayConnect(localClient)

	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {
		defer wg.Done()
		transfer(localClient, dstServer)
	}()

	go func() {
		defer wg.Done()
		transfer(dstServer, localClient)
	}()

	wg.Wait()
}
