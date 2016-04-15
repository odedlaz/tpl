# untem (untemplate-me)
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
$ echo 'Hello {{ "NAME" | getenv:"John" }}.' | bin/untem
$ Hello John.
$ echo 'Hello {{ "NAME" | getenv:"John" }}.' | NAME="Jane" bin/untem
$ Hello Jane.
```

can also be used with a file:

```bash
$ echo 'Hello {{ "NAME" | getenv:"John" }}.' > /tmp/john.tpl
$ bin/untem /tmp/john.tpl
$ Hello John.
```

### httpget
```bash
$ echo 'my ip is: {{ "http://api.ipify.org" | httpget }}' | bin/untem
$ my ip is: 192.0.79.33
```

## how to build

just run go build as you're used to.
don't forget to `go-get`!

```bash
rm -rf bin/* && go build -o bin/untem
```
