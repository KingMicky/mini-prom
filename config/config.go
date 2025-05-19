package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"time"
)

type Target struct {
	URL      string        `yaml:"url"`
	Interval time.Duration `yaml:"interval"`
}

type Config struct {
	Targets []Target `yaml:"targets"`
}

func LoadConfig(filename string) (*Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	return &cfg, err
}
