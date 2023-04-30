package mastotool

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	"github.com/mattn/go-mastodon"
	"github.com/stretchr/testify/assert"
)

// package main

// // func Test_aa(t *testing.T) {
// // 	server := httptest.NewServer(http.HandlerFunc(
// // 		func(resp http.ResponseWriter, req *http.Request) {
// // 			if req.URL.Path != "/api/v1/apps" {
// // 				t.Errorf("unexpected request path %s",
// // 					req.URL.Path)
// // 				return
// // 			}
// // 			result := createApplicationResponse{
// // 				ID:           "563419",
// // 				Name:         "test app",
// // 				Website:      "",
// // 				RedirectURI:  "urn:ietf:wg:oauth:2.0:oob",
// // 				ClientID:     "client_id",
// // 				ClientSecret: "client_secret",
// // 				VapidKey:     "vapid_key",
// // 			}
// // 			err := json.NewEncoder(resp).Encode(&result)
// // 			if err != nil {
// // 				t.Fatal(err)
// // 				return
// // 			}
// // 		},
// // 	))
// // 	defer server.Close()
// // 	client := NewCreateApplicationClient(server.URL)
// // 	resp, err := client.CreateApplication()
// // 	if err != nil {
// // 		t.Fatal(err)
// // 	}
// // 	assert.Equal(t, resp.ClientID, "client_id")
// // 	assert.Equal(t, resp.ClientSecret, "client_secret")

// // }

// // func TestConfig_createApplication(t *testing.T) {
// // 	type fields struct {
// // 		Hostname      string
// // 		AppWebsiteURI string
// // 	}
// // 	tests := []struct {
// // 		name    string
// // 		fields  fields
// // 		wantErr bool
// // 	}{}
// // 	for _, tt := range tests {
// // 		t.Run(tt.name, func(t *testing.T) {
// // 			c := &Config{
// // 				Hostname:      tt.fields.Hostname,
// // 				AppWebsiteURI: tt.fields.AppWebsiteURI,
// // 			}
// // 			if err := c.CreateApplication(); (err != nil) != tt.wantErr {
// // 				t.Errorf("Config.createApplication() error = %v, wantErr %v", err, tt.wantErr)
// // 			}
// // 		})
// // 	}
// // }

func TestAuthenticationURL(t *testing.T) {
	type args struct {
		appConfig *mastodon.AppConfig
	}
	successServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(
			`{
		  "id": "563419",
		  "name": "test app",
		  "website": null,
		  "redirect_uri": "urn:ietf:wg:oauth:2.0:oob",
		  "client_id": "TWhM-tNSuncnqN7DBJmoyeLnk6K3iJJ71KKXxgL1hPM",
		  "client_secret": "ZEaFUFmF0umgBX1qKJDjaU99Q31lDkOU8NutzTOoliw",
		  "vapid_key": "BCk-QqERU0q-CfYZjcuB6lnyyOYfJ2AifKqfeGIm7Z-HiTU5T9eTG5GxVA0_OH5mMlI4UkkDTpaZwozy0TzdZ2M="
		}`))
		if err != nil {
			t.Fatalf("writing mocked response: %s", err)
		}
	}))
	goodURL, err := url.Parse(fmt.Sprintf("%s/oauth/authorize?client_id=TWhM-tNSuncnqN7DBJmoyeLnk6K3iJJ71KKXxgL1hPM&redirect_uri=urn%%3Aietf%%3Awg%%3Aoauth%%3A2.0%%3Aoob&response_type=code&scope=", successServer.URL))
	if err != nil {
		t.Fatalf("processing expected URL: %s", err)
	}
	tests := []struct {
		name    string
		args    args
		want    *url.URL
		wantErr bool
	}{
		{
			name: "success",
			args: args{appConfig: &mastodon.AppConfig{
				Server: successServer.URL,
			}},
			want:    goodURL,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AuthenticationURL(tt.args.appConfig)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthenticationURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthenticationURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAuthenticationConfig(t *testing.T) {
	server := "https://myserver"
	got := NewAuthenticationConfig(server)
	assert.Equal(t, server, got.Server, "Expecting server to be %s, but got %s", server, got.Server)
	assert.Equal(t, website, got.Website, "Expecting website to be %s, but got %s", website, got.Website)
}
