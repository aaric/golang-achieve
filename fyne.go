package main

import (
	"fmt"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"time"
)

func updateTime(clock *widget.Label) {
	formatted := time.Now().Format("Time: 03:04:05")
	clock.SetText(formatted)
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Clock")

	clock := widget.NewLabel("")
	updateTime(clock)
	//myWindow.SetContent(widget.NewLabel("This is some content."))
	myWindow.SetContent(clock)
	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()

	myWindow.ShowAndRun()
	tidyUp()
}

func tidyUp() {
	fmt.Println("Exited")
}
