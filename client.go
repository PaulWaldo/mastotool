package mastotool

import (
	"context"
	"fmt"
	"net/url"

	"github.com/mattn/go-mastodon"
)

type ApiClient struct {
	app    *mastodon.Application
	client *mastodon.Client
	Server string
}

const (
	clientName = "Mastotool"
	website    = "https://github.com/PaulWaldo/mastotool"
)

func (ac *ApiClient) AuthenticationURL() (*url.URL, error) {
	var err error
	ac.app, err = mastodon.RegisterApp(context.Background(), &mastodon.AppConfig{
		Server:       ac.Server,
		ClientName:   clientName,
		Scopes:       "read write follow",
		Website:      website,
		RedirectURIs: "urn:ietf:wg:oauth:2.0:oob",
	})
	if err != nil {
		return nil, err
	}
	fmt.Printf("client-id    : %s\n", ac.app.ClientID)
	fmt.Printf("client-secret: %s\n", ac.app.ClientSecret)
	fmt.Printf("Auth URL: %s\n", ac.app.AuthURI)

	u, err := url.Parse(ac.app.AuthURI)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// a := app.NewWithID("com.github.PaulWaldo.mastotool")
// w := a.NewWindow("TODO List")

// authTokenKey := "authToken"
// authToken := a.Preferences().String(authTokenKey)
// authToken := ""
// var mApp *mastodon.Application
// var err error
// var c *mastodon.Client
// if "" == authToken {
// 	mApp, err = mastodon.RegisterApp(context.Background(), &mastodon.AppConfig{
// 		Server:       "https://stranger.social",
// 		ClientName:   "client-name",
// 		Scopes:       "read write follow",
// 		Website:      "https://github.com/mattn/go-mastodon",
// 		RedirectURIs: "urn:ietf:wg:oauth:2.0:oob",
// 	})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("client-id    : %s\n", mApp.ClientID)
// 	fmt.Printf("client-secret: %s\n", mApp.ClientSecret)
// 	fmt.Printf("Auth URL: %s\n", mApp.AuthURI)

// 	u, err := url.Parse(mApp.AuthURI)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	err = a.OpenURL(u)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// input := widget.NewEntry()
// 	// input.SetPlaceHolder("Authorization code from browser")
// 	// content := container.NewVBox(input, widget.NewButton("OK", func() {
// 	// 	authToken = input.Text
// 	// 	a.Preferences().SetString(authTokenKey, authToken)
// 	// }))
// 	// w.SetContent(content)
// 	// w.ShowAndRun()
// 	fmt.Println("Authorization code from browser:")
// 	n, err := fmt.Scanf("%s", &authToken)
// 	if n != 1 || err != nil {
// 		log.Fatalf("error scanning: n=%d, err=%s", n, err)
// 	}

// 	c = mastodon.NewClient(&mastodon.Config{
// 		Server:       "https://stranger.social",
// 		ClientID:     mApp.ClientID,
// 		ClientSecret: mApp.ClientSecret,
// 		AccessToken:  authToken,
// 	})
// 	err = c.AuthenticateToken(context.Background(), authToken, "urn:ietf:wg:oauth:2.0:oob")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	acct, err := c.GetAccountCurrentUser(context.Background())
// 	// timeline, err := c.GetTimelineHome(context.Background(), nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("Account is %v\n", acct)
