PREFIX ?= /usr/local
BINDIR ?= $(PREFIX)/bin
SHAREDIR ?= $(PREFIX)/share
BIN ?= wordgen
BUILD_VERSION := $(shell git describe --tags --always --dirty 2>/dev/null || cat VERSION)

.SILENT:

.PHONY: build run install uninstall clean

build:
	go build -o ./bin/$(BIN) -ldflags="-X main.version=$(BUILD_VERSION)" ./cmd/wordgen

run: build
	./bin/$(BIN)

install:
	install -Dm755 ./bin/$(BIN) $(BINDIR)/$(BIN)
	install -Dm644 ./LICENSE $(SHAREDIR)/licenses/$(BIN)/LICENSE

uninstall:
	rm -f $(BINDIR)/$(BIN)
	rm -rf $(SHAREDIR)/licenses/$(BIN)

clean:
	rm -rf ./bin/*
