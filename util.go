package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
)

func chk(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func ctrlCHandlerClient(cn net.Conn) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			if sig == os.Interrupt {
				cn.Close()
				fmt.Println("Stopping client.")
				os.Exit(0)
			}
		}
	}()
}

func ctrlCHandler() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			if sig == os.Interrupt {
				fmt.Println("Stopping server.")
				os.Exit(0)
			}
		}
	}()
}

func indexOfAddr(element net.Addr, data []net.Conn) int {
	for k, v := range data {
		if element == v.RemoteAddr() {
			return k
		}
	}
	return -1
}

func removeClient(i int) {
	clients[i] = clients[len(clients)-1]
	clients = clients[:len(clients)-1]
}
