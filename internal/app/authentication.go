package app

import (
	"context"
	"net/url"

	"github.com/mattn/go-mastodon"
)

const (
	clientName = "Mastotool"
	website    = "https://github.com/PaulWaldo/mastotool"
)

func NewAuthenticationConfig(server string) *mastodon.AppConfig {
	return &mastodon.AppConfig{
		Server:       server,
		ClientName:   clientName,
		Scopes:       "read write follow",
		Website:      website,
		RedirectURIs: "urn:ietf:wg:oauth:2.0:oob",
	}
}

// AuthenticationURL returns a URL that authenticates the user to his account.
func AuthenticationURL(appConfig *mastodon.AppConfig) (*url.URL, error) {
	app, err := mastodon.RegisterApp(context.Background(), appConfig)
	if err != nil {
		return nil, err
	}

	u, err := url.Parse(app.AuthURI)
	if err != nil {
		return nil, err
	}
	return u, nil
}
