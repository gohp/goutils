package socks5

import "io"

/**
* @Author: Jam Wong
* @Date: 2020/7/22
 */

func transfer(dst io.Writer, src io.Reader) {
	// TODO 流量加密
	io.Copy(dst, src)
}
