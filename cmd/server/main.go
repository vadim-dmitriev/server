package main

import (
	"log"

	"github.com/vadim-dmitriev/server/internal/config/yaml"
	"github.com/vadim-dmitriev/server/internal/server/tcp"
)

func main() {
	config := yaml.NewConfig("config.yaml")

	server, err := tcp.NewServer(config)
	if err != nil {
		log.Fatalf("failed to create server: %s", err)
	}

	server.Run()
}
