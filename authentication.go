package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type createApplicationClient struct {
	Hostname      string
	AppWebsiteURI string
	Client        *http.Client
}

func NewCreateApplicationClient(hostname string) *createApplicationClient {
	return &createApplicationClient{
		Hostname: hostname,
		Client:   &http.Client{},
	}
}

type Config struct {
	Hostname      string
	AppWebsiteURI string
}

type CreateApplicationResponse struct {
	ID           string `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	Website      string `json:"website,omitempty"`
	RedirectURI  string `json:"redirect_uri,omitempty"`
	ClientID     string `json:"client_id,omitempty"`
	ClientSecret string `json:"client_secret,omitempty"`
	VapidKey     string `json:"vapid_key,omitempty"`
}

// curl -X POST \
// 	-F 'client_name=Test Application' \
// 	-F 'redirect_uris=urn:ietf:wg:oauth:2.0:oob' \
// 	-F 'scopes=read write push' \
// 	-F 'website=https://myapp.example' \
// 	https://mastodon.example/api/v1/apps

func (c *createApplicationClient) CreateApplication() (*CreateApplicationResponse, error) {
	uri := fmt.Sprintf("%s/api/v1/apps", c.Hostname)
	data := url.Values{
		"client_name":   {"Test Application"},
		"redirect_uris": {"urn:ietf:wg:oauth:2.0:oob"},
		"scopes":        {"read write push"},
		"website":       {"https://myapp.example"},
	}
	resp, err := c.Client.PostForm(uri, data)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Got %d calling apps endpoint: %s", resp.StatusCode, resp.Status)
	}
	defer resp.Body.Close()
	bodyBytes, _ := io.ReadAll(resp.Body)

	bodyString := string(bodyBytes)
	fmt.Println("API Response as String:\n" + bodyString)

	var applicationResponse CreateApplicationResponse
	err = json.Unmarshal(bodyBytes, &applicationResponse)
	if err != nil {
		return nil, err
	}
	fmt.Printf("API Response as struct %+v\n", applicationResponse)
	return &applicationResponse, nil
}
