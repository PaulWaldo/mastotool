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
	AuthCode       binding.String
	// APIKey         binding.String
	ClientID     binding.String
	ClientSecret binding.String
}

const (
	ServerKey = "MastodonServer"
	APIKeyKey = "APIKey"
)

type myApp struct {
	mastodonConfig    *mastodon.Config
	mastodonClient    *mastodon.Client
	authorizationCode string
	followedTags      []*mastodon.FollowedTag
	keepTags          binding.ExternalIntList
	removeTags        binding.ExternalIntList
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
	fmt.Printf("Getting followed tags, client is \n%+v\n", ma.mastodonClient.Config)
	ma.followedTags, err = ma.mastodonClient.GetFollowedTags(context.Background(), nil)
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

func Run() {
	a := app.NewWithID("com.github.PaulWaldo.mastotool")
	prefs := preferences{
		MastodonServer: binding.BindPreferenceString("MastodonServer", a.Preferences()),
		AuthCode:       binding.BindPreferenceString("AuthorizationCode", a.Preferences()),
		ClientID:       binding.BindPreferenceString("ClientID", a.Preferences()),
		ClientSecret:   binding.BindPreferenceString("ClientSecret", a.Preferences()),
		// APIKey:         binding.BindPreferenceString(APIKeyKey, a.Preferences()),
	}
	server, _ := prefs.MastodonServer.Get()
	clientID, _ := prefs.ClientID.Get()
	clientSecret, _ := prefs.ClientSecret.Get()
	authToken, _ := prefs.AuthCode.Get()
	mastodonConfig := &mastodon.Config{Server: server, AccessToken: authToken, ClientID: clientID, ClientSecret: clientSecret}
	mastodonClient := mastodon.NewClient(mastodonConfig)
	myApp := &myApp{mastodonConfig: mastodonConfig, mastodonClient: mastodonClient}
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
						// app, err := mastodon.RegisterApp(context.Background(), appConfig)
						app, err := mastodon.RegisterApp(context.Background(), mastotool.NewAuthenticationConfig(val))
						fmt.Printf("Got token %+v\n", app)
						// u, err := mastotool.AuthenticationURL(mastotool.NewAuthenticationConfig(val))
						if err != nil {
							dialog.NewError(err, w).Show()
							return
						}
						prefs.ClientID.Set(app.ClientID)
						myApp.mastodonConfig.ClientID = app.ClientID
						prefs.ClientSecret.Set(app.ClientSecret)
						myApp.mastodonConfig.ClientSecret = app.ClientSecret
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
						authCodeEntry := widget.NewEntryWithData(prefs.AuthCode)
						authCodeEntry.Validator = nil
						dialog.NewForm("Authorization Code", "Save", "Cancel", []*widget.FormItem{
							{
								Text:     "Authorization Code",
								Widget:   authCodeEntry,
								HintText: "XXXXXXXXXXXXXXX",
							}},
							func(b bool) {
								if b {
									val, _ := prefs.AuthCode.Get()
									myApp.authorizationCode = val
									myApp.mastodonConfig.AccessToken = val
									fmt.Printf("After authorizing, client is \n%+v\n", myApp.mastodonConfig)
									myApp.mastodonClient = mastodon.NewClient(myApp.mastodonConfig)
									// c := mastodon.NewClient(myApp.mastodonConfig)
									err = myApp.mastodonClient.AuthenticateToken(context.Background(), myApp.authorizationCode, "urn:ietf:wg:oauth:2.0:oob")
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
	err := myApp.getFollowedTags()
	if err != nil {
		dialog.NewError(err, w).Show()
		fyne.LogError("In main, getting followed tags", err)
	}
	w.SetContent(myApp.makeFollowedTagsUI())
	w.Resize(fyne.Size{Width: 400, Height: 400})
	w.ShowAndRun()
}
