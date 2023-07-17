package ui

import (
	"testing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/widget"
	"github.com/mattn/go-mastodon"
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
	assert.Equal(t, 3, lc.LeftList.Length())
	assert.Equal(t, 2, lc.RightList.Length())
	for i, v := range leftTags {
		got := getListItem(lc.LeftList, i).(*widget.Label)
		assert.Equal(t, v.Name, got.Text, "Expecting left list item %d to be %s, but got %s", i, v.Name, got.Text)
	}
	for i, v := range rightTags {
		got := getListItem(lc.RightList, i).(*widget.Label)
		assert.Equal(t, v.Name, got.Text, "Expecting right list item %d to be %s, but got %s", i, v.Name, got.Text)
	}
}

func TestListChoices_ListHeadersAreCorrect(t *testing.T) {
	a := test.NewApp()
	w := a.NewWindow("")
	lc := NewListChoices()
	w.SetContent(lc)
	w.Resize(fyne.Size{Width: 400, Height: 400})

	assert.True(t, lc.LeftLabel.TextStyle.Bold,
		"Expecting left label style to be Bold")
	assert.Equal(t,
		fyne.TextAlignCenter, lc.LeftLabel.Alignment,
		"Expecting left label alignment to be %d, but got %d",
		fyne.TextAlignCenter, lc.LeftLabel.Alignment)
	assert.True(t, lc.RightLabel.TextStyle.Bold,
		"Expecting right label style to be Bold")
	assert.Equal(t,
		fyne.TextAlignCenter, lc.LeftLabel.Alignment,
		"Expecting right label alignment to be %d, but got %d",
		fyne.TextAlignCenter, lc.RightLabel.Alignment)
}

func TestListChoices_TagMovingButtonTapsMoveTags(t *testing.T) {
	numFollowedTags := 2
	allFollowedTags := createTags("Tag", numFollowedTags)
	a := test.NewApp()
	w := a.NewWindow("")
	lc := NewListChoices()
	lc.SetLeftItems(allFollowedTags)
	ui := container.NewMax(lc)
	w.SetContent(ui)
	w.Resize(fyne.Size{Width: 400, Height: 400})
	assert.Equal(t, len(allFollowedTags), lc.LeftList.Length(), "Left List length")
	assert.True(t, lc.MoveRightButton.Disabled(), "Move right button diabled")
	assert.True(t, lc.MoveLeftButton.Disabled(), "Move left button disabled")

	// Move all tags from left list to right list
	for numRemove := 1; numRemove <= len(allFollowedTags); numRemove++ {
		lc.LeftList.Select(0)
		assert.False(t, lc.MoveRightButton.Disabled(), "Move right button should be enabled when left item selected")
		test.Tap(lc.MoveRightButton)
		assert.Equal(t, numRemove, lc.RightList.Length())
		assert.Equal(t, len(allFollowedTags)-numRemove, lc.LeftList.Length())
	}

	// // Move all tags back to left list
	// for numRemove := 1; numRemove <= len(allFollowedTags); numRemove++ {
	// 	lc.RightList.Select(0)
	// 	assert.False(t, lc.MoveLeftButton.Disabled(), "Move left button should be enabled when right item selected")
	// 	test.Tap(lc.MoveLeftButton)
	// 	assert.Equal(t, numRemove, lc.LeftList.Length())
	// 	assert.Equal(t, len(allFollowedTags)-numRemove, lc.RightList.Length())
	// }

	expectedLeftLen := 0
	expectedRightLen := len(allFollowedTags)
	assert.Equal(t, expectedLeftLen, lc.LeftList.Length())
	assert.Equal(t, expectedRightLen, lc.RightList.Length())
	// Move all tags back to left list
	for numRemove := 0; numRemove < len(allFollowedTags); numRemove++ {
		t.Logf("Item to remove %d", numRemove)
		selectID := len(allFollowedTags) - numRemove - 1
		t.Logf("Right selecting %d", selectID)
		lc.RightList.Select(selectID)
		assert.False(t, lc.MoveLeftButton.Disabled(), "Move left button should be enabled when right item selected")
		test.Tap(lc.MoveLeftButton)
		t.Log("Moved left")
		t.Logf("Left length: %d Right Length %d", lc.LeftList.Length(), lc.RightList.Length())
		expectedLeftLen++
		expectedRightLen--
		assert.Equal(t, expectedLeftLen, lc.LeftList.Length())
		assert.Equal(t, expectedRightLen, lc.RightList.Length())

	}
}

