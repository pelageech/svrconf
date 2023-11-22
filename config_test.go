package svrconf_test

import (
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/pelageech/svrconf"
)

var strConfig = `
host: localhost
port: 8080
read_timeout: 54s
`

func TestLoadServerConfig(t *testing.T) {
	expected := &svrconf.ServerConfig{
		Host:        "localhost",
		Port:        "8080",
		ReadTimeout: 54 * time.Second,
	}

	r := strings.NewReader(strConfig)

	actual, err := svrconf.LoadServerConfig(r)
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected %v but got %v", expected, actual)
	}
}
