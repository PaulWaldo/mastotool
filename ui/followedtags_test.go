package ui

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/widget"
	"github.com/mattn/go-mastodon"
	"github.com/stretchr/testify/assert"
)

func TestMakeFollowedTagsUI_PopulatesList(t *testing.T) {
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
	t.Parallel()
	for i := range tests {
		tt := &tests[i]
		t.Run(tt.name, func(t *testing.T) {
			a := test.NewApp()
			w := a.NewWindow("")
			keepTags := createTags("Tag", tt.fields.numFollowedTags)
			removeTags := []*mastodon.FollowedTag{}
			ma := myApp{keepTags: keepTags, removeTags: removeTags}
			container := ma.MakeFollowedTagsUI()
			allFollowedTags := createTags("Tag", tt.fields.numFollowedTags)
			w.SetContent(container)
			w.Resize(fyne.Size{Width: 400, Height: 400})
			keepList := ma.listChoices.LeftList
			removeList := ma.listChoices.RightList
			assert.Equal(t, tt.fields.numFollowedTags, keepList.Length(), "Expecting keep list widget to have %d items, got %d", len(allFollowedTags), keepList.Length())
			assert.Equal(t, 0, removeList.Length())

			for i, v := range allFollowedTags {
				got := getListItem(keepList, i).(*widget.Label)
				assert.Equal(t, v.Name, got.Text, "Expecting keep list item %d to be %s, but got %s", i, v.Name, got.Text)
			}

			// Initial state of Unfollow button should be disabled
			assert.True(t, ma.unfollowButton.Disabled(), "Initial state of Unfollow button should be disabled")
		})
	}
}

func TestFollowedTagsUI_TagMovingButtonPressesChangesUnfollowButtonEnabled(t *testing.T) {
	t.Parallel()
	allFollowedTags := createTags("Tag", 3)
	a := test.NewApp()
	w := a.NewWindow("")
	ma := myApp{}
	w.SetContent(ma.MakeFollowedTagsUI())
	ma.SetFollowedTags(allFollowedTags)
	w.Resize(fyne.Size{Width: 400, Height: 400})
	assert.Equal(t, len(allFollowedTags), ma.listChoices.LeftList.Length())
	assert.True(t, ma.unfollowButton.Disabled())

	for i := 0; i < 2; i++ {
		// Move one tag to remove list. Unfollow button should be enabled
		ma.listChoices.LeftList.Select(0)
		test.Tap(ma.listChoices.MoveRightButton)
		assert.False(t, ma.unfollowButton.Disabled())
	}
	assert.Equal(t, 2, ma.listChoices.RightList.Length())

	// Move a tag back to left. Unfollow button should be enabled
	ma.listChoices.RightList.Select(0)
	test.Tap(ma.listChoices.MoveLeftButton)
	assert.False(t, ma.unfollowButton.Disabled())

	// Move last tag back to left. Unfollow button should be disabled
	ma.listChoices.RightList.Select(0)
	test.Tap(ma.listChoices.MoveLeftButton)
	assert.True(t, ma.unfollowButton.Disabled())
}

func TestFollowedTagsUI_TappingRefreshButtonRepopulatesTags(t *testing.T) {
	// t.Parallel()
	a := test.NewApp()
	w := a.NewWindow("")
	keepTags := createTags("KTag", 3)
	removeTags := createTags("RTag", 1)
	ma := myApp{keepTags: keepTags, removeTags: removeTags, window: w, app: a, prefs: NewPreferences(a)}
	_ = ma.prefs.AccessToken.Set("access")
	_ = ma.prefs.ClientID.Set("clientid")
	_ = ma.prefs.ClientSecret.Set("secret")
	w.SetContent(ma.MakeFollowedTagsUI())
	w.Resize(fyne.Size{Width: 400, Height: 400})
	assert.Equal(t, 3, ma.listChoices.LeftList.Length())
	assert.Equal(t, 1, ma.listChoices.RightList.Length())

	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedPath := "/api/v1/followed_tags"
		if r.URL.Path != expectedPath {
			t.Fatalf("unexpected request path %s, expecting %s", r.URL.Path, expectedPath)
			return
		}
		resp := []mastodon.FollowedTag{{Name: "NewTag", Following: true}}
		enc := json.NewEncoder(w)
		err := enc.Encode(resp)
		if err != nil {
			t.Fatalf("writing mocked response: %s", err)
		}
	}))
	defer serv.Close()
	ma.prefs = NewPreferences(a)
	err := ma.prefs.MastodonServer.Set(serv.URL)
	assert.NoError(t, err)
	test.Tap(ma.refreshButton)

	assert.Equal(t, 1, ma.listChoices.LeftList.Length(), "After refresh, expecting 1 item, got %d", ma.listChoices.LeftList.Length())
	assert.Equal(t, 0, ma.listChoices.RightList.Length(), "After refresh, expecting 0 items, got %d", ma.listChoices.RightList.Length())
}
