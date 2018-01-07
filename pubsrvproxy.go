package main

import (
	"flag"
	"log"
)

func main() {
	cfgFile := flag.String("c", "", "configuration file")
	listenAddr := flag.String("l", "", "listening address")
	serverAddr := flag.String("s", "", "server address")

	flag.Parse()

	// Load configuration file
	var cfg *config
	if len(*cfgFile) > 0 {
		var err error
		cfg, err = readConfigFile(*cfgFile)
		if err != nil {
			log.Fatal(err)
			return
		}

	} else {
		cfg = defaultConfig()
	}

	// Overridable settings
	if len(*listenAddr) > 0 {
		cfg.Listen = *listenAddr
	}

	if len(*serverAddr) > 0 {
		cfg.ServerAddr = *serverAddr
	}

	// Start
	StartServer(cfg)
}
