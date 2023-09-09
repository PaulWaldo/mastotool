package app

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
			name: "returns valid authentication URL",
			args: args{appConfig: &mastodon.AppConfig{
				Server: successServer.URL,
			}},
			want:    goodURL,
			wantErr: false,
		},
	}
	t.Parallel()
	for _, tt := range tests {
		tt := tt
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

func TestNewAuthenticationConfig_ProperlyPopulatesStructure(t *testing.T) {
	t.Parallel()
	server := "https://myserver"
	got := NewAuthenticationConfig(server)
	assert.Equal(t, server, got.Server, "Expecting server to be %s, but got %s", server, got.Server)
	assert.Equal(t, website, got.Website, "Expecting website to be %s, but got %s", website, got.Website)
}
