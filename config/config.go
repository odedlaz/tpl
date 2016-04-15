package config

import (
	"crypto/tls"
	"io/ioutil"

	libkv "github.com/docker/libkv/store"
	"github.com/ghodss/yaml"
)

type Settings struct {
	KV struct {
		URL               string                 `yaml:"url"`
		Type              string                 `yaml:"type"`
		ClientTLS         *libkv.ClientTLSConfig `yaml:"client-tls"`
		TLS               *tls.Config            `yaml:"tls"`
		PersistConnection bool                   `yaml:"persistent-connection"`
		ConnectionTimeout int                    `yaml:"connection-timeout"`
		Bucket            string                 `yaml:"bucket"`
		Username          string                 `yaml:"username"`
		Password          string                 `yaml:"password"`
	}
}

func Load(filepath string) (*Settings, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return &Settings{}, err
	}
	settings := Settings{}
	if err = yaml.Unmarshal([]byte(data), &settings); err != nil {
		return &Settings{}, err
	}
	return &settings, nil
}
