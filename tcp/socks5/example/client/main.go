package main

import (
	"github.com/wzyonggege/goutils/choke"
	"github.com/wzyonggege/goutils/tcp/socks5"
)

/**
* @Author: Jam Wong
* @Date: 2020/7/22
 */

func main() {
	client := socks5.NewClient("127.0.0.1:10201", "127.0.0.1:10200")
	go client.Run()
	choke.Choke()
}
