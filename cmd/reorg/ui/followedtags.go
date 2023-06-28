package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/mattn/go-mastodon"
)

type followedTags struct {
	keepTags    []*mastodon.FollowedTag
	removeTags  []*mastodon.FollowedTag
	listChoices listChoices
}

func (ft *followedTags) MakeFollowedTagsUI() fyne.CanvasObject {
	ft.keepTags = []*mastodon.FollowedTag{}
	ft.removeTags = []*mastodon.FollowedTag{}
	ft.listChoices = *NewListChoices(ft.keepTags, ft.removeTags)
	return container.NewMax(&ft.listChoices)
}

func (ui *followedTags) SetFollowedTags(t []*mastodon.FollowedTag) {
	ui.keepTags = t
	ui.removeTags = []*mastodon.FollowedTag{}
	ui.listChoices.leftList.Refresh()
	ui.listChoices.rightList.Refresh()
}
