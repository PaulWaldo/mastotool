package mastotool
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
