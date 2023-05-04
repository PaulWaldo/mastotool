package ui

import (
	"context"
	"fmt"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/PaulWaldo/mastotool"
	"github.com/mattn/go-mastodon"
)

type preferences struct {
	MastodonServer binding.String
	AccessToken    binding.String
	ClientID       binding.String
	ClientSecret   binding.String
}

const (
	ServerKey = "MastodonServer"
	APIKeyKey = "APIKey"
)

type myApp struct {
	followedTags []*mastodon.FollowedTag
	keepTags     binding.ExternalIntList
	removeTags   binding.ExternalIntList
	prefs        preferences
	ftui         FollowedTagsUI
}

func (ma *myApp) getFollowedTags() error {
	var err error
	c := NewClientFromPrefs(ma.prefs)
	fmt.Printf("Getting followed tags, client is \n%+v\n", c.Config)
	ma.followedTags, err = c.GetFollowedTags(context.Background(), nil)
	if err != nil {
		return err
	}
	keepIndex := make([]int, len(ma.followedTags))
	fmt.Printf("Got %d followed tags: ", len(ma.followedTags))
	for i := 0; i < len(ma.followedTags); i++ {
		keepIndex[i] = i
		fmt.Printf("%s, ", ma.followedTags[i].Name)
	}
	fmt.Println("")
	ma.keepTags = binding.BindIntList(&keepIndex)
	ma.removeTags = binding.BindIntList(&[]int{})
	return nil
}

func NewClientFromPrefs(p preferences) *mastodon.Client {
	server, _ := p.MastodonServer.Get()
	clientID, _ := p.ClientID.Get()
	clientSecret, _ := p.ClientSecret.Get()
	accessToken, _ := p.AccessToken.Get()
	conf := &mastodon.Config{
		Server:       server,
		ClientID:     clientID,
		ClientSecret: clientSecret,
		AccessToken:  accessToken,
	}
	return mastodon.NewClient(conf)
}

var authURI *url.URL

func (ma *myApp) getAuthCode(w fyne.Window) {
	accessTokenEntry := widget.NewEntry()
	accessTokenEntry.Validator = nil
	dialog.NewForm("Authorization Code", "Save", "Cancel", []*widget.FormItem{
		{
			Text:     "Authorization Code",
			Widget:   accessTokenEntry,
			HintText: "XXXXXXXXXXXXXXX",
		}},
		func(b bool) {
			if b {
				c := NewClientFromPrefs(ma.prefs)
				fmt.Printf("After authorizing, client is \n%+v\n", c.Config)
				err := c.AuthenticateToken(context.Background(), accessTokenEntry.Text, "urn:ietf:wg:oauth:2.0:oob")
				if err != nil {
					dialog.NewError(err, w).Show()
					fyne.LogError("Authenticating token", err)
					return
				}
				_ = ma.prefs.AccessToken.Set(c.Config.AccessToken)
				err = ma.getFollowedTags()
				if err != nil {
					dialog.NewError(err, w).Show()
					fyne.LogError("Getting followed tags after auth", err)
					return
				}
			}
		},
		w).Show()
}

func Run() {
	a := app.NewWithID("com.github.PaulWaldo.mastotool")
	prefs := preferences{
		MastodonServer: binding.BindPreferenceString("MastodonServer", a.Preferences()),
		AccessToken:    binding.BindPreferenceString("AcessToken", a.Preferences()),
		ClientID:       binding.BindPreferenceString("ClientID", a.Preferences()),
		ClientSecret:   binding.BindPreferenceString("ClientSecret", a.Preferences()),
	}
	myApp := &myApp{prefs: prefs}
	w := a.NewWindow("MastoTool")
	w.SetMainMenu(fyne.NewMainMenu(
		fyne.NewMenu("File",
			fyne.NewMenuItem("Authenticate", func() {
				serverUrlEntry := widget.NewEntryWithData(prefs.MastodonServer)
				serverUrlEntry.Validator = nil
				serverUrlEntry.SetPlaceHolder("https://mymastodonserver.com")
				form := dialog.NewForm("Mastodon Server", "Authenticate", "Abort", []*widget.FormItem{
					{Text: "Server", Widget: serverUrlEntry, HintText: "URL of your Mastodon server"},
				}, func(b bool) {
					if b {
						val, _ := prefs.MastodonServer.Get()
						fmt.Printf("Server is %s\n", val)
						app, err := mastodon.RegisterApp(context.Background(), mastotool.NewAuthenticationConfig(val))
						fmt.Printf("Got token %+v\n", app)
						if err != nil {
							dialog.NewError(err, w).Show()
							return
						}
						_ = prefs.ClientID.Set(app.ClientID)
						_ = prefs.ClientSecret.Set(app.ClientSecret)
						authURI, err = url.Parse(app.AuthURI)
						if err != nil {
							dialog.NewError(err, w).Show()
							return
						}
						err = a.OpenURL(authURI)
						if err != nil {
							dialog.NewError(err, w).Show()
							fyne.LogError(fmt.Sprintf("Calling URL.open on '%s'", authURI), err)
							return
						}
						myApp.getAuthCode(w)
						c := NewClientFromPrefs(myApp.prefs)
						_, err = c.VerifyAppCredentials(context.Background())
						if err != nil {
							dialog.NewError(err, w).Show()
							fyne.LogError("In Authenticate menu, Authenticating token", err)
						}
					}
				}, w)
				form.Resize(fyne.Size{Width: 300, Height: 300})
				form.Show()
			}),
		),
	))
	c := NewClientFromPrefs(myApp.prefs)
	_, err := c.VerifyAppCredentials(context.Background())
	if err != nil {
		dialog.NewError(err, w).Show()
		fyne.LogError("In main, Authenticating token", err)
	}
	err = myApp.getFollowedTags()
	if err != nil {
		dialog.NewError(err, w).Show()
		fyne.LogError("In main, getting followed tags", err)
	}

	myApp.ftui = *NewFollowedTagsUI()
	myApp.ftui.MakeFollowedTagsUI()
	myApp.ftui.SetFollowedTags(myApp.followedTags)
	w.SetContent(myApp.ftui.container)
	w.Resize(fyne.Size{Width: 400, Height: 400})
	w.ShowAndRun()
}
