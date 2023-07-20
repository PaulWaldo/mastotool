package ui

import (
	"context"
	"fmt"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/PaulWaldo/mastotool"
	"github.com/mattn/go-mastodon"
)

func (ma *myApp) authenticate() {
	serverUrlEntry := widget.NewEntryWithData(ma.prefs.MastodonServer)
	serverUrlEntry.Validator = nil
	serverUrlEntry.SetPlaceHolder("https://mymastodonserver.com")
	form := dialog.NewForm("Mastodon Server", "Authenticate", "Abort", []*widget.FormItem{
		{Text: "Server", Widget: serverUrlEntry, HintText: "URL of your Mastodon server"},
	}, func(confirmed bool) {
		if confirmed {
			val, _ := ma.prefs.MastodonServer.Get()
			fmt.Printf("Server is %s\n", val)
			app, err := mastodon.RegisterApp(context.Background(), mastotool.NewAuthenticationConfig(val))
			fmt.Printf("Got token %+v\n", app)
			if err != nil {
				dialog.NewError(err, ma.window).Show()
				return
			}
			_ = ma.prefs.ClientID.Set(app.ClientID)
			_ = ma.prefs.ClientSecret.Set(app.ClientSecret)
			authURI, err := url.Parse(app.AuthURI)
			if err != nil {
				dialog.NewError(err, ma.window).Show()
				return
			}
			err = ma.app.OpenURL(authURI)
			if err != nil {
				dialog.NewError(err, ma.window).Show()
				fyne.LogError(fmt.Sprintf("Calling URL.open on '%s'", authURI), err)
				return
			}
			ma.getAuthCode()
			c := NewClientFromPrefs(ma.prefs)
			_, err = c.VerifyAppCredentials(context.Background())
			if err != nil {
				dialog.NewError(err, ma.window).Show()
				fyne.LogError("In Authenticate menu, Authenticating token", err)
			}
		}
	}, ma.window)
	form.Resize(fyne.Size{Width: 300, Height: 300})
	form.Show()
}

func (ma *myApp) forgetCredentials() {
	dialog.NewConfirm("Log out", "Logging out will remove your authentication token", func(b bool) {
		if b {
			ma.prefs.forgetCredentials()
			ma.SetFollowedTags([]*mastodon.FollowedTag{})
		}
	}, ma.window).Show()
}

// getAuthCode allows the user to input the Authentication Token provided by Mastodon into the preferences
func (ma *myApp) getAuthCode() {
	accessTokenEntry := widget.NewEntry()
	accessTokenEntry.Validator = nil
	dialog.NewForm("Authorization Code", "Save", "Cancel", []*widget.FormItem{
		{
			Text:     "Authorization Code",
			Widget:   accessTokenEntry,
			HintText: "XXXXXXXXXXXXXXX",
		}},
		func(confirmed bool) {
			if confirmed {
				c := NewClientFromPrefs(ma.prefs)
				fmt.Printf("After authorizing, client is \n%+v\n", c.Config)
				err := c.AuthenticateToken(context.Background(), accessTokenEntry.Text, "urn:ietf:wg:oauth:2.0:oob")
				if err != nil {
					dialog.NewError(err, ma.window).Show()
					fyne.LogError("Authenticating token", err)
					return
				}
				_ = ma.prefs.AccessToken.Set(c.Config.AccessToken)
				// err = ma.getFollowedTags()
				// if err != nil {
				// 	dialog.NewError(err, w).Show()
				// 	fyne.LogError("Getting followed tags after auth", err)
				// 	return
				// }
			}
		},
		ma.window).Show()
}