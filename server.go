package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

var clients []net.Conn

func startServer(port string) {
	ctrlCHandler()

	if port == "" {
		fmt.Println("Please provide port number.")
		return
	}

	PORT := ":" + port
	l, err := net.Listen("tcp", PORT)
	chk(err)
	defer l.Close()

	for {
		c, err := l.Accept()
		chk(err)
		fmt.Println(c.RemoteAddr().String() + " connected!")
		clients = append(clients, c)
		fmt.Println(clients)

		go handleConnection(c)
	}
}

func handleConnection(c net.Conn) {
	defer c.Close()
	
	var netData Message
	for {
		dec := gob.NewDecoder(c)
		err := dec.Decode(&netData)
		if chkDisconnect(c, err) {
			removeClient(indexOfAddr(c.RemoteAddr(), clients))
			return
		}

		fmt.Print("-> ", string(netData.Content) + " (" + netData.Name + ") ")

		for _, client := range clients {
			enc := gob.NewEncoder(client)
			enc.Encode(netData)
		}
	}
}

func chkDisconnect(c net.Conn, err error) bool {
	fmt.Print("chkdisconnect: ")
	fmt.Println(clients)
	if err != nil {
		fmt.Println(c.RemoteAddr().String() + " disconnected.")
		return true
	}
	return false
}
