# filters

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
