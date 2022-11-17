package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"time"
)

// clock
func updateTime(clock *widget.Label) {
	formatted := time.Now().Format("Time: 03:04:05")
	clock.SetText(formatted)
}

// test ui
func makeUI() (*widget.Label, *widget.Entry) {
	label := widget.NewLabel("Hello World")
	entry := widget.NewEntry()

	entry.OnChanged = func(content string) {
		label.SetText("Hello " + content)
	}
	return label, entry
}

func main() {
	myApp := app.New()

	// hello world
	/*myWindow := myApp.NewWindow("Hello World")
	myWindow.SetContent(widget.NewLabel("This is some content."))*/

	// clock
	/*myWindow := myApp.NewWindow("Clock")
	clock := widget.NewLabel("")
	updateTime(clock)
	myWindow.SetContent(clock)
	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()*/

	// open window
	/*myWindow := myApp.NewWindow("New Window")
	myWindow.SetContent(widget.NewButton("Open Window", func() {
		w2 := myApp.NewWindow("Window 2")
		w2.SetContent(widget.NewLabel("This is other window."))
		w2.Show()
	}))*/

	// canvas
	/*myWindow := myApp.NewWindow("Canvas")
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
	myCanvas.SetContent(content)*/

	// test ui
	/*myWindow := myApp.NewWindow("Test UI")
	myWindow.SetContent(container.NewVBox(makeUI()))*/

	// system tray
	/*myWindow := myApp.NewWindow("System Tray")

	if desktopApp, ok := myApp.(desktop.App); ok {
		myTrayMenu := fyne.NewMenu("Tray Menu",
			fyne.NewMenuItem("Show", func() {
				myWindow.Show()
			}),
		)
		desktopApp.SetSystemTrayMenu(myTrayMenu)
	}

	myWindow.SetCloseIntercept(func() {
		myWindow.Hide()
	})
	myWindow.SetContent(widget.NewLabel("This is system tray."))*/

	// layout | https://developer.fyne.io/explore/layouts
	myWindow := myApp.NewWindow("Layout")
	myWindow.SetContent(container.NewBorder(
		container.NewHBox(widget.NewLabel("top1"), widget.NewLabel("top2")),
		container.NewCenter(widget.NewLabel("bottom")),
		container.NewVBox(widget.NewLabel("left1"), widget.NewLabel("left2")),
		widget.NewLabel("right"),
		container.NewGridWithColumns(2,
			widget.NewLabel("go1"),
			widget.NewLabel("go2"),
			widget.NewLabel("go3"),
		),
	))

	// widget | https://developer.fyne.io/explore/widgets
	//myWindow.SetContent(widget.NewEntry())

	myWindow.Resize(fyne.NewSize(400, 300))
	myWindow.SetMaster()
	//myWindow.ShowAndRun()
	myWindow.Show()

	myApp.Run()
	tidyUp()
}

func tidyUp() {
	fmt.Println("Exited")
}
