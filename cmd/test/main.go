package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"github.com/PaulWaldo/mastotool/reorg/ui"
	"github.com/mattn/go-mastodon"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	// ftui := ui.FollowedTagsUI{
	// 	KeepTags:   []*mastodon.FollowedTag{{Name: "aaa"}, {Name: "bbb"}},
	// 	RemoveTags: []*mastodon.FollowedTag{{Name: "AAAAA"}},
	// }
	// w.SetContent(container.NewMax(ftui.MakeFollowedTagsUI()))
	lc := ui.NewListChoices()
	lc.SetLeftItems([]*mastodon.FollowedTag{{Name: "aaa"}, {Name: "bbb"}})
	lc.SetRightItems([]*mastodon.FollowedTag{{Name: "AAAAA"}})
	w.SetContent(container.NewMax(lc))

	w.Resize(fyne.Size{Width: 400, Height: 400})
	w.ShowAndRun()
}
