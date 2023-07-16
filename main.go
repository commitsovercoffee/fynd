package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// creates new instance of app.
	a := app.New()

	// creates new window for app instance.
	w := a.NewWindow("Hello World")

	// widget tree for the window.
	w.SetContent(widget.NewLabel("Hello World!"))

	// run the app.
	w.ShowAndRun()
}
