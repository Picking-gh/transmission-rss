/*
 * Copyright (C) 2018 Aurélien Chabot <aurelien@chabot.fr>
 *
 * SPDX-License-Identifier: MIT
 */

package main

import (
	"log"
	"os"

	"github.com/go-yaml/yaml"
)

const defaultServerHost = "localhost"
const defaultServerPort = 9091
const defaultUpdateInterval = 10

// Config is handling the config parsing
type Config struct {
	Server struct {
		Host string
		Port uint16
		User string `yaml:"username"`
		Pswd string `yaml:"password"`
	}
	UpdateInterval uint64 `yaml:"update_interval"`
	Feeds          []string
}

// NewConfig return a new Config object
func NewConfig(filename string) *Config {
	var config Config
	source, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(source, &config)
	if err != nil {
		log.Fatal(err)
	}
	if config.Server.Host == "" {
		config.Server.Host = defaultServerHost
	}
	if config.Server.Port == 0 {
		config.Server.Port = defaultServerPort
	}
	if config.UpdateInterval == 0 {
		config.UpdateInterval = defaultUpdateInterval
	}
	return &config
}
