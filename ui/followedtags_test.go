package ui

import (
	"fmt"
	"testing"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/widget"
	"github.com/mattn/go-mastodon"
	"github.com/stretchr/testify/assert"
)

func getListItem(l *widget.List, index int) fyne.CanvasObject {
	l.ScrollTo(index)
	listRenderer := test.WidgetRenderer(l)
	items := listRenderer.Objects()
	scrollRenderer := test.WidgetRenderer(items[0].(fyne.Widget))
	scrollObjs := scrollRenderer.Objects()
	listContainer := scrollObjs[0].(*fyne.Container)
	listCanvas := listContainer.Objects
	listItem := listCanvas[index].(fyne.Widget)
	listItemRenderer := test.WidgetRenderer(listItem)
	listItemCanvas := listItemRenderer.Objects()
	return listItemCanvas[1].(*widget.Label)
}

func TestFollowedTagsUI_MakeFollowedTagsUI(t *testing.T) {
	type fields struct {
		followedTags []*mastodon.FollowedTag
	}
	numFollowedTags := 3
	allFollowedTags := make([]*mastodon.FollowedTag, numFollowedTags)
	for i := 0; i < numFollowedTags; i++ {
		allFollowedTags[i] = &mastodon.FollowedTag{
			Name: fmt.Sprintf("Tag%d", i),
			History: []mastodon.FollowedTagHistory{
				{
					Day:      mastodon.UnixTimeString{Time: time.Now()},
					Accounts: 10 * i,
					Uses:     100 * i,
				},
				{
					Day:      mastodon.UnixTimeString{Time: time.Now().AddDate(0, 0, -1)},
					Accounts: 20 * i,
					Uses:     200 * i,
				},
				{
					Day:      mastodon.UnixTimeString{Time: time.Now().AddDate(0, 0, -2)},
					Accounts: 30 * i,
					Uses:     300 * i,
				},
			},
		}
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Initial tags in keep list, none in remove list",
			fields: fields{
				followedTags: allFollowedTags,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := test.NewApp()
			w := a.NewWindow("")
			ui := NewFollowedTagsUI(tt.fields.followedTags)
			w.SetContent(ui.MakeFollowedTagsUI())
			w.Resize(fyne.Size{Width: 400, Height: 400})
			assert.Equal(t, len(allFollowedTags), ui.keepListWidget.Length(), "Expecting keep list widget to have %d items, got %d", len(allFollowedTags), ui.keepListWidget.Length())
			assert.Equal(t, 0, ui.removeListWidget.Length())

			for i, v := range allFollowedTags {
				got := getListItem(ui.keepListWidget, i).(*widget.Label)
				assert.Equal(t, v.Name, got.Text, "Expecting keep list item %d to be %s, but got %s", i, v.Name, got.Text)
			}
		})
	}
}
