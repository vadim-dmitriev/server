package yaml

import (
	"log"

	"github.com/vadim-dmitriev/server/internal/config"
)

type YAMLConfig struct {
	port string
}

func NewConfig(filepath string) config.Configer {
	config := &YAMLConfig{
		port: "8000",
	}

	log.Printf("YAML config created '%#v'", config)

	return config
}

func (yc *YAMLConfig) GetPort() string {
	return yc.port
}
