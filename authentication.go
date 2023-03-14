package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Config struct {
	Hostname      string
	AppWebsiteURI string
}

type applicationResponse struct {
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

func (c *Config) CreateApplication() error {
	uri := fmt.Sprintf("%s/api/v1/apps", c.Hostname)
	resp, err := http.Post(uri, "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Got %d calling apps endpoint: %s", resp.StatusCode, resp.Status)
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	bodyString := string(bodyBytes)
	fmt.Println("API Response as String:\n" + bodyString)

	var applicationResponse applicationResponse
	err = json.Unmarshal(bodyBytes, &applicationResponse)
	if err != nil {
		return err
	}
	fmt.Printf("API Response as struct %+v\n", applicationResponse)
	return nil
}
