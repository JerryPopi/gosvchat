package main

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var (
	grid       *tview.Grid
	textView   *tview.TextView
	inputField *tview.InputField
	app        *tview.Application
	statusBar  *tview.TextView
)

// todo configs
// ? input history (nep)

func createUI() {
	fmt.Println(Config.Client)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()

	defer os.Exit(0)

	app = tview.NewApplication()

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
			parseInput(inputField.GetText())
		}
	}

	inputField.
		SetLabel("> ").
		SetLabelColor(tcell.ColorRed).
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

	if err := app.SetRoot(grid, true).EnableMouse(false).Run(); err != nil {
		panic(err)
	}
}

func appendMessage(message Message) {
	fmt.Fprintf(textView, "%+v\n", message)
	if message.Name != Config.Client.Name {
		if !Config.Env.OverrideCustomColors {
			//					  rcc usr ptr msg
			color := "[" + message.CustomColor + "]"
			// fmt.Fprintf(textView, "%s%s %s[white]\n", color, message.Name, message.Content)
			fmt.Fprint(textView, color, message.Name, " ", message.Content, "[white]\n")
		} else {
			color := "[" + Config.Env.RemoteColor + "]"
			// fmt.Fprintf(textView, "%s%s %s[white]\n", color, message.Name, message.Content)
			fmt.Fprint(textView, color, message.Name, " ", message.Content, "[white]\n")
		}

	} else {
		color := "[" + Config.Env.LocalColor + "]"
		// 					   lc   msg
		// fmt.Fprintf(textView, "[%s] %s[white]\n", Config.Env.LocalColor, message.Content)
		fmt.Fprint(textView, color, message.Content, "[white]\n")


	}
}
