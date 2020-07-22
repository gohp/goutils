package socks5

import (
	"encoding/binary"
	"io"
	"log"
	"net"
	"sync"
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

func (server *TcpServer) handleRequest(conn *net.TCPConn) {
	log.Println("client recv ")
	buf := make([]byte, 263)
	n, _ := io.ReadAtLeast(conn, buf, 5)

	var dstIP []byte
	switch buf[0] {
	case 0x01: // ipv4
		dstIP = buf[1 : net.IPv4len+1]
	case 0x03: // domain
		ipAddr, err := net.ResolveIPAddr("ip", string(buf[2:n-2]))
		if err != nil {
			return
		}
		dstIP = ipAddr.IP
	case 0x04: // ipv6
		dstIP = buf[1 : net.IPv6len+1]
	default:
		return
	}
	// port
	dstPort := buf[n-2:]
	dstAddr := &net.TCPAddr{
		IP:   dstIP,
		Port: int(binary.BigEndian.Uint16(dstPort)),
	}

	client, _ := net.DialTCP("tcp", nil, dstAddr)

	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {
		defer wg.Done()
		transfer(conn, client)
	}()

	go func() {
		defer wg.Done()
		transfer(client, conn)
	}()

	wg.Wait()
}
