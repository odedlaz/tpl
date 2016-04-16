package KeyValueStoreFilter

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
	"github.com/odedlaz/tpl/template"
)

func init() {
	consul.Register()
	etcd.Register()
	zookeeper.Register()
	boltdb.Register()
	template.RegisterFilter("kvget", get)
}

func get(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	key := in.String()
	kv, err := libkv.NewStore(
		store.Backend(template.Settings.KV.Type),
		[]string{template.Settings.KV.URL},
		&store.Config{
			ClientTLS:         template.Settings.KV.ClientTLS,
			TLS:               template.Settings.KV.TLS,
			Bucket:            template.Settings.KV.Bucket,
			Username:          template.Settings.KV.Username,
			Password:          template.Settings.KV.Password,
			PersistConnection: template.Settings.KV.PersistConnection,
			ConnectionTimeout: time.Duration(template.Settings.KV.ConnectionTimeout) * time.Second,
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
