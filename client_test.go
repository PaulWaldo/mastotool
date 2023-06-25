package mastotool

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/mattn/go-mastodon"
	"github.com/stretchr/testify/assert"
)

func TestRemoveFollowedTags(t *testing.T) {
	type args struct {
		c    mastodon.Client
		tags []*mastodon.FollowedTag
	}
	unfollowTags := []*mastodon.FollowedTag{
		{Name: "test"},
		{Name: "xyz"},
	}
	successServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pathParts := strings.Split(r.URL.Path, "/")
		tag := pathParts[4]
		expectedPath := fmt.Sprintf("/api/v1/tags/%s/unfollow", tag)
		if r.URL.Path != expectedPath {
			t.Fatalf("unexpected request path %s, expecting %s", r.URL.Path, expectedPath)
			return
		}
		_, err := w.Write(
			[]byte(fmt.Sprintf(`
			{
				"name": "%s",
				"url": "http://mastodon.example/tags/%s",
				"history": [
					{
					"day": "1668556800",
					"accounts": "0",
					"uses": "0"
					}
				],
				"following": false
			}`, tag, tag)),
		)
		if err != nil {
			t.Fatalf("writing mocked response: %s", err)
		}
	}))
	defer successServer.Close()

	failureServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
	}))
	defer successServer.Close()

	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "removes followed tags for valid tag removal request",
			args: args{
				c: mastodon.Client{
					Client: *successServer.Client(),
					Config: &mastodon.Config{Server: successServer.URL},
				},
				tags: unfollowTags},
			wantErr: nil,
		},
		{
			name: "failure returns Unathorized if not authorized",
			args: args{
				c: mastodon.Client{
					Client: *failureServer.Client(),
					Config: &mastodon.Config{Server: failureServer.URL},
				},
				tags: unfollowTags,
			},
			wantErr: errors.New("bad request: 401 Unauthorized"),
		},
		{
			name: "returns failure if no tags specified",
			args: args{
				c: mastodon.Client{
					Client: *successServer.Client(),
					Config: &mastodon.Config{Server: successServer.URL},
				},
				tags: []*mastodon.FollowedTag{}},
			wantErr: ErrEmptyFollowedTagsToRemove,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := RemoveFollowedTags(tt.args.c, tt.args.tags)
			if tt.wantErr != nil {
				assert.EqualError(t, tt.wantErr, err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
