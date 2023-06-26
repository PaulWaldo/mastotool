package listchoices

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/mattn/go-mastodon"
)

type ListChoices struct {
	widget.BaseWidget
	LeftItems, RightItems []*mastodon.FollowedTag
	LeftLabel, RightLabel string

	leftList, rightList               *widget.List
	moveLeftButton, moveRightButton   *widget.Button
	leftSelectionId, rightSelectionId widget.ListItemID
}

func NewListChoices(
	leftItems, rightItems []*mastodon.FollowedTag,
	// createItem func() fyne.CanvasObject,
	// updateItem func(widget.ListItemID, fyne.CanvasObject),
) *ListChoices {
	lc := &ListChoices{
		leftList: widget.NewList(
			func() int { return len(leftItems) },
			func() fyne.CanvasObject {
				return widget.NewLabel("XXXXXXXXXXXXXXX")
			},
			func(i widget.ListItemID, o fyne.CanvasObject) {
				o.(*widget.Label).SetText(leftItems[i].Name)
			},
		),
		rightList: widget.NewList(
			func() int { return len(rightItems) },
			func() fyne.CanvasObject {
				return widget.NewLabel("XXXXXXXXXXXXXXX")
			},
			func(i widget.ListItemID, o fyne.CanvasObject) {
				o.(*widget.Label).SetText(rightItems[i].Name)
			},
		),
		moveRightButton: widget.NewButtonWithIcon("Remove", theme.NavigateNextIcon(), func() {

		}),
		moveLeftButton: widget.NewButtonWithIcon("Keep", theme.NavigateBackIcon(), func() {

		}),
	}
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

func (lc *ListChoices) CreateRenderer() fyne.WidgetRenderer {
	lcr := listChoicesRenderer{
		listChoices: lc,
		container:   container.NewHBox(lc.leftList, lc.moveLeftButton, lc.moveRightButton, lc.rightList),
	}
	return lcr
}

var _ fyne.WidgetRenderer = listChoicesRenderer{}

type listChoicesRenderer struct {
	listChoices *ListChoices
	container   *fyne.Container
}

func (lcr listChoicesRenderer) Destroy() {}

func (lcr listChoicesRenderer) Layout(s fyne.Size) {}

func (lcr listChoicesRenderer) MinSize() fyne.Size { return fyne.Size{} }

func (lcr listChoicesRenderer) Objects() []fyne.CanvasObject {
	// lc := lcr.listChoices
	return []fyne.CanvasObject{
		lcr.container,
	}
}

func (lcr listChoicesRenderer) Refresh() {
	lcr.container.Refresh()
}
