package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type preferences struct {
	MastodonServer binding.String
	APIKey         binding.String
}

const (
	ServerKey = "MastodonServer"
	APIKeyKey = "APIKey"
)

func Run() {
	a := app.NewWithID("com.github.PaulWaldo.mastotool")
	prefs := preferences{
		MastodonServer: binding.BindPreferenceString("MastodonServer", a.Preferences()),
		APIKey:         binding.BindPreferenceString(APIKeyKey, a.Preferences()),
	}
	val, _ := prefs.MastodonServer.Get()
	fmt.Printf("Server is %s\n", val)
	w := a.NewWindow("MastoTool")
	w.SetMainMenu(fyne.NewMainMenu(
		fyne.NewMenu("File",
			fyne.NewMenuItem("Authenticate", func() {
				serverUrlEntry := widget.NewEntryWithData(prefs.MastodonServer)
				serverUrlEntry.SetPlaceHolder("https://mymastodonserver.com")
				form := dialog.NewForm("Mastodon Server", "Authenticate", "Abort", []*widget.FormItem{
					{Text: "Server", Widget: serverUrlEntry, HintText: "URL of your Mastodon server"},
				}, func(b bool) {
					if b {
						val, _ := prefs.MastodonServer.Get()
						fmt.Printf("Server is %s\n", val)
					}
				}, w)
				// form.Resize(fyne.Size{Width: 300, Height: 300})
				form.Show()
			}),
		),
	))
	w.Resize(fyne.Size{Width: 400, Height: 400})
	w.ShowAndRun()
}
