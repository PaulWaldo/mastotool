package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type myApp struct {
	prefs  *preferences
	app    fyne.App
	window fyne.Window
}

func Run() {
	myApp := myApp{}
	myApp.app = app.NewWithID("com.github.PaulWaldo.mastotool")
	myApp.prefs = NewPreferences(myApp.app)
	myApp.window = myApp.app.NewWindow("MastoTool")
	myApp.window.SetMainMenu(fyne.NewMainMenu(
		fyne.NewMenu("File", fyne.NewMenuItem("Authenticate", myApp.authenticate))),
	)
	myApp.window.Resize(fyne.Size{Width: 400, Height: 400})
	myApp.window.ShowAndRun()
}
