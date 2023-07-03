package ui

import (
	"testing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/widget"
	"github.com/mattn/go-mastodon"
	"github.com/stretchr/testify/assert"
)

func TestSetFollowedTags_PopulatesList(t *testing.T) {
	type fields struct {
		numFollowedTags int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "with 3 tags",
			fields: fields{
				numFollowedTags: 3,
			},
		},
	}
	for i := range tests {
		tt := &tests[i]
		t.Run(tt.name, func(t *testing.T) {
			a := test.NewApp()
			w := a.NewWindow("")
			keepTags := createTags("Tag", tt.fields.numFollowedTags)
			removeTags := []*mastodon.FollowedTag{}
			ftui := FollowedTagsUI{KeepTags: keepTags, RemoveTags: removeTags}
			container := ftui.MakeFollowedTagsUI()
			allFollowedTags := createTags("Tag", tt.fields.numFollowedTags)
			w.SetContent(container)
			w.Resize(fyne.Size{Width: 400, Height: 400})
			keepList := ftui.ListChoices.leftList
			removeList := ftui.ListChoices.rightList
			assert.Equal(t, tt.fields.numFollowedTags, keepList.Length(), "Expecting keep list widget to have %d items, got %d", len(allFollowedTags), keepList.Length())
			assert.Equal(t, 0, removeList.Length())

			for i, v := range allFollowedTags {
				got := getListItem(keepList, i).(*widget.Label)
				assert.Equal(t, v.Name, got.Text, "Expecting keep list item %d to be %s, but got %s", i, v.Name, got.Text)
			}

			// Initial state of Unfollow button should be disabled
			// assert.True(t, ucontainerunfollowButton.Disabled(), "Initial state of Unfollow button should be disabled")
		})
	}
}

// func TestFollowedTagsUI_TagMovingButtonPressesMoveTags(t *testing.T) {
// 	numFollowedTags := 2
// 	allFollowedTags := getTagList(numFollowedTags)
// 	a := test.NewApp()
// 	w := a.NewWindow("")
// 	ui := NewFollowedTagsUI(allFollowedTags)
// 	w.SetContent(ui.MakeFollowedTagsUI())
// 	w.Resize(fyne.Size{Width: 400, Height: 400})
// 	assert.Equal(t, len(allFollowedTags), ui.keepListWidget.Length())
// 	assert.True(t, ui.removeButton.Disabled())
// 	assert.True(t, ui.keepButton.Disabled())
// 	assert.True(t, ui.unfollowButton.Disabled())

// 	// Move all tags from keep list to remove list
// 	for numRemove := 1; numRemove <= len(allFollowedTags); numRemove++ {
// 		ui.keepListWidget.Select(0)
// 		assert.False(t, ui.removeButton.Disabled())
// 		test.Tap(ui.removeButton)
// 		assert.Equal(t, numRemove, ui.removeListWidget.Length())
// 		assert.Equal(t, len(allFollowedTags)-numRemove, ui.keepListWidget.Length())
// 		assert.False(t, ui.unfollowButton.Disabled())
// 	}

// 	// Move all tags back to keep list
// 	for numRemove := 1; numRemove <= len(allFollowedTags); numRemove++ {
// 		ui.removeListWidget.Select(0)
// 		assert.False(t, ui.unfollowButton.Disabled())
// 		ui.removeListWidget.Select(0)
// 		assert.False(t, ui.keepButton.Disabled())
// 		test.Tap(ui.keepButton)
// 		assert.Equal(t, numRemove, ui.keepListWidget.Length())
// 		assert.Equal(t, len(allFollowedTags)-numRemove, ui.removeListWidget.Length())
// 	}
// 	assert.True(t, ui.unfollowButton.Disabled())
// }
