SOURCEDIR=.
BINARY=tpl
VERSION=0.2-dev
BUILD_TIME=`date +%FT%T%z`

LDFLAGS=-ldflags "-X github.com/odedlaz/tpl/core.Version=${VERSION} -X github.com/odedlaz/tpl/core.BuildTime=${BUILD_TIME}"

.DEFAULT_GOAL: all

.PHONY: all
all: build build-alpine


.PHONY: build
build:
		go build ${LDFLAGS} -o ${SOURCEDIR}/bin/${BINARY}

.PHONY: build-alpine
build-alpine:
		CGO_ENABLED=0 go build ${LDFLAGS} -a -installsuffix cgo -o ${SOURCEDIR}/bin/${BINARY}-alpine

.PHONY: install
install:
	go install ${LDFLAGS}

.PHONY: clean
clean:
	rm -rf ${SOURCEDIR}/bin/*
