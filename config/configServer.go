package config

import (
	"flag"
	"log"
	"os"
)

func NewConfigServer() *ConfigServer {
	cnf := &ConfigServer{}
	return ParseFlagsServer(cnf)
}

type ConfigServer struct {
	FlagRunAddr string
}

func ParseFlagsServer(cnf *ConfigServer) *ConfigServer {

	flag.StringVar(&cnf.FlagRunAddr, "a", ":8080", "address and port to run server")

	flag.Parse()

	if envRunAddr := os.Getenv("ADDRESS"); envRunAddr != "" {
		cnf.FlagRunAddr = envRunAddr
	}
	log.Printf("Starting UpServer on %s\n", cnf.FlagRunAddr)
	return cnf
}
