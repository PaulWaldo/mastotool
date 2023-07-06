package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/mattn/go-mastodon"
)

type FollowedTagsUI struct {
	KeepTags    []*mastodon.FollowedTag
	RemoveTags  []*mastodon.FollowedTag
	ListChoices *ListChoices
	// Container   *fyne.Container
}

func (ft *FollowedTagsUI) MakeFollowedTagsUI() fyne.CanvasObject {
	ft.ListChoices = NewListChoices()

	ft.ListChoices.LeftLabel.Text = "To Keep"
	ft.ListChoices.LeftLabel.TextStyle = fyne.TextStyle{Bold: true}
	ft.ListChoices.LeftLabel.Alignment = fyne.TextAlignCenter

	ft.ListChoices.RightLabel.Text = "To Remove"
	ft.ListChoices.RightLabel.TextStyle = fyne.TextStyle{Bold: true}
	ft.ListChoices.RightLabel.Alignment = fyne.TextAlignCenter

	ft.ListChoices.LeftItems = ft.KeepTags
	ft.ListChoices.RightItems = ft.RemoveTags
	return container.NewMax(ft.ListChoices)
}

func (ui *FollowedTagsUI) SetFollowedTags(t []*mastodon.FollowedTag) {
	// ui.keepTags = t
	// ui.removeTags = []*mastodon.FollowedTag{}
	ui.ListChoices.SetLeftItems(t)
	ui.ListChoices.SetRightItems([]*mastodon.FollowedTag{})
}
