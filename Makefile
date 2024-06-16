VERSION := v0.2.0
PREFIX ?= /usr/local
BINDIR ?= $(PREFIX)/bin
SHAREDIR ?= $(PREFIX)/share

build:
	@go build -o bin/wordgen -trimpath -ldflags="-s -w -X main.version=$(VERSION) -buildid=" .

run: build
	@./bin/wordgen

install:
	@install -Dm755 bin/wordgen $(DESTDIR)$(BINDIR)/wordgen
	@install -Dm644 LICENSE $(DESTDIR)$(SHAREDIR)/licenses/wordgen/LICENSE

uninstall:
	@rm -f $(DESTDIR)$(BINDIR)/wordgen
	@rm -f $(DESTDIR)$(SHAREDIR)/licenses/wordgen/LICENSE

clean:
	@rm -f bin/*

prune:
	@docker system prune -a -f

test:
	@go test -v ./...

release: clean test build-all prune
	@./scripts/release.sh
	@./scripts/update-aur.sh

# Reproducible Builds (requires docker)
build-all: linux-amd64 windows-amd64 darwin-amd64

linux-amd64:
	@./scripts/build.sh linux amd64

windows-amd64:
	@./scripts/build.sh windows amd64

darwin-amd64:
	@./scripts/build.sh darwin amd64

.PHONY: build run install uninstall clean prune test release build-all \
        linux-amd64 windows-amd64 darwin-amd64
