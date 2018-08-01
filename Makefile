
SHELL := /bin/bash

VERSION := $(shell git describe --always)
SEMVER := "^[^0-9]*\([0-9]*\)[.]\([0-9]*\)[.]\([0-9]*\)\([0-9A-Za-z-]*\)\\$"

binary:
	go build ./cmd/mata

clean:
	rm -fr ./mata

usage:
	go run -ldflags "-X main.version=2" ./cmd/mata/mata.go

release:
	[[ $(VERSION) =~ ^[0-9]+[.][0-9]+([.][0.9]*)?$  ]] && export Version="$(VERSION)" && goreleaser --rm-dist