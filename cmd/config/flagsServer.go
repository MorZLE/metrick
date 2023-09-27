package config

import (
	"flag"
	"log"
	"os"
)

var FlagRunAddr string

func ParseFlags() {

	flag.StringVar(&FlagRunAddr, "a", ":8080", "address and port to run server")

	flag.Parse()

	if envRunAddr := os.Getenv("ADDRESS"); envRunAddr != "" {
		FlagRunAddr = envRunAddr
	}
	log.Printf("Starting UpServer on %s\n", FlagRunAddr)
}
