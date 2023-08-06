package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"github.com/mattn/go-mastodon"
)

const (
	PrefKeyServer       = "MastodonServer"
	PrefKeyAccessToken  = "AcessToken"
	PrefKeyClientID     = "ClientID"
	PrefKeyClientSecret = "ClientSecret"
)

// AppPrefs stores user data locally between application runs
type AppPrefs struct {
	MastodonServer binding.String // User's Mastodon server
	AccessToken    binding.String // Token provided by Mastodon
	ClientID       binding.String
	ClientSecret   binding.String
}

// NewClientFromPrefs creates a new Mastodon client from user preferences
func NewClientFromPrefs(p AppPrefs) *mastodon.Client {
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

func NewPreferences(a fyne.App) AppPrefs {
	return AppPrefs{
		MastodonServer: binding.BindPreferenceString(PrefKeyServer, a.Preferences()),
		AccessToken:    binding.BindPreferenceString(PrefKeyAccessToken, a.Preferences()),
		ClientID:       binding.BindPreferenceString(PrefKeyClientID, a.Preferences()),
		ClientSecret:   binding.BindPreferenceString(PrefKeyClientSecret, a.Preferences()),
	}
}

func ClearCredentialsPrefs() {
	p := fyne.CurrentApp().Preferences()
	p.RemoveValue(PrefKeyAccessToken)
	p.RemoveValue(PrefKeyClientID)
	p.RemoveValue(PrefKeyClientSecret)
	p.RemoveValue(PrefKeyServer)
}
