package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"github.com/PaulWaldo/mastotool/pkg/listchoices"
	"github.com/mattn/go-mastodon"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	tags := []*mastodon.FollowedTag{
		{Name: "aaaaa"},
		{Name: "bbbbbbbb"},
		{Name: "XXX"},
	}
	w.SetContent(container.NewMax(
		listchoices.NewListChoices(tags, []*mastodon.FollowedTag{{Name: "hghfgfhghdf"}}),
		))

	w.ShowAndRun()
}
