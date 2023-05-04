package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/mattn/go-mastodon"
)

type FollowedTagsUI struct {
	KeepTags     []*mastodon.FollowedTag
	RemoveTags   []*mastodon.FollowedTag
	container    *fyne.Container
}

func NewFollowedTagsUI() *FollowedTagsUI {
	return &FollowedTagsUI{
		KeepTags:   []*mastodon.FollowedTag{},
		RemoveTags: []*mastodon.FollowedTag{},
	}
}

func (ui *FollowedTagsUI) SetFollowedTags(ft []*mastodon.FollowedTag) {
	ui.KeepTags = make([]*mastodon.FollowedTag, len(ft))
	copy(ui.KeepTags, ft)
}

func (ui *FollowedTagsUI) MakeFollowedTagsUI() *fyne.Container {
	keepList := widget.NewList(
		func() int {
			return len(ui.KeepTags)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(ui.KeepTags[i].Name)
		})
	keepLabel := widget.NewLabel("Tags to keep")

	removeList := widget.NewList(
		func() int {
			return len(ui.RemoveTags)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(ui.RemoveTags[i].Name)
		})
	removeLabel := widget.NewLabel("Tags to remove")

	keepBox := container.NewBorder(keepLabel, nil, nil, nil, keepList)
	removeBox := container.NewBorder(removeLabel, nil, nil, nil, removeList)
	ui.container = container.NewHBox(keepBox, removeBox)
	return ui.container
}
