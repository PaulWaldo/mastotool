package ui

import (
	"log"

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
	container                         *fyne.Container
}

func NewListChoices() *ListChoices {
	lc := &ListChoices{}
	lc.LeftItems = []*mastodon.FollowedTag{}
	lc.RightItems = []*mastodon.FollowedTag{}
	lc.LeftList = widget.NewList(
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
	lc.RightList = widget.NewList(
		func() int { return len(lc.RightItems) },
		func() fyne.CanvasObject {
			return widget.NewLabel("XXXXXXXXXXXXXXX")
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
		// Add tag to right list
		lc.RightItems = slices.Insert(lc.RightItems, 0, lc.LeftItems[lc.LeftSelectionId])

		// Remove tag from left list
		lc.LeftItems = slices.Delete(lc.LeftItems, lc.LeftSelectionId, lc.LeftSelectionId+1)

		lc.container.Refresh()
	}
	lc.MoveLeftButton.OnTapped = func() {
		log.Println("Move left button tapped: Before")
		log.Println("LeftItems:")
		for _, t := range lc.LeftItems {
			log.Printf("\t%s\n", t.Name)
		}
		log.Println("RightItems:")
		for _, t := range lc.RightItems {
			log.Printf("\t%s\n", t.Name)
		}
		log.Printf("Left SelectID: %d, Right Select ID: %d\n", lc.LeftSelectionId, lc.RightSelectionId)
		// Add tag to left list
		lc.LeftItems = slices.Insert(lc.LeftItems, 0, lc.RightItems[lc.RightSelectionId])

		// Remove tag from right list
		lc.RightItems = slices.Delete(lc.RightItems, lc.RightSelectionId, lc.RightSelectionId+1)

		log.Println("After move:")
		log.Println("LeftItems:")
		for _, t := range lc.LeftItems {
			log.Printf("\t%s\n", t.Name)
		}
		log.Println("RightItems:")
		for _, t := range lc.RightItems {
			log.Printf("\t%s\n", t.Name)
		}
		log.Println()
		lc.container.Refresh()
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
	lc.LeftItems = t
	lc.LeftList.Refresh()
}

func (lc *ListChoices) SetRightItems(t []*mastodon.FollowedTag) {
	lc.RightItems = t
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
	lc.container = container.NewHBox(keepBox, buttonBox, removeBox)

	lcr := listChoicesRenderer{
		listChoices: lc,
		container:   lc.container,
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
