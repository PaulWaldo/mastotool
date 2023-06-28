package main

import (
	"context"

	"github.com/PaulWaldo/mastotool/cmd/reorg/ui"
	"github.com/mattn/go-mastodon"
)

func GetFollowedTags(prefs ui.Preferences) ([]*mastodon.FollowedTag, error) {
	var err error
	c := ui.NewClientFromPrefs(prefs)
	followedTags, err := c.GetFollowedTags(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	return followedTags, nil
}
