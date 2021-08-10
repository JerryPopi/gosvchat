package main

import (
	"encoding/gob"
	"fmt"
	"net"
	// "strconv"
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
	for {
		var netData Message
		dec := gob.NewDecoder(c)
		err := dec.Decode(&netData)
		if chkDisconnect(c, err) {
			removeClient(indexOfAddr(c.RemoteAddr(), clients))
			return
		}

		for i, cl := range clients {
			fmt.Print("-> ", string(netData.Content) + " ")
			fmt.Print(i)
			fmt.Println(" > " + cl.RemoteAddr().String() + " (" + netData.Name + ")")
			enc := gob.NewEncoder(cl)
			enc.Encode(netData)
		}
	}
}

func chkDisconnect(c net.Conn, err error) bool {
	fmt.Println(clients)
	if err != nil {
		fmt.Println(c.RemoteAddr().String() + " disconnected.")
		// i := indexOfAddr(c.RemoteAddr(), clients)
		// debug(strconv.Itoa(i))
		// if i == -1 {
		// 	// fmt.Println("Error: cant fint address")
		// 	return true
		// } else {
		// 	// removeClient(i)
		// 	return true
		// }

		return true

		// c.Close()
		// return
	}
	return false
}

// func debug(str string){
// 	fmt.Println("DEBUG: " + "\033[33m" + str + "\033[0m")
// }