package ui

import (
	"testing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/widget"
	"github.com/stretchr/testify/assert"
)

func TestListChoices_DisplaysCorrectItems(t *testing.T) {
	a := test.NewApp()
	w := a.NewWindow("")
	leftTags := createTags("Left", 3)
	rightTags := createTags("Right", 2)
	lc := NewListChoices()
	lc.SetLeftItems(leftTags)
	lc.SetRightItems(rightTags)
	w.SetContent(container.NewMax(lc))
	w.Resize(fyne.Size{Width: 400, Height: 400})
	assert.Equal(t, 3, lc.leftList.Length())
	assert.Equal(t, 2, lc.rightList.Length())
	for i, v := range leftTags {
		got := getListItem(lc.leftList, i).(*widget.Label)
		assert.Equal(t, v.Name, got.Text, "Expecting left list item %d to be %s, but got %s", i, v.Name, got.Text)
	}
	for i, v := range rightTags {
		got := getListItem(lc.rightList, i).(*widget.Label)
		assert.Equal(t, v.Name, got.Text, "Expecting right list item %d to be %s, but got %s", i, v.Name, got.Text)
	}
}

func TestListChoices_ListHeadersAreCorrect(t *testing.T) {
	a := test.NewApp()
	w := a.NewWindow("")
	lc := NewListChoices()
	w.SetContent(lc)
	w.Resize(fyne.Size{Width: 400, Height: 400})

	assert.True(t, lc.leftLabel.TextStyle.Bold,
		"Expecting left label style to be Bold")
	assert.Equal(t,
		fyne.TextAlignCenter, lc.leftLabel.Alignment,
		"Expecting left label alignment to be %d, but got %d",
		fyne.TextAlignCenter, lc.leftLabel.Alignment)
	assert.True(t, lc.rightLabel.TextStyle.Bold,
		"Expecting right label style to be Bold")
	assert.Equal(t,
		fyne.TextAlignCenter, lc.leftLabel.Alignment,
		"Expecting right label alignment to be %d, but got %d",
		fyne.TextAlignCenter, lc.rightLabel.Alignment)
}
