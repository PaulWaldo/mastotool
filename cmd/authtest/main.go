package main

import (
	"context"
	"fmt"
	"log"
	"net/url"

	"github.com/mattn/go-mastodon"
)

func main() {
	appConfig := &mastodon.AppConfig{
		Server:       "https://stranger.social",
		ClientName:   "client-name",
		Scopes:       "read write follow",
		Website:      "https://github.com/mattn/go-mastodon",
		RedirectURIs: "urn:ietf:wg:oauth:2.0:oob",
	}
	app, err := mastodon.RegisterApp(context.Background(), appConfig)
	if err != nil {
		log.Fatal(err)
	}

	// Have the user manually get the token and send it back to us
	u, err := url.Parse(app.AuthURI)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Open your browser to \n%s\n and copy/paste the given token\n", u)
	var token string
	fmt.Print("Paste the token here:")
	fmt.Scanln(&token)
	config := &mastodon.Config{
		Server:       "https://stranger.social",
		ClientID:     app.ClientID,
		ClientSecret: app.ClientSecret,
		AccessToken:  token,
	}

	c := mastodon.NewClient(config)
	err = c.AuthenticateToken(context.Background(), token, "urn:ietf:wg:oauth:2.0:oob")
	if err != nil {
		panic(err)
	}

	acct, err := c.GetAccountCurrentUser(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Account is %v\n", acct)

	ft, err := c.GetFollowedTags(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\n\n%d Followed Tags:\n", len(ft))
	for i := 0; i < len(ft); i++ {
		fmt.Printf("Followed Tag %d: %v\n", i, ft[i])
	}

}
