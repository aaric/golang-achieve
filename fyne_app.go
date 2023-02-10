package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"golang-achieve/theme"
)

func main3(title string) {
	myApp := app.New()
	myApp.Settings().SetTheme(&theme.MyTheme{})
	myWindow := myApp.NewWindow(title)
	myWindow.Resize(fyne.NewSize(400, 300))

	//content := widget.NewForm(widget.NewFormItem("用户名：", widget.NewEntry()))
	//myWindow.SetContent(content)
	myWindow.SetContent(widget.NewLabel("中国"))

	myWindow.ShowAndRun()
}
