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

const (
	addrDomain = 0x03
	addrIpv4   = 0x01
	addrIpv6   = 0x04
)

type ProxyServer struct {
	listenAddr string
}

func NewServer(addr string) *ProxyServer {
	s := &ProxyServer{
		listenAddr: addr,
	}
	return s
}

func (server *ProxyServer) Run() {
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
			log.Println(err)
			continue
		}
		go server.handleRequest(conn)
	}
}

func (server *ProxyServer) handleRequest(conn *net.TCPConn) {
	buf := make([]byte, 263)
	n, _ := io.ReadAtLeast(conn, buf, 5)

	var dstIP []byte
	switch buf[0] {
	case addrIpv4: // ipv4
		dstIP = buf[1 : net.IPv4len+1]
	case addrDomain: // domain
		ipAddr, err := net.ResolveIPAddr("ip", string(buf[2:n-2]))
		if err != nil {
			return
		}
		dstIP = ipAddr.IP
	case addrIpv6: // ipv6
		dstIP = buf[1 : net.IPv6len+1]
	default:
		return
	}
	// port
	dstPort := buf[n-2:]
	port := int(binary.BigEndian.Uint16(dstPort))
	log.Printf("recv request %s:%d\n", string(dstIP), port)

	client, _ := net.DialTCP("tcp", nil, &net.TCPAddr{IP: dstIP, Port: port})

	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {
		defer wg.Done()
		defer func() {
			if err := recover(); err != nil {
				log.Println("panic recover", err)
			}
		}()
		transfer(conn, client)
	}()

	go func() {
		defer wg.Done()
		defer func() {
			if err := recover(); err != nil {
				log.Println("panic recover", err)
			}
		}()
		transfer(client, conn)
	}()

	wg.Wait()
}
