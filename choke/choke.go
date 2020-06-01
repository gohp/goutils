package choke

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Choke() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Println("os exit")
			return
		case syscall.SIGHUP:
		// TODO reload
		default:
			return
		}
	}
}
