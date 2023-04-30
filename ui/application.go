package ui

import (
	"context"
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/mattn/go-mastodon"
)

type preferences struct {
	MastodonServer binding.String
	// APIKey         binding.String
	ClientID     binding.String
	ClientSecret binding.String
}

const (
	ServerKey = "MastodonServer"
	APIKeyKey = "APIKey"
)

var (
	mastodonConfig *mastodon.Config
	mastodonClient *mastodon.Client
	followedTags []*mastodon.FollowedTag
)

func makeFollowedTagsUI() *fyne.Container {
	l := widget.NewList(func() int {
		return len(followedTags)
	},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(followedTags[i].Name)
		},
	)
	return container.New(layout.NewHBoxLayout(), l)

}

func setFollowedTags() {
	ft, err := mastodonClient.GetFollowedTags(context.Background(), nil)
	if err != nil {
		return nil, err
	}
}

func Run() {
	a := app.NewWithID("com.github.PaulWaldo.mastotool")
	prefs := preferences{
		MastodonServer: binding.BindPreferenceString("MastodonServer", a.Preferences()),
		ClientID:       binding.BindPreferenceString("ClientID", a.Preferences()),
		ClientSecret:   binding.BindPreferenceString("ClientSecret", a.Preferences()),
		// APIKey:         binding.BindPreferenceString(APIKeyKey, a.Preferences()),
	}
	server, _ := prefs.MastodonServer.Get()
	clientID, _ := prefs.ClientID.Get()
	clientSecret, _ := prefs.ClientSecret.Get()
	mastodonConfig = &mastodon.Config{Server: server, ClientID: clientID, ClientSecret: clientSecret}
	mastodonClient = mastodon.NewClient(mastodonConfig)
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
					}
				}, w)
				form.Resize(fyne.Size{Width: 300, Height: 300})
				form.Show()
			}),
		),
	))
	ftui, err := makeFollowedTagsUI()
	if err != nil {
		log.Fatalf("getting followed tags ui: %w", err)
	}
	w.SetContent(ftui)
	w.Resize(fyne.Size{Width: 400, Height: 400})
	w.ShowAndRun()
}
