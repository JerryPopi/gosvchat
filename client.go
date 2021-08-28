package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
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

	go createUI()

	go writer(c)
	// go reader(c)

	for{
		time.Sleep(50 * time.Millisecond)
	}
}

// func reader(c net.Conn){
// 	defer c.Close()

// 	enc := gob.NewEncoder(c)

// 	for {
// 		if messageText != "" {
// 			message := Message{Content: messageText, Name: Client.Name, Timestamp: time.Now()}
// 			enc.Encode(message)
// 		}

// 		time.Sleep(20 * time.Millisecond)
// 	}
// }

func writer(c net.Conn) {
	defer c.Close()
	dec := gob.NewDecoder(c)

	for {
		var message Message
		// fmt.Println("POLYCHIH SAOBSHTENIE")
		if err := dec.Decode(&message); err != nil {
			// fmt.Println("ERR " + err.Error())
		} else {
			appendMessage(message)
		}
		time.Sleep(20 * time.Millisecond)
	}
}

func parseInput(input string){
	enc := gob.NewEncoder(c)

	message := Message{Content: input, Name: Client.Name, Timestamp: time.Now()}
	enc.Encode(message)

	inputField.SetText("")
}