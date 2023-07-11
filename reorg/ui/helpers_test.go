package ui

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/widget"
	"github.com/mattn/go-mastodon"
)

func createTags(prefix string, n int) []*mastodon.FollowedTag {
	tags := make([]*mastodon.FollowedTag, n)
	for i := 0; i < n; i++ {
		tags[i] = &mastodon.FollowedTag{
			Name: fmt.Sprintf("%s%d", prefix, i),
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
	return tags
}

func getListItem(l *widget.List, index int) fyne.CanvasObject {
	if index+1 > l.Length() {
		panic(fmt.Sprintf("Attemping to access list index %d, but list length is %d", index, l.Length()))
	}
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
	return listItemCanvas[1] //.(*widget.Label)
}
