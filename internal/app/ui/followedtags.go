package ui

import (
	"context"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/PaulWaldo/mastotool/internal/app"
	"github.com/mattn/go-mastodon"
)

func (ma *myApp) MakeFollowedTagsUI() fyne.CanvasObject {
	ma.listChoices = NewListChoices()

	ma.listChoices.LeftLabel.Text = "To Keep"
	ma.listChoices.LeftLabel.TextStyle = fyne.TextStyle{Bold: true}
	ma.listChoices.LeftLabel.Alignment = fyne.TextAlignCenter

	ma.listChoices.RightLabel.Text = "To Remove"
	ma.listChoices.RightLabel.TextStyle = fyne.TextStyle{Bold: true}
	ma.listChoices.RightLabel.Alignment = fyne.TextAlignCenter

	ma.listChoices.LeftItems = ma.keepTags
	ma.listChoices.RightItems = ma.removeTags

	oldMoveRightAction := ma.listChoices.MoveRightButton.OnTapped
	ma.listChoices.MoveRightButton.OnTapped = func() {
		oldMoveRightAction()
		ma.unfollowButton.Enable()
	}

	oldMoveLeftAction := ma.listChoices.MoveLeftButton.OnTapped
	ma.listChoices.MoveLeftButton.OnTapped = func() {
		oldMoveLeftAction()
		if ma.listChoices.RightList.Length() > 0 {
			ma.unfollowButton.Enable()
		} else {
			ma.unfollowButton.Disable()
		}
	}

	ma.unfollowButton = widget.NewButtonWithIcon("Unfollow", theme.DeleteIcon(), func() {
		dialog.ShowCustomConfirm(
			"Confirm Unfollow", "Unfollow", "Cancel",
			makeRemoveConfirmUI(ma.listChoices.RightItems), func(b bool) {
				if b {
					c := NewClientFromPrefs(ma.prefs)
					err := app.RemoveFollowedTags(*c, ma.listChoices.RightItems)
					if err != nil {
						dialog.NewError(err, ma.window)
						return
					} else {
						tags := make([]string, len(ma.listChoices.RightItems))
						for i, t := range ma.listChoices.RightItems {
							tags[i] = t.Name
						}
						dialog.NewInformation(
							"Success",
							"Tags successfully unfollowed",
							ma.window).Show()
					}
					ma.refreshFollowedTags()
				}
			}, ma.window)
	})
	ma.unfollowButton.Disable()

	ma.refreshButton = widget.NewButtonWithIcon("Refresh", theme.ViewRefreshIcon(), func() {
		ma.refreshFollowedTags()
	})
	bottom := container.NewBorder(nil, nil, nil, ma.refreshButton, ma.unfollowButton)
	return container.NewBorder(nil, bottom, nil, nil, ma.listChoices)
}

func (ma *myApp) SetFollowedTags(t []*mastodon.FollowedTag) {
	if ma.listChoices == nil {
		return
	}
	ma.listChoices.SetLeftItems(t)
	ma.listChoices.SetRightItems([]*mastodon.FollowedTag{})
	ma.listChoices.Refresh()
}

// RemoveFollowedTags removes the given list of tags from the user's following list.
// These tags will no longer show up in the user's feed
func (ma *myApp) RemoveFollowedTags(w fyne.Window) func() {
	return func() {
		dialog.ShowCustomConfirm("Confirm Unfollow", "Unfollow", "Cancel", makeRemoveConfirmUI(ma.listChoices.RightItems), func(b bool) {
			if b {
				c := NewClientFromPrefs(ma.prefs)
				err := app.RemoveFollowedTags(*c, ma.listChoices.RightItems)
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
				ma.refreshFollowedTags()
			}
		}, w)
	}
}

// refreshFollowedTags gets the list of followed tags and populates the keepTags and removeTags based on this
func (ma *myApp) refreshFollowedTags() {
	if !ma.isLoggedIn() {
		return
	}
	c := NewClientFromPrefs(ma.prefs)
	tags, err := c.GetFollowedTags(context.Background(), nil)
	if err != nil {
		tags = []*mastodon.FollowedTag{}
		dialog.ShowError(err, ma.window)
	}
	ma.SetFollowedTags(tags)
}

func (ma *myApp) isLoggedIn() bool {
	s, err := ma.prefs.MastodonServer.Get()
	if err != nil || len(s) == 0 {
		return false
	}
	s, err = ma.prefs.AccessToken.Get()
	if err != nil || len(s) == 0 {
		return false
	}
	s, err = ma.prefs.ClientID.Get()
	if err != nil || len(s) == 0 {
		return false
	}
	s, err = ma.prefs.ClientSecret.Get()
	if err != nil || len(s) == 0 {
		return false
	}

	return true
}

// makeRemoveConfirmUI creates the UI for the tag removal confirmation action
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
