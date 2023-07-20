package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"github.com/mattn/go-mastodon"
)

const AppID = "com.github.PaulWaldo.mastotool"

type myApp struct {
	prefs                Preferences
	app                  fyne.App
	window               fyne.Window
	keepTags, removeTags []*mastodon.FollowedTag
	listChoices          *ListChoices
	unfollowButton       *widget.Button
	refreshButton        *widget.Button
}

func Run() {
	ma := myApp{}
	ma.app = app.NewWithID(AppID)
	ma.prefs = NewPreferences(ma.app)
	ma.window = ma.app.NewWindow("MastoTool")
	ma.window.SetMainMenu(fyne.NewMainMenu(
		fyne.NewMenu("Server",
			fyne.NewMenuItem("Log In", ma.authenticate),
			fyne.NewMenuItem("Log Out", ma.forgetCredentials),
		)),
	)
	ma.window.SetContent(ma.MakeFollowedTagsUI())
	ma.window.Resize(fyne.Size{Width: 400, Height: 400})
	go ma.getFollowedTags()
	ma.window.ShowAndRun()
}
