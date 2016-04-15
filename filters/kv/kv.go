package kv

import (
	"fmt"
	"time"

	"github.com/docker/libkv"
	"github.com/docker/libkv/store"
	"github.com/docker/libkv/store/boltdb"
	"github.com/docker/libkv/store/consul"
	"github.com/docker/libkv/store/etcd"
	"github.com/docker/libkv/store/zookeeper"
	"github.com/flosch/pongo2"
	"github.com/odedlaz/untemplate-me/config"
)

func init() {
	consul.Register()
	etcd.Register()
	zookeeper.Register()
	boltdb.Register()
	pongo2.RegisterFilter("kvget", get)
}

func getSettings() *config.Settings {
	return pongo2.Globals["settings"].(*config.Settings)
}

func get(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	settings := getSettings()
	key := in.String()
	kv, err := libkv.NewStore(
		store.Backend(settings.KV.Type),
		[]string{settings.KV.URL},
		&store.Config{
			ClientTLS:         settings.KV.ClientTLS,
			TLS:               settings.KV.TLS,
			Bucket:            settings.KV.Bucket,
			Username:          settings.KV.Username,
			Password:          settings.KV.Password,
			PersistConnection: settings.KV.PersistConnection,
			ConnectionTimeout: time.Duration(settings.KV.ConnectionTimeout) * time.Second,
		},
	)

	if err != nil {
		return nil, &pongo2.Error{
			Sender:   "filter:kv",
			ErrorMsg: fmt.Sprintf("Error creating kv client: %s", err.Error()),
		}
	}

	pair, err := kv.Get(key)
	if err != nil && param.IsNil() {
		return nil, &pongo2.Error{
			Sender:   "filter:kv",
			ErrorMsg: fmt.Sprintf("Error trying accessing value at key: %v", key),
		}
	}

	if err != nil {
		return pongo2.AsValue(param.String()), nil
	}

	return pongo2.AsValue(string(pair.Value)), nil
}
