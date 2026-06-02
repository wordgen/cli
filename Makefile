BIN := wordgen

PREFIX ?= $(HOME)/.local
BINDIR ?= $(PREFIX)/bin
SHAREDIR ?= $(PREFIX)/share

BUILD_VERSION := $(shell git describe --tags --always --dirty 2>/dev/null || cat VERSION)

.PHONY: build run install uninstall clean

build:
	@go build -o "./bin/$(BIN)" -ldflags="-X main.version=$(BUILD_VERSION)" ./cmd/wordgen

run: build
	@"./bin/$(BIN)"

install:
	install -Dm755 "./bin/$(BIN)" -t "$(DESTDIR)$(BINDIR)/"
	install -Dm644 ./LICENSE -t "$(DESTDIR)$(SHAREDIR)/licenses/$(BIN)/"

uninstall:
	rm -f "$(DESTDIR)$(BINDIR)/$(BIN)"
	rm -rf "$(DESTDIR)$(SHAREDIR)/licenses/$(BIN)"

clean:
	rm -rf ./bin
