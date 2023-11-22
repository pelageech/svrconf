package svrconf

import (
	"io"
	"time"

	"gopkg.in/yaml.v3"
)

// ServerConfigOptional is a configuration for http.Server.
// All the fields can be absent in file configurations.
type ServerConfigOptional struct {
	Host        *string        `yaml:"host,omitempty"`
	Port        *string        `yaml:"port,omitempty"`
	ReadTimeout *time.Duration `yaml:"read_timeout,omitempty"`
}

// LoadServerConfigOptional parses yaml from the io.Reader r and returns
// a pointer to the ServerConfigOptional.
//
// If decoding is failed, it returns an error.
func LoadServerConfigOptional(r io.Reader) (*ServerConfigOptional, error) {
	var cfg ServerConfigOptional

	dec := yaml.NewDecoder(r)
	if err := dec.Decode(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
