package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/mattn/go-mastodon"
)

type FollowedTagsUI struct {
	KeepTags   []*mastodon.FollowedTag
	RemoveTags []*mastodon.FollowedTag
	// KeepTagsBoundList   binding.UntypedList
	// RemoveTagsBoundList binding.UntypedList
	keepButton        *widget.Button
	removeButton      *widget.Button
	unfollowButton    *widget.Button
	container         *fyne.Container
	keepSelectionId   *widget.ListItemID
	removeSelectionId *widget.ListItemID
	keepListWidget    *widget.List
	removeListWidget  *widget.List
}

func NewFollowedTagsUI(ft []*mastodon.FollowedTag) FollowedTagsUI {
	return FollowedTagsUI{
		KeepTags:   ft,
		RemoveTags: []*mastodon.FollowedTag{},
		// KeepTagsBoundList:   binding.NewUntypedList(),
		// RemoveTagsBoundList: binding.NewUntypedList(),
	}
}

func (ui *FollowedTagsUI) SetFollowedTags(ft []*mastodon.FollowedTag) {
	ui.KeepTags = ft
	ui.RemoveTags = []*mastodon.FollowedTag{}
	ui.keepListWidget.Refresh()
	ui.removeListWidget.Refresh()
}

func (ui *FollowedTagsUI) MakeFollowedTagsUI() *fyne.Container {
	ui.keepButton = widget.NewButtonWithIcon("Keep", theme.NavigateBackIcon(), func() {})
	ui.removeButton = widget.NewButtonWithIcon("Remove", theme.NavigateNextIcon(), func() {})
	ui.unfollowButton = widget.NewButtonWithIcon("Unfollow", theme.DeleteIcon(), func() {})
	ui.keepButton.Disable()
	ui.removeButton.Disable()
	ui.unfollowButton.Disable()
	buttons := container.NewBorder(container.NewVBox(ui.removeButton, ui.keepButton), nil, nil, nil)

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
	ui.keepListWidget = keepList

	keepList.OnSelected = func(id widget.ListItemID) {
		ui.removeSelectionId = &id
		ui.keepButton.Disable()
		ui.removeButton.Enable()
	}

	keepLabel := widget.NewLabel("Tags to keep")
	keepLabel.TextStyle = fyne.TextStyle{Bold: true}

	ui.removeListWidget = widget.NewList(
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

	ui.removeListWidget.OnSelected = func(id widget.ListItemID) {
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
		ui.KeepTags = append(ui.KeepTags, ui.RemoveTags[*ui.keepSelectionId])

		// Remove tag from Remove list
		copy(ui.RemoveTags[*ui.keepSelectionId:], ui.RemoveTags[*ui.keepSelectionId+1:])
		ui.RemoveTags = ui.RemoveTags[:len(ui.RemoveTags)-1]

		ui.container.Refresh()
	}

	keepBox := container.NewBorder(keepLabel, nil, nil, nil, keepList)
	buttonBox := container.NewBorder(nil, nil, nil, nil, buttons)
	removeBox := container.NewBorder(removeLabel, nil, nil, nil, ui.removeListWidget)
	ui.container = container.NewBorder(nil, ui.unfollowButton, nil, nil,
		container.NewHBox(keepBox, buttonBox, removeBox),
	)
	return ui.container
}

// https://github.com/fyne-io/developer.fyne.io/pull/54/files/d4a55ebe251f1d55b5abb5dc140c0e3d53c31787
// https://github.com/mJehanno/developer.fyne.io/blob/master/tutorial/list-with-data.md
// func NewBoundList(followedTagListBinding binding.UntypedList) *widget.List {
// 	l := widget.NewListWithData(followedTagListBinding,
// 		func() fyne.CanvasObject {
// 			return container.NewVBox(widget.NewLabel(""))
// 		},
// 		func(di binding.DataItem, co fyne.CanvasObject) {
// 			v, _ := di.(binding.Untyped).Get()
// 			s := binding.BindStruct(v.(*mastodon.FollowedTag))
// 			str, err := s.GetItem("Name")
// 			if err != nil {
// 				panic(err)
// 			}
// 			co.(*fyne.Container).Objects[0].(*widget.Label).Bind(str.(binding.String))

// 			// name := binding.NewString()

// 			// _ = name.Set(v.(*mastodon.FollowedTag).Name)
// 			// co.(*fyne.Container).Objects[0].(*widget.Label).Bind(name)
// 		},
// 	)
// 	return l
// }
