package config

import (
	"encoding/json"
	"io"
	"os"
)

type DataSources struct {
	CSV  string `json:"csv"`
	JSON string `json:"json"`
}

type ExternalAPI struct {
	ClientID string `json:"client_id"`
	Secret   string `json:"secret"`
	TokenURL string `json:"token_url"`
	BaseURL  string `json:"base_url"`
}

type Config struct {
	DataSources DataSources `json:"data_source"`
	ExternalAPI ExternalAPI `json:"external_api"`
	Env         string      `json:"env"`
	Port        int         `json:"port"`
}

func New(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return readConfig(file)
}

func readConfig(r io.Reader) (*Config, error) {
	decoder := json.NewDecoder(r)
	config := Config{}
	err := decoder.Decode(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
