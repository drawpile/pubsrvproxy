package main

import (
	"github.com/BurntSushi/toml"
)

type config struct {
	Listen      string
	ServerAddr  string
	CacheTime   int64
	UserView    bool
	ShowUserIps bool
}

func defaultConfig() *config {
	return &config{
		Listen:      "localhost:8080",
		ServerAddr:  "http://localhost:27780/",
		CacheTime:   60,
		UserView:    false,
		ShowUserIps: false,
	}
}

func readConfigFile(path string) (*config, error) {
	cfg := defaultConfig()

	if _, err := toml.DecodeFile(path, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
