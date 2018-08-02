
SHELL := /bin/bash

VERSION := $(shell git describe --always)
SEMVER := "^[^0-9]*\([0-9]*\)[.]\([0-9]*\)[.]\([0-9]*\)\([0-9A-Za-z-]*\)\\$"

binary:
	go build ./cmd/mata

clean:
	rm -fr ./mata

release:
	[[ $(VERSION) =~ ^[0-9]+[.][0-9]+([.][0.9]*)?$  ]] && export Version="$(VERSION)" && goreleaser --rm-dist
