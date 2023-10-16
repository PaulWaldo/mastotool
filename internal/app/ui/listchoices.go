package ui

import (
	"log"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/mattn/go-mastodon"
	"golang.org/x/exp/slices"
)

var _ fyne.Widget = &ListChoices{}

type ListChoices struct {
	widget.BaseWidget
	LeftItems, RightItems []*mastodon.FollowedTag
	LeftLabel, RightLabel *widget.Label

	LeftList, RightList               *widget.List
	MoveLeftButton, MoveRightButton   *widget.Button
	LeftSelectionId, RightSelectionId widget.ListItemID
	maxItemLength                     int
	maxItem                           string
	rItemsMutex, lItemsMutex          sync.RWMutex
}

func NewListChoices() *ListChoices {
	lc := &ListChoices{}
	lc.LeftItems = []*mastodon.FollowedTag{}
	lc.RightItems = []*mastodon.FollowedTag{}
	lc.LeftList = widget.NewList(
		func() int {
			lc.lItemsMutex.RLock()
			n := len(lc.LeftItems)
			lc.lItemsMutex.RUnlock()
			return n
		},
		func() fyne.CanvasObject {
			if lc.maxItem == "" {
				lc.maxItem = "XXX"
			}
			return widget.NewLabel(lc.maxItem)
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(lc.LeftItems[i].Name)
		},
	)
	lc.RightList = widget.NewList(
		func() int { return len(lc.RightItems) },
		func() fyne.CanvasObject {
			if lc.maxItem == "" {
				lc.maxItem = "XXX"
			}
			return widget.NewLabel(lc.maxItem)
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(lc.RightItems[i].Name)
		},
	)

	lc.LeftLabel = widget.NewLabel("Left")
	lc.LeftLabel.TextStyle = fyne.TextStyle{Bold: true}
	lc.LeftLabel.Alignment = fyne.TextAlignCenter

	lc.RightLabel = widget.NewLabel("Right")
	lc.RightLabel.TextStyle = fyne.TextStyle{Bold: true}
	lc.RightLabel.Alignment = fyne.TextAlignCenter

	lc.MoveRightButton = widget.NewButtonWithIcon("Remove", theme.NavigateNextIcon(), func() {})
	lc.MoveLeftButton = widget.NewButtonWithIcon("Keep", theme.NavigateBackIcon(), func() {})
	lc.MoveRightButton.Disable()
	lc.MoveLeftButton.Disable()
	lc.MoveRightButton.OnTapped = func() {
		if lc.LeftSelectionId > lc.LeftList.Length()-1 {
			log.Printf("Refusing to attempt to right move item #%d when list length is only %d",
				lc.LeftSelectionId, lc.LeftList.Length())
			return
		}
		// Add tag to right list
		lc.RightItems = slices.Insert(lc.RightItems, 0, lc.LeftItems[lc.LeftSelectionId])

		// Remove tag from left list
		lc.LeftItems = slices.Delete(lc.LeftItems, lc.LeftSelectionId, lc.LeftSelectionId+1)

		lc.RightList.UnselectAll()
		lc.Refresh()
		if lc.LeftSelectionId > lc.LeftList.Length()-1 {
			lc.LeftList.Select(lc.LeftList.Length() - 1)
		}
		if lc.LeftList.Length() == 0 {
			lc.RightList.Select(0)
		}
	}
	lc.MoveLeftButton.OnTapped = func() {
		if lc.RightSelectionId > lc.RightList.Length()-1 {
			log.Printf("Refusing to attempt to left move item #%d when list length is only %d",
				lc.RightSelectionId, lc.RightList.Length())
			return
		}
		// Add tag to left list
		lc.LeftItems = slices.Insert(lc.LeftItems, 0, lc.RightItems[lc.RightSelectionId])

		// Remove tag from right list
		lc.RightItems = slices.Delete(lc.RightItems, lc.RightSelectionId, lc.RightSelectionId+1)

		lc.LeftList.UnselectAll()
		lc.Refresh()
		if lc.RightSelectionId > lc.RightList.Length()-1 {
			lc.RightList.Select(lc.RightList.Length() - 1)
		}
		if lc.RightList.Length() == 0 {
			lc.LeftList.Select(0)
		}
	}

	lc.LeftList.OnSelected = func(id widget.ListItemID) {
		lc.LeftSelectionId = id
		lc.MoveRightButton.Enable()
		lc.MoveLeftButton.Disable()
	}
	lc.RightList.OnSelected = func(id widget.ListItemID) {
		lc.RightSelectionId = id
		lc.MoveLeftButton.Enable()
		lc.MoveRightButton.Disable()
	}
	lc.ExtendBaseWidget(lc)
	return lc
}

func (lc *ListChoices) SetLeftItems(t []*mastodon.FollowedTag) {
	lc.lItemsMutex.Lock()
	lc.LeftItems = t
	for _, m := range t {
		l := fyne.MeasureText(m.Name, fyne.CurrentApp().Settings().Theme().Size(theme.SizeNameText), fyne.TextStyle{})
		if int(l.Width) > lc.maxItemLength {
			lc.maxItemLength = int(l.Width)
			lc.maxItem = m.Name
		}
	}
	lc.lItemsMutex.Unlock()
	lc.LeftList.Refresh()
}

func (lc *ListChoices) SetRightItems(t []*mastodon.FollowedTag) {
	lc.rItemsMutex.Lock()
	lc.RightItems = t
	lc.rItemsMutex.Unlock()
	lc.RightList.Refresh()
}

func (lc *ListChoices) CreateRenderer() fyne.WidgetRenderer {
	buttons := container.NewBorder(
		container.NewVBox(lc.MoveRightButton, lc.MoveLeftButton),
		nil, nil, nil)
	keepBox := container.NewBorder(lc.LeftLabel, nil, nil, nil, lc.LeftList)
	// Create filler for the buttons to keep them from being at the very top
	fill := widget.NewLabel("")
	buttonBox := container.NewBorder(fill, nil, nil, nil, buttons)
	removeBox := container.NewBorder(lc.RightLabel, nil, nil, nil, lc.RightList)
	container := container.NewHBox(keepBox, buttonBox, removeBox)

	return widget.NewSimpleRenderer(container)
}
