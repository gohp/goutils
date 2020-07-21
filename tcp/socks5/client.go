package socks5

import (
	"log"
	"net"
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

func (client *TcpClient) handleRequest(localClient *net.TCPConn, serverAddr *net.TCPAddr)  {
	panic("no implement")
}