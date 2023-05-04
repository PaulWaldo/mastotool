package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/mattn/go-mastodon"
)

type FollowedTagsUI struct {
	KeepTags     binding.IntList
	RemoveTags   binding.IntList
	followedTags []*mastodon.FollowedTag
	container    *fyne.Container
}

func NewFollowedTagsUI() *FollowedTagsUI {
	return &FollowedTagsUI{
		KeepTags:   binding.BindIntList(&[]int{}),
		RemoveTags: binding.BindIntList(&[]int{}),
	}
}

func (ui *FollowedTagsUI) SetFollowedTags(ft []*mastodon.FollowedTag) {
	keepIDs := make([]int, len(ft))
	for i := 0; i < len(ft); i++ {
		keepIDs[i] = i
	}
	var err error
	err = ui.RemoveTags.Set([]int{})
	if err != nil {
		fyne.LogError("Setting RemoveTags: ", err)
	}
	err = ui.KeepTags.Set(keepIDs)
	if err != nil {
		fyne.LogError("Setting KeepTags: ", err)
	}
}

func (ui *FollowedTagsUI) MakeFollowedTagsUI() *fyne.Container {
	keepList := widget.NewListWithData(ui.KeepTags,
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			o.(*widget.Label).Bind(binding.IntToString(i.(binding.Int)))
		},
	)
	keepLabel := widget.NewLabel("Tags to keep")
	removeList := widget.NewListWithData(ui.RemoveTags,
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			o.(*widget.Label).Bind(binding.IntToString(i.(binding.Int)))
		},
	)
	removeLabel := widget.NewLabel("Tags to remove")

	keepBox := container.NewBorder(keepLabel, nil, nil, nil, keepList)
	removeBox := container.NewBorder(removeLabel, nil, nil, nil, removeList)
	ui.container = container.NewHBox(keepBox, removeBox)
	return ui.container
}
