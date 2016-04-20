# tpl
A small utility that transforms templates to text.

The idea is to have a bare-bone [confd](http://confd.io) alternative, that follows the unix philosophy: _"Do One Thing and Do It Well"_.

In other words -> It just transforms template files to text, and spits the output to stdout.

[pongo2](https://github.com/flosch/pongo2) templating language was selected because it's familiar and powerful.


## why

I needed a small binary that can consume dynamic data from different sources.

confd is awesome, but it does much more than just transform templates.

plus, many times specific filters are missing and I needed a way to add new filters easily.

## how

```bash
# pipe your template
$ echo 'Hello {{ "NAME" | getenv }}.' | bin/tpl
# or get it from a file
$ bin/tpl /path/to/template/file
```

### configuration

some filters might require a config file.
tpl will default to `tpl.yml`, and if it doesn't exit - will ignore it.

example configuration:

```yaml
kv:
   url: "localhost:2379"
   type: "etcd" # etcd | consul | zk | boltdb
   bucket: mybucket
   connection-timeout: 10
   persistent-connection: true
```

* currently, only ``kv`` filter requires a config file.

## examples

### Elasticsearch Dockerfile

what if you want to create several Dockerfile(s) for different Elasticsearch versions, and even specific plugins?


```bash
FROM elasticsearch:{{"VERSION" | getenv:"latest" }}
MAINTAINER Oded odedlaz@gmail.com

{% if "PLUGINS" | getenv:"" != "" %}
# install all the plugins
{% for plugin in "PLUGINS" | getenv | stringsplit:"," %}
RUN usr/share/elasticsearch/bin/plugin install {{ plugin }}
{% endfor %}
{% endif %}
```

now run it!

```bash
# without any arguments
$ bin/tpl /path/to/Dockerfile.tpl
FROM elasticsearch:latest
MAINTAINER Oded odedlaz@gmail.com

# with VERSION env variable
$ VERSION="1.7" bin/tpl /path/to/Dockerfile.tpl
FROM elasticsearch:1.7
MAINTAINER Oded odedlaz@gmail.com

# with the kopf and marvel plugins
$ VERSION="1.7" PLUGINS="kopf,marvel" bin/tpl /path/to/Dockerfile.tpl
FROM elasticsearch:1.7
MAINTAINER Oded odedlaz@gmail.com

# install all the plugins
RUN usr/share/elasticsearch/bin/plugin install kopf
RUN usr/share/elasticsearch/bin/plugin install marvel
```

## filters

pongo2 currently supports all [django 1.7 filters](https://docs.djangoproject.com/en/1.7/ref/templates/builtins/).


### splitstring
```bash
$ echo '{% for digit in "0,1,1,2,3,5,8" | stringsplit:"," %}{{digit}}\n{% endfor %}\n' | bin/tpl
0
1
1
2
3
5
8
```

### getenv
```bash
$ echo 'Hello {{ "NAME" | getenv:"John" }}.' | NAME="Jane" bin/tpl
Hello Jane.
# can also work with default values ->
$ echo 'Hello {{ "NAME" | getenv:"John" }}.' | bin/tpl
Hello John.
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
Hello John.
```

also works with default values via `kvget:DEFAULT`

### httpget
```bash
$ echo 'my ip is: {{ "http://api.ipify.org" | httpget }}' | bin/tpl
my ip is: 192.0.79.33
```

### pathexists
```bash
$ tmpfile=$(mktemp)
$ echo 'does the file exist? {% if "tmpfile" | getenv | pathexists %}yes{% else %}no{% endif %}' | bin/tpl
does the file exist? yes
```

### cat

```bash
$ echo 'tpl version: {{ "GOPATH" | getenv | stringformat: "%s/src/github.com/odedlaz/tpl/VERSION" | cat }}' | bin/tpl
# or with variable substitution -
$ echo "tpl version: {{ \"$GOPATH/src/github.com/odedlaz/tpl/VERSION\" | cat }}" | bin/tpl
tpl version: 0.1
```

also works with default values via `cat:DEFAULT`

## how to build

just run go build as you're used to.
don't forget to `go-get`!

```bash
# build for your architecture
go build -o bin/tpl
# build for alpine
CGO_ENABLED=0 go build -a -installsuffix cgo -o bin/tpl
```

I'll add a Makefile soon.
