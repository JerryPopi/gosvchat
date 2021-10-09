package main

import (
	"fmt"
	"strings"
)

func parseCommand(input string){
	split := strings.Fields(input)
	switch split[0] {
	case "help":
		statusBar.SetText("Available commands: :help, :rename, :quit")
	case "rename":
		if len(split) < 2 {
			statusBar.SetText("Incorrect usage: :rename <name>")
		} else {
			Config.Client.Name = strings.Join(split[1:], " ")
			statusBar.SetText("Changed name to " + Config.Client.Name)
		}
	case "user":
		statusBar.SetText(Config.Client.Name)
	case "info":
		fmt.Fprintf(textView, "%+v\n", Config)
	default:
		statusBar.SetText("Unknown command. Use :help")
	}
}