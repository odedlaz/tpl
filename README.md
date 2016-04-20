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

## how to build

```bash
# build
make build
```

don't forget to `go-get`!
