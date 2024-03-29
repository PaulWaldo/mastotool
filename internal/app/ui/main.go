package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"github.com/mattn/go-mastodon"
)

const AppID = "com.github.PaulWaldo.mastotool"

type myApp struct {
	prefs                 AppPrefs
	app                   fyne.App
	window                fyne.Window
	keepTags, removeTags  []*mastodon.FollowedTag
	listChoices           *ListChoices
	unfollowButton        *widget.Button
	refreshButton         *widget.Button
	loginMenu, logoutMenu *fyne.MenuItem
	serverText            *canvas.Text
}

func Run() {
	ma := myApp{}
	ma.app = app.NewWithID(AppID)
	ma.prefs = NewPreferences(ma.app)
	ma.window = ma.app.NewWindow("MastoTool")
	ma.loginMenu = fyne.NewMenuItem("Log In", ma.authenticate)
	ma.logoutMenu = fyne.NewMenuItem("Log Out", ma.forgetCredentials)
	ma.window.SetMainMenu(fyne.NewMainMenu(
		fyne.NewMenu("Server", ma.loginMenu, ma.logoutMenu)),
	)
	ma.setAuthMenuStatus()
	ma.window.SetContent(ma.MakeFollowedTagsUI())
	ma.window.Resize(fyne.Size{Width: 1, Height: 400})
	if ma.isLoggedIn() {
		ma.refreshFollowedTags()
	} else {
		ma.authenticate()
	}
	ma.window.ShowAndRun()
}

func (ma *myApp) setAuthMenuStatus() {
	ma.logoutMenu.Disabled = !ma.isLoggedIn()
	ma.loginMenu.Disabled = ma.isLoggedIn()
}
