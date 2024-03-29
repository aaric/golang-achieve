package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"log"
	"net/url"
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

func tidyUp() {
	fmt.Println("Exited")
}

func main2(fun string) {
	myApp := app.New()

	switch fun {
	case "helloApp":
		// 1. hello world
		helloApp(myApp)
	case "clockApp":
		// 2. clock
		clockApp(myApp)
	case "openWindowApp":
		// 3. open window
		openWindowApp(myApp)
	case "canvasApp":
		// 4. canvas
		canvasApp(myApp)
	case "testUiApp":
		// 5. test ui
		testUiApp(myApp)
	case "trayApp":
		// 6. system tray
		trayApp(myApp)
	case "layoutApp":
		// 7. layout | https://developer.fyne.io/explore/layouts
		layoutApp(myApp)
	case "widgetApp":
		// 8. widget | https://developer.fyne.io/api/v2.2/widget/
		widgetApp(myApp)
	case "themeApp":
		// 9. theme
		themeApp(myApp)
	default:
		fmt.Println("not match")
	}

	myApp.Run()
	tidyUp()
}

func helloApp(myApp fyne.App) {
	myWindow := myApp.NewWindow("Hello World")
	myWindow.SetContent(widget.NewLabel("This is some content."))
}

func clockApp(myApp fyne.App) {
	myWindow := myApp.NewWindow("Clock")
	clock := widget.NewLabel("")
	updateTime(clock)
	myWindow.SetContent(clock)
	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()
}

func openWindowApp(myApp fyne.App) {
	myWindow := myApp.NewWindow("New Window")
	myWindow.SetContent(widget.NewButton("Open Window", func() {
		w2 := myApp.NewWindow("Window 2")
		w2.SetContent(widget.NewLabel("This is other window."))
		w2.Show()
	}))
}

func canvasApp(myApp fyne.App) {
	myWindow := myApp.NewWindow("Canvas")
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
}

func testUiApp(myApp fyne.App) {
	myWindow := myApp.NewWindow("Test UI")
	myWindow.SetContent(container.NewVBox(makeUI()))
}

func trayApp(myApp fyne.App) {
	myWindow := myApp.NewWindow("System Tray")

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
	myWindow.SetContent(widget.NewLabel("This is system tray."))
}

func layoutApp(myApp fyne.App) {
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
}

func widgetApp(myApp fyne.App) {
	myWindow := myApp.NewWindow("Widget")
	url, _ := url.Parse("https://www.baidu.com")
	path, _ := fyne.LoadResourceFromPath("./Icon.png")
	myWindow.SetContent(container.NewVScroll(container.NewVBox(
		widget.NewAccordion(widget.NewAccordionItem("Accordion", widget.NewLabel("hello world"))),
		widget.NewButton("Button", func() {
			log.Println("hello", "world")
		}),
		widget.NewCard("Card", "this is a card", widget.NewIcon(theme.AccountIcon())),
		widget.NewCheck("Check", func(b bool) {
			log.Println("checkbox", "clicked")
		}),
		widget.NewEntry(),
		widget.NewPasswordEntry(),
		//widget.NewFileIcon(fyne.URI()),
		widget.NewForm(
			widget.NewFormItem("Username", widget.NewEntry()),
			widget.NewFormItem("Password", widget.NewPasswordEntry()),
		),
		widget.NewHyperlink("www.baidu.com", url),
		widget.NewIcon(path),
		widget.NewProgressBar(),
		widget.NewRadioGroup([]string{"go1", "go2"}, func(selected string) {
			log.Println(selected)
		}),
		widget.NewSelect([]string{"go1", "go2"}, func(selected string) {
			log.Println(selected)
		}),
		widget.NewSeparator(),
		widget.NewSlider(0, 100),
		widget.NewToolbar(
			widget.NewToolbarAction(theme.ContentCutIcon(), func() {
				log.Println("cut")
			}),
			widget.NewToolbarSeparator(),
			widget.NewToolbarAction(theme.ContentCopyIcon(), func() {
				log.Println("copied")
			}),
			widget.NewToolbarAction(theme.ContentPasteIcon(), func() {
				log.Println("pasted")
			}),
			widget.NewToolbarSpacer(),
			widget.NewToolbarAction(theme.HelpIcon(), func() {
				log.Println("helped")
			}),
		),
		widget.NewMultiLineEntry(),
		container.NewAppTabs(
			container.NewTabItem("tab1", widget.NewLabel("tab1 content")),
			container.NewTabItem("tab2", widget.NewLabel("tab2 content")),
		),
	)))
}

func themeApp(myApp fyne.App) {
	myWindow := myApp.NewWindow("Theme")
	myWindow.SetContent(container.NewVSplit(container.NewHSplit(
		widget.NewButton("DarkTheme", func() {
			myApp.Settings().SetTheme(theme.DarkTheme())
		}),
		widget.NewButton("LightTheme", func() {
			myApp.Settings().SetTheme(theme.LightTheme())
		}),
	), widget.NewLabel("This is switch theme demo.")))

	myWindow.Resize(fyne.NewSize(400, 300))
	myWindow.SetMaster()
	//myWindow.ShowAndRun()
	myWindow.Show()
}
