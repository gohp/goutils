package socks5

import (
	"log"
	"net"
)

/**
* @Author: Jam Wong
* @Date: 2020/7/21
 */

type TcpServer struct {
	listenAddr string
}

func NewServer(addr string) *TcpServer {
	s := &TcpServer{
		listenAddr: addr,
	}
	return s
}

func (server *TcpServer) Run() {
	// listen client
	listenAddr, err := net.ResolveTCPAddr("tcp", server.listenAddr)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("listen on: %s ", server.listenAddr)

	listener, err := net.ListenTCP("tcp", listenAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Fatal(err)
		}
		go server.handleRequest(conn)
	}
}

func (server *TcpServer) handleRequest(conn *net.TCPConn)  {
	panic("no implement")
}