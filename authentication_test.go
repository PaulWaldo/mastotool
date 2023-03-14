package main

import "testing"

func TestConfig_createApplication(t *testing.T) {
	type fields struct {
		Hostname      string
		AppWebsiteURI string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				Hostname:      tt.fields.Hostname,
				AppWebsiteURI: tt.fields.AppWebsiteURI,
			}
			if err := c.CreateApplication(); (err != nil) != tt.wantErr {
				t.Errorf("Config.createApplication() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
