package config

import (
	"bytes"
	"testing"
)

var r = bytes.NewBufferString(`{
	"data_source": {
		"csv": "CSV",
		"json": "JSON"
	},
	"external_api": {
		"client_id": "ClientID",
		"secret": "Secret",
		"token_url": "TokenURL",
		"base_url": "BaseURL"
	},
	"env": "Env",
	"port": 0
}`)

var config = Config{
	DataSources: DataSources{
		CSV:  "CSV",
		JSON: "JSON",
	},
	ExternalAPI: ExternalAPI{
		ClientID: "ClientID",
		Secret:   "Secret",
		TokenURL: "TokenURL",
		BaseURL:  "BaseURL",
	},
	Env:  "Env",
	Port: 0,
}

func TestReadConfig(t *testing.T) {
	t.Run("should parse config JSON into Config{}", func(t *testing.T) {
		expected := config
		actual, _ := readConfig(r)
		if *actual != expected {
			t.Errorf("actual %v, expected %v", actual, expected)
		}
	})
}
