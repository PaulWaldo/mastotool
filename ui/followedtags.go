package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/mattn/go-mastodon"
)

type FollowedTagsUI struct {
	KeepTags     []*mastodon.FollowedTag
	RemoveTags   []*mastodon.FollowedTag
	keepButton   *widget.Button
	removeButton *widget.Button
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
	ui.keepButton = widget.NewButtonWithIcon("Keep", theme.NavigateBackIcon(), func() {})
	ui.removeButton = widget.NewButtonWithIcon("Remove", theme.NavigateNextIcon(), func() {})
	ui.keepButton.Disable()
	ui.removeButton.Disable()
	buttons := container.NewVBox(ui.removeButton, ui.keepButton)

	keepList := widget.NewList(
		func() int {
			return len(ui.KeepTags)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("XXXXXXXXXXXXXXX")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(ui.KeepTags[i].Name)
		},
	)
	keepList.OnSelected = func(id widget.ListItemID) {
		ui.keepButton.Disable()
		ui.removeButton.Enable()
	}

	keepLabel := widget.NewLabel("Tags to keep")
	keepLabel.TextStyle = fyne.TextStyle{Bold: true}

	removeList := widget.NewList(
		func() int {
			return len(ui.RemoveTags)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(ui.RemoveTags[i].Name)
		},
	)
	removeList.OnSelected = func(id widget.ListItemID) {
		ui.keepButton.Enable()
		ui.removeButton.Disable()
	}

	removeLabel := widget.NewLabel("Tags to remove")
	removeLabel.TextStyle = fyne.TextStyle{Bold: true}

	keepBox := container.NewBorder(keepLabel, nil, nil, nil, keepList)
	buttonBox := container.NewBorder(nil, nil, nil, nil, buttons)
	removeBox := container.NewBorder(removeLabel, nil, nil, nil, removeList)
	ui.container = container.NewHBox(keepBox, buttonBox, removeBox)
	return ui.container
}
