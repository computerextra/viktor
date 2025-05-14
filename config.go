package main

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type (
	Config struct {
		Version  float32
		Database database
		Folder   folder
		Mail     mail
		Sage     sage
	}

	database struct {
		Url string
	}

	folder struct {
		Upload  string
		Archive string
	}

	mail struct {
		From     string
		Server   string
		Port     int
		User     string
		Password string
	}

	sage struct {
		Server   string
		Port     int
		User     string
		Password string
		Database string
	}
)

func NewConfig() *Config {
	err := os.WriteFile("test.toml", configFile, 0644)
	if err != nil {
		panic(fmt.Errorf("failed to create config: %v", err))

	}

	var conf Config
	_, err = toml.DecodeFile("test.toml", &conf)
	if err != nil {
		panic(fmt.Errorf("failed to decode config.toml: %v", err))
	}
	err = os.Remove("test.toml")
	if err != nil {
		panic(fmt.Errorf("failed to delete config file: %v", err))
	}

	return &conf
}
