# CHANGELOG

- [v0.2](#v02-dev)
- [v0.1](#v01)

# v0.2-dev

new features:

* pathexists filter
* stringsplit filter
* edit in place flag (similar to sed -i)

changes:

* cat now supports glob: `{{ "/path/to/dir/\*" | cat }}`

# v0.1

first release with the following filters:
* getenv
* httpget
* kvget
* cat
