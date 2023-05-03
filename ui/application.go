package ui

import (
	"context"
	"fmt"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
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
}

func (ma *myApp) makeFollowedTagsUI() *fyne.Container {
	ma.keepTags = binding.BindIntList(&[]int{})
	ma.removeTags = binding.BindIntList(&[]int{})
	keep := widget.NewListWithData(ma.keepTags,
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			o.(*widget.Label).Bind(i.(binding.String))
		},
	)
	remove := widget.NewListWithData(ma.removeTags,
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			o.(*widget.Label).Bind(i.(binding.String))
		},
	)
	return container.New(layout.NewHBoxLayout(), keep, remove)

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
	fmt.Printf("Got %d followed tags\n", len(ma.followedTags))
	for i := 0; i < len(ma.followedTags); i++ {
		keepIndex[i] = i
	}
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
						u, err := url.Parse(app.AuthURI)
						if err != nil {
							dialog.NewError(err, w).Show()
							return
						}
						err = a.OpenURL(u)
						if err != nil {
							dialog.NewError(err, w).Show()
							fyne.LogError("Calling URL.open", err)
							return
						}
						AccessTokenEntry := widget.NewEntryWithData(prefs.AccessToken)
						AccessTokenEntry.Validator = nil
						dialog.NewForm("Authorization Code", "Save", "Cancel", []*widget.FormItem{
							{
								Text:     "Authorization Code",
								Widget:   AccessTokenEntry,
								HintText: "XXXXXXXXXXXXXXX",
							}},
							func(b bool) {
								if b {
									c := NewClientFromPrefs(myApp.prefs)
									fmt.Printf("After authorizing, client is \n%+v\n", c.Config)
									at, _ := myApp.prefs.AccessToken.Get()
									err = c.AuthenticateToken(context.Background(), at, "urn:ietf:wg:oauth:2.0:oob")
									if err != nil {
										dialog.NewError(err, w).Show()
										fyne.LogError("Authenticating token", err)
										return
									}
									err = myApp.getFollowedTags()
									if err != nil {
										dialog.NewError(err, w).Show()
										fyne.LogError("Getting followed tags after auth", err)
										return
									}
								}
							},
							w).Show()
					}
				}, w)
				form.Resize(fyne.Size{Width: 300, Height: 300})
				form.Show()
			}),
		),
	))
	at, _ := myApp.prefs.AccessToken.Get()
	c := NewClientFromPrefs(myApp.prefs)
	err := c.AuthenticateToken(context.Background(), at, "urn:ietf:wg:oauth:2.0:oob")
	if err != nil {
		dialog.NewError(err, w).Show()
		fyne.LogError("In main, Authenticating token", err)
	}
	err = myApp.getFollowedTags()
	if err != nil {
		dialog.NewError(err, w).Show()
		fyne.LogError("In main, getting followed tags", err)
	}
	w.SetContent(myApp.makeFollowedTagsUI())
	w.Resize(fyne.Size{Width: 400, Height: 400})
	w.ShowAndRun()
}
