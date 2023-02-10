package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main3(title string) {
	myApp := app.New()
	myWindow := myApp.NewWindow(title)
	myWindow.Resize(fyne.NewSize(400, 300))

	myWindow.SetContent(widget.NewLabel("This is some content."))

	myWindow.ShowAndRun()
}
