package svrconf_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/pelageech/svrconf"
)

var strOptionalConfig = `
host: localhost
port: 8080
`

func TestLoadServerConfigOptional(t *testing.T) {
	temp := &svrconf.ServerConfig{
		Host: "localhost",
		Port: "8080",
	}

	expected := &svrconf.ServerConfigOptional{
		Host:        &temp.Host,
		Port:        &temp.Port,
		ReadTimeout: nil,
	}

	r := strings.NewReader(strOptionalConfig)

	actual, err := svrconf.LoadServerConfigOptional(r)
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected %v but got %v", expected, actual)
	}
}
