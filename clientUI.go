package main

import (
	"fmt"

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

func createUI() {
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println("stacktrace from panic: \n" + string(debug.Stack()))
	// 	}
	// }()

	// setInputFocus := func(tview.Primitive) {
	// 	app.SetFocus(inputField)
	// }

	// defer os.Exit(0);

	defer fmt.Println("Returned from clientUI.go")

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
			// textView.Write([]byte(inputField.GetText()))
			// fmt.Fprintf(textView, "%s: %s\n", Client.Name, inputField.GetText())
			parseInput(inputField.GetText())
			// fmt.Fprintf(textView, "%s\n", inputField.GetText())
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

	if err := app.SetRoot(grid, true).EnableMouse(false).Run(); err != nil {
		panic(err)
	}

	// textView.Focus(setInputFocus)
	// statusBar.Focus(setInputFocus)

}

func appendMessage(message Message){
	// textView.Write([]byte(message.Content))
	// fmt.Println(message)
	fmt.Fprintf(textView, "%s: %s\n", message.Name, message.Content)
	// fmt.Println(c)
	// fmt.Println(message.Content)
	// fmt.Println(message.Content)
	// textView.SetText(message.Content)
}