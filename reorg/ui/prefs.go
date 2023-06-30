package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"github.com/mattn/go-mastodon"
)

// Preferences stores user data locally between application runs
type Preferences struct {
	MastodonServer binding.String // User's Mastodon server
	AccessToken    binding.String // Token provided by Mastodon
	ClientID       binding.String
	ClientSecret   binding.String
}

// NewClientFromPrefs creates a new Mastodon client from user preferences
func NewClientFromPrefs(p Preferences) *mastodon.Client {
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

func NewPreferences(a fyne.App) Preferences {
	return Preferences{
		MastodonServer: binding.BindPreferenceString("MastodonServer", a.Preferences()),
		AccessToken:    binding.BindPreferenceString("AcessToken", a.Preferences()),
		ClientID:       binding.BindPreferenceString("ClientID", a.Preferences()),
		ClientSecret:   binding.BindPreferenceString("ClientSecret", a.Preferences()),
	}
}
