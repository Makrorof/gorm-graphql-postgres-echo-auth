package configs

import (
	"errors"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	PostgresConf PostgresConf `yaml:"db"`
	LogConf      LogConf      `yaml:"log"`
	APIConf      APIConf      `yaml:"api"`
}

func readConf(filename string) (*Config, error) {
	buf, err := os.ReadFile(filename)
	if err != nil {
		if _, err := createDefaultConf(filename); err != nil {
			return nil, err
		}

		log.Println("[SERVER] A new configuration file has been created. Please configure your settings.")
		fmt.Println("[SERVER] A new configuration file has been created. Please configure your settings.")
		os.Exit(0)

		return nil, errors.New("A new configuration file has been created. Please configure your settings.")
	}

	c := &Config{}
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %w", filename, err)
	}

	return c, err
}

func createDefaultConf(filename string) (*Config, error) {
	c := &Config{
		PostgresConf: PostgresConf{Host: "localhost", Port: 5432, Username: "postgres", Password: "postgres", Name: "postgres"},
		LogConf:      LogConf{Path: "/var/log/gorm_graphql_postgres/", Level: 4},
		APIConf:      APIConf{Port: 3131},
	}

	data, err := yaml.Marshal(c)

	if file, err := os.Create("config.yaml"); err == nil {
		defer file.Close()
		if _, err := file.Write(data); err != nil {
			return nil, err
		}
	}

	return c, err
}

func ReadConf(filename string) (*Config, error) {
	if cfg, err := readConf(filename); err != nil {
		return nil, err
	} else {
		return cfg, nil
	}
}
