package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"golang-achieve/theme"
	"log"
)

func main3(title string) {
	myApp := app.New()
	myApp.Settings().SetTheme(&theme.MyTheme{})
	myWindow := myApp.NewWindow(title)
	myWindow.Resize(fyne.NewSize(400, 300))
	myWindow.CenterOnScreen()

	loginForm := widget.NewForm(
		widget.NewFormItem("用户名", widget.NewEntry()),
		widget.NewFormItem("密码", widget.NewPasswordEntry()),
	)
	loginForm.SubmitText = "登录"
	loginForm.OnSubmit = func() {
		log.Print("已登录")
	}
	loginForm.CancelText = "取消"
	loginForm.OnCancel = func() {
		log.Print("已取消")
		myWindow.Close()
	}

	content := container.New(layout.NewMaxLayout(), loginForm)
	myWindow.SetContent(content)

	myWindow.ShowAndRun()
}
