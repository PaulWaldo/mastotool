package mastotool

import (
	"context"
	"errors"

	"github.com/mattn/go-mastodon"
)

// ErrEmptyFollowedTagsToRemove is for the case where an empty list of followed tags are requested to be removed
var ErrEmptyFollowedTagsToRemove = errors.New("removing followed tags: empty list of tags")

// RemoveFollowedTags removes the given list of tags from the user's following list.
// These tags will no longer show up in the user's feed
func RemoveFollowedTags(c mastodon.Client, tags []*mastodon.FollowedTag) error {
	if len(tags) == 0 {
		return ErrEmptyFollowedTagsToRemove
	}

	for _, tag := range tags {
		_, err := c.TagUnfollow(context.Background(), tag.Name)
		if err != nil {
			return err
		}
	}
	return nil
}
