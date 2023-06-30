package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/mattn/go-mastodon"
)

var _ fyne.Widget = &ListChoices{}

type ListChoices struct {
	widget.BaseWidget
	LeftItems, RightItems []*mastodon.FollowedTag
	leftLabel, rightLabel *widget.Label

	leftList, rightList               *widget.List
	moveLeftButton, moveRightButton   *widget.Button
	leftSelectionId, rightSelectionId widget.ListItemID
}

func NewListChoices() *ListChoices {
	lc := &ListChoices{}
	lc.LeftItems = []*mastodon.FollowedTag{}
	lc.RightItems = []*mastodon.FollowedTag{}
	lc.leftList = widget.NewList(
		func() int {
			return len(lc.LeftItems)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("XXXXXXXXXXXXXXX")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(lc.LeftItems[i].Name)
		},
	)
	lc.rightList = widget.NewList(
		func() int { return len(lc.RightItems) },
		func() fyne.CanvasObject {
			return widget.NewLabel("XXXXXXXXXXXXXXX")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(lc.RightItems[i].Name)
		},
	)
	lc.leftLabel = widget.NewLabel("")
	lc.rightLabel = widget.NewLabel("")
	lc.moveRightButton = widget.NewButtonWithIcon("Remove", theme.NavigateNextIcon(), func() {
	})
	lc.moveLeftButton = widget.NewButtonWithIcon("Keep", theme.NavigateBackIcon(), func() {
	})
	lc.leftList.OnSelected = func(id widget.ListItemID) {
		lc.leftSelectionId = id
		lc.moveRightButton.Enable()
		lc.moveLeftButton.Disable()
	}
	lc.rightList.OnSelected = func(id widget.ListItemID) {
		lc.rightSelectionId = id
		lc.moveLeftButton.Enable()
		lc.moveRightButton.Disable()
	}
	lc.ExtendBaseWidget(lc)
	return lc
}

func (lc *ListChoices) SetLeftItems(t []*mastodon.FollowedTag) {
	lc.LeftItems = t
	lc.leftList.Refresh()
}

func (lc *ListChoices) SetRightItems(t []*mastodon.FollowedTag) {
	lc.RightItems = t
	lc.rightList.Refresh()
}

func (lc *ListChoices) CreateRenderer() fyne.WidgetRenderer {
	buttons := container.NewVBox(lc.moveLeftButton, lc.moveRightButton)
	// keepBox := container.NewBorder(lc.leftLabel, nil, nil, nil, lc.leftList)
	// buttonBox := container.NewBorder(nil, nil, nil, nil, buttons)
	// removeBox := container.NewBorder(lc.rightLabel, nil, nil, nil, lc.rightList)
	// ui.container = container.NewBorder(nil, ui.unfollowButton, nil, nil,
	// 	container.NewHBox(keepBox, buttonBox, removeBox),
	// )

	lcr := listChoicesRenderer{
		listChoices: lc,
		container:   container.NewHBox(lc.leftList, buttons, lc.rightList),
	}
	return lcr
}

var _ fyne.WidgetRenderer = listChoicesRenderer{}

type listChoicesRenderer struct {
	listChoices *ListChoices
	container   *fyne.Container
}

func (lcr listChoicesRenderer) Destroy() {
}

func (lcr listChoicesRenderer) Layout(s fyne.Size) {
	lcr.container.Resize(s)
}

func (lcr listChoicesRenderer) MinSize() fyne.Size {
	return lcr.container.MinSize()
}

func (lcr listChoicesRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{
		lcr.container,
	}
}

func (lcr listChoicesRenderer) Refresh() {
	lcr.container.Refresh()
}
