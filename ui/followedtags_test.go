package ui

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/test"
	"github.com/mattn/go-mastodon"
	"github.com/stretchr/testify/assert"
)

// func Test_a(t *testing.T) {
// 	tags := []*mastodon.FollowedTag{
// 		{Name: "name1"},
// 		{Name: "name2"},
// 		{Name: "name3"},
// 	}
// 	inter := make([]interface{}, len(tags))

// 	for i, v := range tags {
// 		inter[i] = v
// 	}
// 	bl := binding.NewUntypedList()
// 	err := bl.Set(inter)
// 	assert.NoError(t, err)
// 	a := test.NewApp()
// 	list := NewBoundList(bl)
// 	w := a.NewWindow("")
// 	w.SetContent(list)
// 	assert.Equal(t, len(tags), list.Length())
// }

// func Test_UI(t *testing.T) {
// 	tags := []*mastodon.FollowedTag{
// 		{Name: "name1"},
// 		{Name: "name2"},
// 		{Name: "name3"},
// 	}

// 	a := test.NewApp()
// 	w := a.NewWindow("")
// 	ui := NewFollowedTagsUI()
// 	ui.SetFollowedTags(tags)
// 	w.SetContent(ui.MakeFollowedTagsUI())
// 	assert.NotNil(t, ui.keepListWidget)
// 	assert.Equal(t, len(tags), ui.keepListWidget.Length())
// 	wr := test.WidgetRenderer(ui.keepListWidget)
// 	items := wr.Objects()
// 	assert.Equal(t, items[0], "name1")
// }

func TestFollowedTagsUI_MakeFollowedTagsUI(t *testing.T) {
	type fields struct {
		followedTags []*mastodon.FollowedTag
	}
	tests := []struct {
		name   string
		fields fields
		want   *fyne.Container
	}{
		{
			name: "A few followed tags",
			fields: fields{
				followedTags: []*mastodon.FollowedTag{
					{
						Name: "tag1",
						History: []mastodon.FollowedTagHistory{
							{
								Day:      mastodon.UnixTimeString{Time: time.Now()},
								Accounts: 300,
								Uses:     3000,
							},
							{
								Day:      mastodon.UnixTimeString{Time: time.Now().AddDate(0, 0, -1)},
								Accounts: 200,
								Uses:     3000,
							},
							{
								Day:      mastodon.UnixTimeString{Time: time.Now().AddDate(0, 0, -2)},
								Accounts: 100,
								Uses:     1000,
							},
						},
					},
					{
						Name: "tag2",
						History: []mastodon.FollowedTagHistory{
							{
								Day:      mastodon.UnixTimeString{Time: time.Now()},
								Accounts: 150,
								Uses:     1500,
							},
							{
								Day:      mastodon.UnixTimeString{Time: time.Now().AddDate(0, 0, -1)},
								Accounts: 100,
								Uses:     1000,
							},
							{
								Day:      mastodon.UnixTimeString{Time: time.Now().AddDate(0, 0, -2)},
								Accounts: 50,
								Uses:     500,
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = test.NewApp()
			ui := &FollowedTagsUI{KeepTags: tt.fields.followedTags}
			if got := ui.MakeFollowedTagsUI(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FollowedTagsUI.MakeFollowedTagsUI() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestFollowedTagsUI_KeepTagsPopulate(t *testing.T) {
	followedTags := []*mastodon.FollowedTag{
		{
			Name: "tag1",
			History: []mastodon.FollowedTagHistory{
				{
					Day:      mastodon.UnixTimeString{Time: time.Now()},
					Accounts: 300,
					Uses:     3000,
				},
				{
					Day:      mastodon.UnixTimeString{Time: time.Now().AddDate(0, 0, -1)},
					Accounts: 200,
					Uses:     3000,
				},
				{
					Day:      mastodon.UnixTimeString{Time: time.Now().AddDate(0, 0, -2)},
					Accounts: 100,
					Uses:     1000,
				},
			},
		},
		{
			Name: "tag2",
			History: []mastodon.FollowedTagHistory{
				{
					Day:      mastodon.UnixTimeString{Time: time.Now()},
					Accounts: 150,
					Uses:     1500,
				},
				{
					Day:      mastodon.UnixTimeString{Time: time.Now().AddDate(0, 0, -1)},
					Accounts: 100,
					Uses:     1000,
				},
				{
					Day:      mastodon.UnixTimeString{Time: time.Now().AddDate(0, 0, -2)},
					Accounts: 50,
					Uses:     500,
				},
			},
		},
	}
	_ = test.NewApp()
	ui := &FollowedTagsUI{KeepTags: followedTags}
	_ = ui.MakeFollowedTagsUI()
	keepList := ui.keepListWidget
	assert.Equal(t, len(followedTags), keepList.Length())
	renderer := test.WidgetRenderer(ui.keepListWidget)
	items := renderer.Objects()
	scrollRenderer := test.WidgetRenderer(items[0].(fyne.Widget))
	scrollObjs:= scrollRenderer.Objects()
	fmt.Printf("Objs: %+v\n", scrollObjs)
	xx := scrollObjs[0].(*fyne.Container)
	fmt.Printf("xx: %+v\n", xx)
	yy:=xx.Objects[0].(fyne.Widget)
	fmt.Printf("yy: %+v\n", yy)
}
