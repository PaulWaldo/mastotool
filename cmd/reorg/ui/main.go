package ui

import (
	"context"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
	"github.com/mattn/go-mastodon"
)

type myApp struct {
	prefs        Preferences
	app          fyne.App
	window       fyne.Window
	followedTags *followedTags
}

func Run() {
	ma := myApp{}
	ma.app = app.NewWithID("com.github.PaulWaldo.mastotool")
	ma.prefs = NewPreferences(ma.app)
	ma.window = ma.app.NewWindow("MastoTool")
	ma.window.SetMainMenu(fyne.NewMainMenu(
		fyne.NewMenu("File", fyne.NewMenuItem("Authenticate", ma.authenticate))),
	)
	ma.window.SetContent(ma.followedTags.MakeFollowedTagsUI())
	ma.window.Resize(fyne.Size{Width: 400, Height: 400})
	go ma.getFollowedTags()
	ma.window.ShowAndRun()
}

// getFollowedTags gets the list of followed tags and populates the keepTags and removeTags based on this
func (ma *myApp) getFollowedTags() {
	c := NewClientFromPrefs(ma.prefs)
	var err error
	ma.followedTags.keepTags, err = c.GetFollowedTags(context.Background(), nil)
	if err != nil {
		ma.followedTags.keepTags = []*mastodon.FollowedTag{}
		dialog.ShowError(err, ma.window)
	}
	ma.followedTags.removeTags = []*mastodon.FollowedTag{}
}
