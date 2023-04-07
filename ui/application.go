package ui

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func Run() {
	a := app.NewWithID("com.github.PaulWaldo.mastotool")
	server := binding.BindPreferenceString("MastodonServer", a.Preferences())
	val, _ := server.Get()
	fmt.Printf("Server is %s\n", val)
	w := a.NewWindow("MastoTool")
	err := server.Set("i am the server name")
	if err != nil {
		log.Fatal(err)
	}
	w.SetMainMenu(fyne.NewMainMenu(
		fyne.NewMenu("File",
			fyne.NewMenuItem("Authenticate", func() {
				dialog.ShowForm("Authenticate to your server", "Authenticate", "Cancel",
					[]*widget.FormItem{widget.NewFormItem("URL of your server", widget.NewEntry())},
					func(b bool) {

					},
					w)
			}),
		),
	))
	w.Resize(fyne.Size{Width: 400, Height: 400})
	w.ShowAndRun()
}
