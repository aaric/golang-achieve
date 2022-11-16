package main

import (
	"fmt"
	"fyne.io/fyne/v2"
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
	myWindow.Resize(fyne.NewSize(400, 300))
	myWindow.SetMaster()

	/*clock := widget.NewLabel("")
	updateTime(clock)
	//myWindow.SetContent(widget.NewLabel("This is some content."))
	myWindow.SetContent(clock)
	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()*/

	myWindow.SetContent(widget.NewButton("Open Window", func() {
		w2 := myApp.NewWindow("Window 2")
		w2.SetContent(widget.NewLabel("This is other windows."))
		w2.Show()
	}))

	//myWindow.ShowAndRun()
	myWindow.Show()

	myApp.Run()
	tidyUp()
}

func tidyUp() {
	fmt.Println("Exited")
}
