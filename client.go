package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

var Client struct {
	Name string
}

var c net.Conn

func startClient(username string, addr string) {
	if username == "" {
		fmt.Print("Please enter username: ")
		reader := bufio.NewReader(os.Stdin)
		username, _ := reader.ReadString('\n')
		Client.Name = strings.TrimSuffix(username, "\n")
	} else {
		Client.Name = username
	}


	conn := addr
	var err error
	c, err = net.Dial("tcp", conn)
	chk(err)
	ctrlCHandlerClient(c)

	createUI(c)
}
// 	go writer(c)

// 	for {
// 		reader := bufio.NewReader(os.Stdin)
// 		fmt.Print(">> ")
// 		text, _ := reader.ReadString('\n')
// 		if text != "" {
// 			message := Message{Content: text, Name: Client.Name, Timestamp: time.Now()}
// 			// fmt.Fprintf(c, text+"\n")
// 			enc := gob.NewEncoder(c)
// 			enc.Encode(message)
// 		}
// 	}

// }

// func writer(c net.Conn) {
// 	defer c.Close()

// 	var message Message

// 	for {
// 		dec := gob.NewDecoder(c)
// 		dec.Decode(&message)
// 		if message.Name != Client.Name {
// 			fmt.Print(message.Content)
// 		}
// 		time.Sleep(50 * time.Millisecond)
// 	}

// }
