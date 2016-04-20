SOURCEDIR=.
BINARY=tpl

LDFLAGS=-ldflags "-X github.com/odedlaz/tpl/core.build=`git rev-parse HEAD`"

.DEFAULT_GOAL: all

.PHONY: all
all: build build-alpine

.PHONY: build
build:
	go build ${LDFLAGS} -o ${SOURCEDIR}/bin/${BINARY}

.PHONY: build-alpine
build-alpine:
	CGO_ENABLED=0 go build ${LDFLAGS} -a -installsuffix cgo -o ${SOURCEDIR}/bin/${BINARY}-alpine

.PHONY: test
test:
	go test ./...

.PHONY: install
install:
	go install ${LDFLAGS}

.PHONY: clean
clean:
	rm -rf ${SOURCEDIR}/bin/*
