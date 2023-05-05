package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/mattn/go-mastodon"
)

type FollowedTagsUI struct {
	KeepTags          []*mastodon.FollowedTag
	RemoveTags        []*mastodon.FollowedTag
	keepButton        *widget.Button
	removeButton      *widget.Button
	container         *fyne.Container
	keepSelectionId   *widget.ListItemID
	removeSelectionId *widget.ListItemID
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
	buttons := container.NewCenter(container.NewVBox(ui.removeButton, ui.keepButton))

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
		ui.removeSelectionId = &id
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
		ui.keepSelectionId = &id
		ui.keepButton.Enable()
		ui.removeButton.Disable()
	}

	removeLabel := widget.NewLabel("Tags to remove")
	removeLabel.TextStyle = fyne.TextStyle{Bold: true}

	ui.removeButton.OnTapped = func() {
		// Add tag to Remove list
		if ui.removeSelectionId == nil {
			return
		}
		ui.RemoveTags = append(ui.RemoveTags, ui.KeepTags[*ui.removeSelectionId])

		// Remove tag from Keep list
		copy(ui.KeepTags[*ui.removeSelectionId:], ui.KeepTags[*ui.removeSelectionId+1:])
		ui.KeepTags = ui.KeepTags[:len(ui.KeepTags)-1]

		ui.container.Refresh()
	}

	ui.keepButton.OnTapped = func() {
		// Add tag to Keep list
		if ui.keepSelectionId == nil {
			return
		}
		ui.KeepTags = append(ui.KeepTags, ui.KeepTags[*ui.keepSelectionId])

		// Remove tag from Remove list
		copy(ui.RemoveTags[*ui.keepSelectionId:], ui.RemoveTags[*ui.keepSelectionId+1:])
		ui.RemoveTags = ui.RemoveTags[:len(ui.RemoveTags)-1]

		ui.container.Refresh()
	}

	keepBox := container.NewBorder(keepLabel, nil, nil, nil, keepList)
	buttonBox := container.NewBorder(nil, nil, nil, nil, buttons)
	removeBox := container.NewBorder(removeLabel, nil, nil, nil, removeList)
	ui.container = container.NewHBox(keepBox, buttonBox, removeBox)
	return ui.container
}

// https://github.com/fyne-io/developer.fyne.io/pull/54/files/d4a55ebe251f1d55b5abb5dc140c0e3d53c31787
// https://github.com/mJehanno/developer.fyne.io/blob/master/tutorial/list-with-data.md
func a(followedTagListBinding binding.UntypedList) *widget.List {
	l := widget.NewListWithData(followedTagListBinding,
		func() fyne.CanvasObject {
			return container.NewVBox(widget.NewLabel(""))
		},
		func(di binding.DataItem, co fyne.CanvasObject) {
			v, _ := di.(binding.Untyped).Get()
			name := binding.NewString()

			_ = name.Set(v.(mastodon.FollowedTag).Name)
			co.(*fyne.Container).Objects[0].(*widget.Label).Bind(name)
		},
	)
	return l
}
