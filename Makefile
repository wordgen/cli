PREFIX ?= /usr/local
BINDIR ?= $(PREFIX)/bin
SHAREDIR ?= $(PREFIX)/share

build:
	@go build -o ./bin/wordgen -trimpath -ldflags="-s -w -buildid=" .

run: build
	@./bin/wordgen

install:
	@install -Dm755 ./bin/wordgen $(DESTDIR)$(BINDIR)/wordgen
	@install -Dm644 ./LICENSE $(DESTDIR)$(SHAREDIR)/licenses/wordgen/LICENSE

uninstall:
	@rm -f $(DESTDIR)$(BINDIR)/wordgen
	@rm -rf $(DESTDIR)$(SHAREDIR)/licenses/wordgen

clean:
	@rm -rf ./bin/*

prune:
	@docker system prune -a -f

release: clean build-all prune
	@./.local/release.sh
	@./.local/update-aur.sh

# Reproducible Builds (requires docker)
build-all: linux-amd64 linux-arm64 \
           windows-amd64 windows-arm64 \
           darwin-amd64 darwin-arm64

linux-amd64:
	@./scripts/build_image.sh linux amd64

linux-arm64:
	@./scripts/build_image.sh linux arm64

windows-amd64:
	@./scripts/build_image.sh windows amd64

windows-arm64:
	@./scripts/build_image.sh windows arm64

darwin-amd64:
	@./scripts/build_image.sh darwin amd64

darwin-arm64:
	@./scripts/build_image.sh darwin arm64

.PHONY: build run install uninstall clean prune release build-all \
        linux-amd64 linux-arm64 \
        windows-amd64 windows-arm64 \
        darwin-amd64 darwin-arm64
