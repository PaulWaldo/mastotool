package ui

import (
	"context"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/PaulWaldo/mastotool"
	"github.com/mattn/go-mastodon"
)

type myApp struct {
	prefs                Preferences
	app                  fyne.App
	window               fyne.Window
	keepTags, removeTags []*mastodon.FollowedTag
	listChoices          *ListChoices
	unfollowButton       *widget.Button
}

func Run() {
	ma := myApp{}
	ma.app = app.NewWithID("com.github.PaulWaldo.mastotool")
	ma.prefs = NewPreferences(ma.app)
	ma.window = ma.app.NewWindow("MastoTool")
	ma.window.SetMainMenu(fyne.NewMainMenu(
		fyne.NewMenu("File", fyne.NewMenuItem("Authenticate", ma.authenticate))),
	)
	ma.window.SetContent(ma.MakeFollowedTagsUI())
	ma.window.Resize(fyne.Size{Width: 400, Height: 400})
	go ma.getFollowedTags()
	ma.window.ShowAndRun()
}

// getFollowedTags gets the list of followed tags and populates the keepTags and removeTags based on this
func (ma *myApp) getFollowedTags() {
	c := NewClientFromPrefs(ma.prefs)
	// var err error
	tags, err := c.GetFollowedTags(context.Background(), nil)
	if err != nil {
		tags = []*mastodon.FollowedTag{}
		dialog.ShowError(err, ma.window)
	}
	// ma.listChoices.RightItems = []*mastodon.FollowedTag{}
	ma.SetFollowedTags(tags)
}

func makeRemoveConfirmUI(tags []*mastodon.FollowedTag) fyne.CanvasObject {
	var sb strings.Builder
	sb.WriteString("These tags will be removed from your following list and not seen in your feed:")
	for _, t := range tags {
		sb.WriteString("\n* " + t.Name)
	}
	rt := widget.NewRichTextFromMarkdown(sb.String())
	rt.Wrapping = fyne.TextWrapWord
	scroll := container.NewVScroll(rt)
	scroll.SetMinSize(fyne.Size{Width: 100, Height: 150})
	return scroll
}

// RemoveFollowedTags removes the given list of tags from the user's following list.
// These tags will no longer show up in the user's feed
func (ma *myApp) RemoveFollowedTags(w fyne.Window) func() {
	return func() {
		dialog.ShowCustomConfirm("Confirm Unfollow", "Unfollow", "Cancel", makeRemoveConfirmUI(ma.listChoices.RightItems), func(b bool) {
			if b {
				c := NewClientFromPrefs(ma.prefs)
				err := mastotool.RemoveFollowedTags(*c, ma.listChoices.RightItems)
				if err != nil {
					dialog.NewError(err, w)
					return
				} else {
					tags := make([]string, len(ma.listChoices.RightItems))
					for i, t := range ma.listChoices.RightItems {
						tags[i] = t.Name
					}
					dialog.NewInformation(
						"Success",
						"Tags successfully unfollowed",
						w).Show()
				}
				ma.getFollowedTags()
				// TODO: this does not work because the data store is not connected properly
				// ma.ftui.container.Refresh()
			}
		}, w)
	}
}

func GetFollowedTags(prefs Preferences) ([]*mastodon.FollowedTag, error) {
	var err error
	c := NewClientFromPrefs(prefs)
	followedTags, err := c.GetFollowedTags(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	return followedTags, nil
}
