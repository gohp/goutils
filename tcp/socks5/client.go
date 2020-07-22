package socks5

import (
	"io"
	"log"
	"net"
	"sync"
)

/**
* @Author: Jam Wong
* @Date: 2020/7/21
 */

type TcpClient struct {
	listenAddr string
	serverAddr string
}

func NewClient(listenAddr string, serverAddr string) *TcpClient {
	c := &TcpClient{
		listenAddr: listenAddr,
		serverAddr: serverAddr,
	}
	return c
}

func (client *TcpClient) Run() {
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
			log.Fatal(err)
		}
		go client.handleRequest(localClient, serverAddr)
	}
}

func (client *TcpClient) handleRequest(localClient *net.TCPConn, serverAddr *net.TCPAddr) {
	log.Println("client transfer to server")

	buf := make([]byte, 263)
	n, err := io.ReadAtLeast(localClient, buf, 2)
	if err != nil {
		return
	}

	// only support socks5
	if buf[0] != 0x05 {
		return
	}

	nMethod := int(buf[1])
	msgLen := nMethod + 2
	if n < msgLen {
		if _, err = io.ReadFull(localClient, buf[n:msgLen]); err != nil {
			return
		}
	} else if n > msgLen {
		return
	}

	/*
		告诉客户端 不需要验证
		+----+--------+
		|VER | METHOD |
		+----+--------+
		| 1  |   1    |
		+----+--------+
	*/
	localClient.Write([]byte{0x05, 0x00})
	if n, err = io.ReadAtLeast(localClient, buf, 5); err != nil {
		return
	}
	if buf[0] != 0x05 {
		return
	}
	if buf[1] != 0x01 {
		return
	}

	dstServer, err := net.DialTCP("tcp", nil, serverAddr)
	if err != nil {
		log.Print("remote addr error")
		log.Print(err)
		return
	}
	defer dstServer.Close()
	defer localClient.Close()

	dstServer.Write(buf[3:n])
	localClient.Write([]byte{0x05, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})

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
