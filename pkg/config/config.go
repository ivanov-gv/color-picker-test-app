package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"path/filepath"
)

type Config struct {
	Http HTTP `yaml:"http"`
	Pg   PG   `yaml:"pg"`
}

type PG struct {
	DbConnString string `yaml:"db_conn_string"`
}

type HTTP struct {
	AppPort string `yaml:"port"`
}

func Parse(confPath string) (*Config, error) {
	filename, err := filepath.Abs(confPath)
	if err != nil {
		return nil, fmt.Errorf("can't get config path: %w", err)
	}
	yamlConf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("can't read conf: %w", err)
	}

	var config Config
	err = yaml.Unmarshal(yamlConf, &config)
	if err != nil {
		return nil, fmt.Errorf("can't unmarshall conf: %w", err)
	}

	return &config, nil
}
