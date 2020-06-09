package main

import (
	"github.com/wzyonggege/goutils/ecode"
	"github.com/wzyonggege/goutils/http"
	"log"
)

/**
* @Author: Jam Wong
* @Date: 2020/6/9
 */

var (
	NetworkErr = ecode.New(10001, "network error")
)

func main()  {
	_, err := http.HttpGet("baidu.com")
	if err != nil {
		log.Fatal(NetworkErr)
	}
}