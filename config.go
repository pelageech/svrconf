package svrconf

import (
	"io"
	"time"

	"gopkg.in/yaml.v3"
)

// ServerConfig is a configuration for http.Server.
type ServerConfig struct {
	Host        string        `yaml:"host"`
	Port        string        `yaml:"port"`
	ReadTimeout time.Duration `yaml:"read_timeout"`
}

// LoadServerConfig parses yaml from the io.Reader r and returns
// a pointer to the ServerConfig.
//
// If decoding is failed, it returns an error.
func LoadServerConfig(r io.Reader) (*ServerConfig, error) {
	var cfg ServerConfig

	dec := yaml.NewDecoder(r)
	if err := dec.Decode(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
