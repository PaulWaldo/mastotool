package ui

// func Test_removeFollowedTags_CallsUnfollowForAllTags(t *testing.T) {
// 	a := test.NewApp()
// 	w := a.NewWindow("")
// 	ui := FollowedTagsUI{RemoveTags: []*mastodon.FollowedTag{{Name: "test1"}, {Name: "test2"}}}
// 	w.SetContent(ui.MakeFollowedTagsUI())
// 	w.Resize(fyne.Size{Width: 400, Height: 400})

// 	successServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		pathParts := strings.Split(r.URL.Path, "/")
// 		tag := pathParts[4]
// 		expectedPath := fmt.Sprintf("/api/v1/tags/%s/unfollow", tag)
// 		if r.URL.Path != expectedPath {
// 			t.Fatalf("unexpected request path %s, expecting %s", r.URL.Path, expectedPath)
// 			return
// 		}
// 		_, err := w.Write(
// 			[]byte(fmt.Sprintf(`
// 			{
// 				"name": "%s",
// 				"url": "http://mastodon.example/tags/%s",
// 				"history": [
// 					{
// 						"day": "1668556800",
// 						"accounts": "0",
// 						"uses": "0"
// 					}
// 					],
// 					"following": false
// 					}`, tag, tag)),
// 		)
// 		if err != nil {
// 			t.Fatalf("writing mocked response: %s", err)
// 		}
// 	}))
// 	defer successServer.Close()

// 	myApp := &myApp{}
// 	myApp.ftui = ui
// 	host := successServer.URL
// 	token := ""
// 	clientId := ""
// 	clientSecret := ""
// 	myApp.prefs.MastodonServer = binding.BindString(&host)
// 	myApp.prefs.AccessToken = binding.BindString(&token)
// 	myApp.prefs.ClientID = binding.BindString(&clientId)
// 	myApp.prefs.ClientSecret = binding.BindString(&clientSecret)
// 	err := myApp.RemoveFollowedTags()
// 	assert.NoError(t, err)
// }
