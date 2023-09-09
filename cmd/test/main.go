package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"github.com/PaulWaldo/mastotool/internal/app/ui"
	"github.com/mattn/go-mastodon"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	lc := ui.NewListChoices()
	lc.SetLeftItems([]*mastodon.FollowedTag{{Name: "aaa"}, {Name: "bbb"}})
	lc.SetRightItems([]*mastodon.FollowedTag{{Name: "AAAAA"}})
	w.SetContent(container.NewStack(lc))

	w.Resize(fyne.Size{Width: 400, Height: 400})
	w.ShowAndRun()
}
