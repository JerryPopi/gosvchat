package main

import (
	// "encoding/gob"
	"encoding/gob"
	"fmt"
	"net"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var (
	conn net.Conn
	grid* tview.Grid
	textView* tview.TextView
	inputField* tview.InputField
	app* tview.Application
	statusBar* tview.TextView
)

func createUI(c net.Conn) {
	conn = c

	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println("stacktrace from panic: \n" + string(debug.Stack()))
	// 	}
	// }()

	setInputFocus := func(tview.Primitive) {
		app.SetFocus(inputField)
	}

	
	app = tview.NewApplication()
	
	// app.SetAfterDrawFunc(setInputFocus)
	
	textView = tview.NewTextView().
	SetDynamicColors(true).
	SetRegions(false).
	SetWordWrap(true).
	SetChangedFunc(func() {
		app.Draw()
	})
	
	inputField = tview.NewInputField()
	inputDone := func(event tcell.Key) {
		switch event {
		case tcell.KeyEnter:
			parseInput()
		}
	}
	inputField.
	SetLabel("> ").
	SetFieldWidth(0).
	SetFieldTextColor(tcell.ColorWhite).
	SetFieldBackgroundColor(tcell.ColorBlack).
	SetDoneFunc(inputDone)
	
	statusBar = tview.NewTextView().
	SetDynamicColors(true).
	SetRegions(false).
	SetWordWrap(false).
	SetTextColor(tcell.ColorGhostWhite).
	SetChangedFunc(func() {
		app.Draw()
	})
	
	grid = tview.NewGrid().
	SetRows(0, 1, 1).
	SetColumns(0).
	AddItem(textView, 0, 0, 1, 1, 0, 0, false).
	AddItem(inputField, 1, 0, 1, 1, 0, 0, true).
	AddItem(statusBar, 2, 0, 1, 1, 0, 0, false)
	
	if err := app.SetRoot(grid, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

	textView.Focus(setInputFocus)
	statusBar.Focus(setInputFocus)
	
	go reciever()
}

func parseInput() {
	statusBar.Clear()
	input := inputField.GetText()
	fmt.Print(input)
	if string([]byte(input)[0]) == ":" {
		if(len(input) > 1){
			parseCommand(input[1:])
		} 	
		inputField.SetText("")
	} else {
		sendMessage(input)
		fmt.Fprintf(textView, "> %s\n", input)
		statusBar.SetText(input)
		inputField.SetText("")
	}
}

func sendMessage(text string) {
	if text != "" {
		message := Message{Content: text, Name: Client.Name, Timestamp: time.Now()}
		enc := gob.NewEncoder(conn)
		err := enc.Encode(message)
		chk(err)
	}
}

func reciever() {
	defer conn.Close()

	var message Message

	for {
		dec := gob.NewDecoder(conn)
		dec.Decode(&message)
		if message.Name != Client.Name {
			fmt.Print(message.Content)
		}
		time.Sleep(50 * time.Millisecond)
	}
}