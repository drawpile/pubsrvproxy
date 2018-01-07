package main

import (
	"github.com/BurntSushi/toml"
	"os"
	"log"
)

type config struct {
	Listen      string
	ServerAddr  string
	ServerHost  string
	ServerPort  int
	CacheTime   int64
	UserView    bool
	ShowUserIps bool
}

func defaultConfig() *config {
	cfg := &config{
		Listen:      "localhost:8080",
		ServerAddr:  "http://localhost:27780/",
		ServerHost:  "",
		ServerPort:  27750,
		CacheTime:   60,
		UserView:    false,
		ShowUserIps: false,
	}

	hostname, err := os.Hostname()
	if err != nil {
		log.Println("Couldn't determine hostname", err)
	} else {
		cfg.ServerHost = hostname
	}

	return cfg
}

func readConfigFile(path string) (*config, error) {
	cfg := defaultConfig()

	if _, err := toml.DecodeFile(path, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
