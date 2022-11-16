package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"time"
)

// clock
func updateTime(clock *widget.Label) {
	formatted := time.Now().Format("Time: 03:04:05")
	clock.SetText(formatted)
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Demo")
	myWindow.Resize(fyne.NewSize(400, 300))
	myWindow.SetMaster()

	// hello world
	/*myWindow.SetContent(widget.NewLabel("This is some content."))*/

	// clock
	/*clock := widget.NewLabel("")
	updateTime(clock)
	myWindow.SetContent(clock)
	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()*/

	// open window
	/*myWindow.SetContent(widget.NewButton("Open Window", func() {
		w2 := myApp.NewWindow("Window 2")
		w2.SetContent(widget.NewLabel("This is other windows."))
		w2.Show()
	}))*/

	// canvas
	myCanvas := myWindow.Canvas()
	blue := color.NRGBA{
		R: 0,
		G: 0,
		B: 180,
		A: 255,
	}
	rect := canvas.NewRectangle(blue)
	go func() {
		time.Sleep(time.Second)
		green := color.NRGBA{
			R: 0,
			G: 180,
			B: 0,
			A: 255,
		}
		rect.FillColor = green
		rect.Refresh()
	}()
	myCanvas.SetContent(rect)

	circle := canvas.NewCircle(blue)
	circle.StrokeWidth = 4
	circle.StrokeColor = blue
	myCanvas.SetContent(circle)

	text := canvas.NewText("hello world", color.Black)
	text.TextStyle.Bold = true

	//content := container.NewWithoutLayout(rect, circle, text)
	content := container.New(layout.NewVBoxLayout(), rect, circle, text)
	myCanvas.SetContent(content)

	// layout | https://developer.fyne.io/explore/layouts

	// widget | https://developer.fyne.io/explore/widgets
	//myWindow.SetContent(widget.NewEntry())

	//myWindow.ShowAndRun()
	myWindow.Show()

	myApp.Run()
	tidyUp()
}

func tidyUp() {
	fmt.Println("Exited")
}
