package config

import (
	log "Challenge/internal/adapters/Logger"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
)

// DatabaseConfigurations database configurations

type DatabaseConfigurations struct {
	Dsn string `koanf:"dsn"`
}

//Kafka configuration

type KafkaConfiguration struct {
	Brokers  []string `koanf:"broker"`
	Topic    string   `koanf:"topic"`
	MaxBytes int      `koanf:"maxBytes"`
}

// Configurations Application wide configurations
type Configurations struct {
	Database DatabaseConfigurations `koanf:"database"`
	Kafka    KafkaConfiguration     `koanf:"kafka"`
}

func LoadConfig() *Configurations {
	k := koanf.New(".")
	err := k.Load(file.Provider("./resources/environment/configmap.yaml"), yaml.Parser())
	if err != nil {
		log.Error.Println("Failed to locate configurations. %v ", err)
	}
	var configuration Configurations
	err = k.Unmarshal("", &configuration)
	if err != nil {
		log.Error.Println("Failed to load configurations. %v", err)
	}
	return &configuration
}
