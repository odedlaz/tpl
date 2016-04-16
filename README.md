# tpl
A small utility that transforms templates to text.

The idea is to have a bare-bone [confd](http://confd.io) alternative, that follows the unix philosophy: _"Do One Thing and Do It Well"_.

In other words -> It just transforms template files to text, and spits the output to stdout.

[pongo2](https://github.com/flosch/pongo2) templating language was selected because it's familiar and powerful.


## why

I needed a small binary that can consume dynamic data from different sources.

confd is awesome, but it does much more than just transform templates.

plus, many times specific filters are missing and I needed a way to add new filters easily.

## filter usage

### getenv
```bash
$ echo 'Hello {{ "NAME" | getenv:"John" }}.' | bin/tpl
$ Hello John.
# can also work with default values ->
$ echo 'Hello {{ "NAME" | getenv:"John" }}.' | NAME="Jane" bin/tpl
$ Hello Jane.
```

### kv

kv support is based of the wonderful work the docker team did with [libkv](https://github.com/docker/libkv).

libkv currently supports:
* etcd
* consul
* zookeeper
* boltdb

you'll need to create the proper configuration to access a kv store.

for example:
```yaml
kv:
   url: "localhost:2379"
   type: "etcd" # etcd | consul | zk | boltdb
   connection-timeout: 10
   persistent-connection: true
```

#### kvget

```bash
$ etcdctl set /person/name John
$ echo 'Hello {{ "/person/name" | kvget:"Jane" }}.' | bin/tpl --config examples/tpl.yml
$ Hello John.
```

also works with default values via `kvget:DEFAULT`

### httpget
```bash
$ echo 'my ip is: {{ "http://api.ipify.org" | httpget }}' | bin/tpl
$ my ip is: 192.0.79.33
```

### cat

```bash
$ echo 'tpl version: {{ "GOPATH" | getenv | stringformat: "%s/src/github.com/odedlaz/tpl/VERSION" | cat }}' | bin/tpl
# or with variable substitution -
$ echo "tpl version: {{ \"$GOPATH/src/github.com/odedlaz/tpl/VERSION\" | cat }}" | bin/tpl
$ tpl version: 0.1
```

also works with default values via `cat:DEFAULT`

## how to build

just run go build as you're used to.
don't forget to `go-get`!

```bash
rm -rf bin/* && go build -o bin/tpl
```
