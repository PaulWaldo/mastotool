package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type myApp struct {
	prefs preferences
	app   fyne.App
}

func main() {
	myApp := myApp{}
	myApp.app = app.NewWithID("com.github.PaulWaldo.mastotool")
	w := myApp.app.NewWindow("MastoTool")
	w.Resize(fyne.Size{Width: 400, Height: 400})
	w.ShowAndRun()

}