func TestListChoices_TagMovingButtonsMoveTagsSelectingFirstListItems(t *testing.T) {
	tags := []*mastodon.FollowedTag{
		{Name: "Tag1"},
		{Name: "Tag2"},
		{Name: "Tag3"},
	}
	tagsCopy := make([]*mastodon.FollowedTag, len(tags))
	copy(tagsCopy, tags)
	a := test.NewApp()
	w := a.NewWindow("")
	lc := NewListChoices()
	lc.SetLeftItems(tagsCopy)
	ui := container.NewMax(lc)
	w.SetContent(ui)
	w.Resize(fyne.Size{Width: 400, Height: 400})
	assert.Equal(t, len(tags), lc.LeftList.Length(), "Left List length")
	assert.True(t, lc.MoveRightButton.Disabled(), "Move right button diabled")
	assert.True(t, lc.MoveLeftButton.Disabled(), "Move left button disabled")

	// Move all tags from left list to right list
	for i := 0; i < len(tags); i++ {
		lc.LeftList.Select(0)
		assert.False(t, lc.MoveRightButton.Disabled(), "Move right button should be enabled when left item selected")
		test.Tap(lc.MoveRightButton)
	}
	assert.Equal(t, 0, lc.LeftList.Length(), "left list should not have any items")
	assert.Equal(t, len(tags), lc.RightList.Length(), "right list should have all items")
	// Right list will be in reverse order
	for i := 0; i < len(tags); i++ {
		l := getListItem(lc.RightList, len(tags)-i-1).(*widget.Label)
		assert.Equal(t, tags[i].Name, l.Text)
	}

	// Move all tags from right list to left list
	for i := 0; i < len(tags); i++ {
		lc.RightList.Select(0)
		assert.False(t, lc.MoveLeftButton.Disabled(), "Move right button should be enabled when left item selected")
		test.Tap(lc.MoveLeftButton)
	}
	assert.Equal(t, 0, lc.RightList.Length(), "right list should not have any items")
	assert.Equal(t, len(tags), lc.LeftList.Length(), "left list should have all items")
	for i := 0; i < len(tags); i++ {
		l := getListItem(lc.LeftList, i).(*widget.Label)
		assert.Equal(t, tags[i].Name, l.Text)
	}
}
func TestListChoices_TagMovingButtonsMoveTagsSelectingLastListItems(t *testing.T) {
	tags := []*mastodon.FollowedTag{
		{Name: "Tag1"},
		{Name: "Tag2"},
		{Name: "Tag3"},
	}
	tagsCopy := make([]*mastodon.FollowedTag, len(tags))
	copy(tagsCopy, tags)
	a := test.NewApp()
	w := a.NewWindow("")
	lc := NewListChoices()
	lc.SetLeftItems(tagsCopy)
	ui := container.NewMax(lc)
	w.SetContent(ui)
	w.Resize(fyne.Size{Width: 400, Height: 400})
	assert.Equal(t, len(tags), lc.LeftList.Length(), "Left List length")
	assert.True(t, lc.MoveRightButton.Disabled(), "Move right button diabled")
	assert.True(t, lc.MoveLeftButton.Disabled(), "Move left button disabled")

	// Move all tags from left list to right list, selecting the bottommost element
	for i := len(tags) - 1; i >= 0; i-- {
		lc.LeftList.Select(i)
		assert.False(t, lc.MoveRightButton.Disabled(), "Move right button should be enabled when left item selected")
		test.Tap(lc.MoveRightButton)
	}
	assert.Equal(t, 0, lc.LeftList.Length(), "left list should not have any items")
	assert.Equal(t, len(tags), lc.RightList.Length(), "right list should have all items")
	for i := 0; i < len(tags); i++ {
		l := getListItem(lc.RightList, i).(*widget.Label)
		assert.Equal(t, tags[i].Name, l.Text)
	}

	// Move all tags from right list to left list
	for i := len(tags) - 1; i >= 0; i-- {
		lc.RightList.Select(i)
		assert.False(t, lc.MoveLeftButton.Disabled(), "Move left button should be enabled when left item selected")
		test.Tap(lc.MoveLeftButton)
	}
	assert.Equal(t, 0, lc.RightList.Length(), "right list should not have any items")
	assert.Equal(t, len(tags), lc.LeftList.Length(), "left list should have all items")
	for i := 0; i < len(tags); i++ {
		l := getListItem(lc.LeftList, i).(*widget.Label)
		assert.Equal(t, tags[i].Name, l.Text)
	}
}
