package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/PaulWaldo/mastotool"
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
					err := mastotool.RemoveFollowedTags(*c, ma.listChoices.RightItems)
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
					tags, err := GetFollowedTags(ma.prefs)
					if err != nil {
						dialog.NewError(err, ma.window)
					}
					ma.SetFollowedTags(tags)
				}
			}, ma.window)
	})
	ma.unfollowButton.Disable()

	ma.refreshButton = widget.NewButtonWithIcon("Refresh", theme.ViewRefreshIcon(), func() {
		ma.getFollowedTags()
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
	ma.listChoices.container.Refresh()
}
