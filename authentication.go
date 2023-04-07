package mastotool

import (
	"context"
	"fmt"
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
func AuthenticationURL(appConfig *mastodon.AppConfig) (*url.URL, error) {
	app, err := mastodon.RegisterApp(context.Background(), appConfig)
	if err != nil {
		return nil, err
	}
	fmt.Printf("client-id    : %s\n", app.ClientID)
	fmt.Printf("client-secret: %s\n", app.ClientSecret)
	fmt.Printf("Auth URL: %s\n", app.AuthURI)

	u, err := url.Parse(app.AuthURI)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"net/url"
// )

// type newApplicationClient struct {
// 	Hostname string
// 	Client   *http.Client
// }

// func NewCreateApplicationClient(hostname string) *newApplicationClient {
// 	return &newApplicationClient{
// 		Hostname: hostname,
// 		Client:   &http.Client{},
// 	}
// }

// // type Config struct {
// // 	Hostname      string
// // 	AppWebsiteURI string
// // }

// type createApplicationResponse struct {
// 	ID           string `json:"id,omitempty"`
// 	Name         string `json:"name,omitempty"`
// 	Website      string `json:"website,omitempty"`
// 	RedirectURI  string `json:"redirect_uri,omitempty"`
// 	ClientID     string `json:"client_id,omitempty"`
// 	ClientSecret string `json:"client_secret,omitempty"`
// 	VapidKey     string `json:"vapid_key,omitempty"`
// }

// func (c *newApplicationClient) createApplication(website string) (*createApplicationResponse, error) {
// 	URI := fmt.Sprintf("%s/api/v1/apps", c.Hostname)
// 	data := url.Values{
// 		"client_name":   {"Test Application"},
// 		"redirect_uris": {"urn:ietf:wg:oauth:2.0:oob"},
// 		"scopes":        {"read write push"},
// 		"website":       {website},
// 	}
// 	resp, err := c.Client.PostForm(URI, data)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if resp.StatusCode != http.StatusOK {
// 		return nil, fmt.Errorf("Got %d calling apps endpoint: %s", resp.StatusCode, resp.Status)
// 	}
// 	defer resp.Body.Close()

// 	bodyBytes, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return nil, err
// 	}
// 	bodyString := string(bodyBytes)
// 	fmt.Println("API Response as String:\n" + bodyString)

// 	var car createApplicationResponse
// 	err = json.Unmarshal(bodyBytes, &car)
// 	if err != nil {
// 		return nil, err
// 	}
// 	fmt.Printf("API Response as struct %+v\n", car)
// 	return &car, nil
// }

// type getTokenResponse struct {
// 	AccessToken string `json:"access_token,omitempty"`
// 	TokenType   string `json:"token_type,omitempty"`
// 	Scope       string `json:"scope,omitempty"`
// 	CreatedAt   int    `json:"created_at,omitempty"`
// }

// func (c *newApplicationClient) getToken(clientID string, clientSecret string) (*getTokenResponse, error) {
// 	URI := fmt.Sprintf("%s/oauth/token", c.Hostname)
// 	data := url.Values{
// 		"client_id":     {clientID},
// 		"client_secret": {clientSecret},
// 		"redirect_uris": {"urn:ietf:wg:oauth:2.0:oob"},
// 		"grant_type":    {"client_credentials"},
// 	}
// 	resp, err := c.Client.PostForm(URI, data)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if resp.StatusCode != http.StatusOK {
// 		return nil, fmt.Errorf("Got %d calling apps endpoint: %s", resp.StatusCode, resp.Status)
// 	}
// 	defer resp.Body.Close()

// 	bodyBytes, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return nil, err
// 	}
// 	bodyString := string(bodyBytes)
// 	fmt.Println("API Response as String:\n" + bodyString)

// 	var gtr getTokenResponse
// 	err = json.Unmarshal(bodyBytes, &gtr)
// 	if err != nil {
// 		return nil, err
// 	}
// 	fmt.Printf("API Response as struct %+v\n", gtr)
// 	return &gtr, nil
// }

// // func (c *newApplicationClient) CreateApplication() (*createApplicationResponse, error) {
// // 	car := c.createApplication()
// // 	appCreateURI := fmt.Sprintf("%s/api/v1/apps", c.Hostname)
// // 	data := url.Values{
// // 		"client_name":   {"Test Application"},
// // 		"redirect_uris": {"urn:ietf:wg:oauth:2.0:oob"},
// // 		"scopes":        {"read write push"},
// // 		"website":       {c.AppWebsiteURI},
// // 	}
// // 	resp1, err := c.Client.PostForm(appCreateURI, data)
// // 	if err != nil {
// // 		return nil, err
// // 	}
// // 	if resp1.StatusCode != http.StatusOK {
// // 		return nil, fmt.Errorf("Got %d calling apps endpoint: %s", resp1.StatusCode, resp1.Status)
// // 	}
// // 	defer resp1.Body.Close()
// // 	bodyBytes, _ := io.ReadAll(resp1.Body)

// // 	bodyString := string(bodyBytes)
// // 	fmt.Println("API Response as String:\n" + bodyString)

// // 	var car createApplicationResponse
// // 	err = json.Unmarshal(bodyBytes, &car)
// // 	if err != nil {
// // 		return nil, err
// // 	}
// // 	fmt.Printf("API Response as struct %+v\n", car)

// // curl -X POST \
// // -F 'client_id=your_client_id_here' \
// // -F 'client_secret=your_client_secret_here' \
// // -F 'redirect_uri=urn:ietf:wg:oauth:2.0:oob' \
// // -F 'grant_type=client_credentials' \
// // https://mastodon.example/oauth/token
// // 	getTokenURI := fmt.Sprintf("%s/oauth/token", c.Hostname)
// // 	data = url.Values{
// // 		"client_id":     {car.ClientID},
// // 		"client_secret": {car.ClientSecret},
// // 		"redirect_uris": {"urn:ietf:wg:oauth:2.0:oob"},
// // 		"grant_type":    {"client_credentials"},
// // 	}
// // 	resp2, err := c.Client.PostForm(getTokenURI, data)
// // 	if err != nil {
// // 		return nil, err
// // 	}
// // 	if resp2.StatusCode != http.StatusOK {
// // 		return nil, fmt.Errorf("Got %d calling apps endpoint: %s", resp1.StatusCode, resp1.Status)
// // 	}
// // 	defer resp2.Body.Close()
// // 	bodyBytes, _ = io.ReadAll(resp2.Body)

// // 	bodyString = string(bodyBytes)
// // 	fmt.Println("API Response as String:\n" + bodyString)

// // 	var applicationResponse CreateApplicationResponse
// // 	err = json.Unmarshal(bodyBytes, &appCreateResp)
// // 	if err != nil {
// // 		return nil, err
// // 	}
// // 	fmt.Printf("API Response as struct %+v\n", appCreateResp)
// // 	return &appCreateResp, nil
// // }
