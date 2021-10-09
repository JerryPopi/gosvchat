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

// var Client struct {
// 	Name string
// }

var c net.Conn

func startClient(username string, addr string) {
	if username == "" {
		if Config.Client.Name == "" {
			fmt.Print("Please enter username: ")
			reader := bufio.NewReader(os.Stdin)
			username, _ := reader.ReadString('\n')
			Config.Client.Name = strings.TrimSuffix(username, "\n")
		}
	}

	conn := addr
	var err error
	c, err = net.Dial("tcp", conn)
	chk(err)

	go createUI()
	go writer(c)

	for {
		time.Sleep(50 * time.Millisecond)
	}
}

func writer(c net.Conn) {
	defer c.Close()
	dec := gob.NewDecoder(c)

	for {
		var message Message

		if err := dec.Decode(&message); err != nil {

		} else {
			appendMessage(message)
		}
		time.Sleep(20 * time.Millisecond)
	}
}

func parseInput(input string) {
	statusBar.Clear()
	if string([]byte(input)[0]) == ":" {
		if len(input) > 1 {
			parseCommand(input[1:])
		}
		inputField.SetText("")
	} else {
		sendMessage(input)
		inputField.SetText("")
	}
}

func sendMessage(input string) {
	enc := gob.NewEncoder(c)

	message := Message{Content: input, Name: Config.Client.Name, CustomColor: Config.Client.CustomColor, Timestamp: time.Now()}
	enc.Encode(message)
}
