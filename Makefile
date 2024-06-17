VERSION := v0.4.0
PREFIX ?= /usr/local
BINDIR ?= $(PREFIX)/bin
SHAREDIR ?= $(PREFIX)/share

build:
	@go build -o ./bin/wordgen -trimpath -ldflags="-s -w -X main.version=$(VERSION) -buildid=" .

run: build
	@./bin/wordgen

install:
	@install -Dm755 ./bin/wordgen $(DESTDIR)$(BINDIR)/wordgen
	@install -Dm644 ./LICENSE $(DESTDIR)$(SHAREDIR)/licenses/wordgen/LICENSE

uninstall:
	@rm -f $(DESTDIR)$(BINDIR)/wordgen
	@rm -rf $(DESTDIR)$(SHAREDIR)/licenses/wordgen

clean:
	@rm -f ./bin/*

prune:
	@docker system prune -a -f

release: clean build-all prune
	@./scripts/release.sh
	@./scripts/update-aur.sh

# Reproducible Builds (requires docker)
build-all: linux-amd64 linux-arm64 \
           windows-amd64 windows-arm64 \
           darwin-amd64 darwin-arm64

linux-amd64:
	@./scripts/build.sh linux amd64

linux-arm64:
	@./scripts/build.sh linux arm64

windows-amd64:
	@./scripts/build.sh windows amd64

windows-arm64:
	@./scripts/build.sh windows arm64

darwin-amd64:
	@./scripts/build.sh darwin amd64

darwin-arm64:
	@./scripts/build.sh darwin arm64

.PHONY: build run install uninstall clean prune release build-all \
        linux-amd64 linux-arm64 \
        windows-amd64 windows-arm64 \
        darwin-amd64 darwin-arm64
