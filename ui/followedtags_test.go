package ui

import (
	"testing"

	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/test"
	"github.com/mattn/go-mastodon"
	"github.com/stretchr/testify/assert"
)

func Test_a(t *testing.T) {
	tags := []*mastodon.FollowedTag{
		{Name: "name1"},
		{Name: "name2"},
		{Name: "name3"},
	}
	inter := make([]interface{}, len(tags))

	for i, v := range tags {
		inter[i] = v
	}
	bl := binding.NewUntypedList()
	err := bl.Set(inter)
	assert.NoError(t, err)
	a := test.NewApp()
	list := NewBoundList(bl)
	w := a.NewWindow("")
	w.SetContent(list)
	assert.Equal(t, len(tags), list.Length())
}
