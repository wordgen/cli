PREFIX ?= /usr/local
BINDIR ?= $(PREFIX)/bin
SHAREDIR ?= $(PREFIX)/share
BIN ?= wordgen

.SILENT:

.PHONY: build run install uninstall clean prune release build-all \
        linux-amd64 linux-arm64 \
        windows-amd64 windows-arm64 \
        darwin-amd64 darwin-arm64

build:
	go build -o ./bin/$(BIN) -trimpath -ldflags="-s -w -buildid=" .

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

prune:
	docker system prune -a -f

release: clean build-all prune
	./.local/release.sh
	./.local/update-aur.sh

build-all: linux-amd64 linux-arm64 \
           windows-amd64 windows-arm64 \
           darwin-amd64 darwin-arm64

linux-amd64:
	./scripts/build_image.sh linux amd64

linux-arm64:
	./scripts/build_image.sh linux arm64

windows-amd64:
	./scripts/build_image.sh windows amd64

windows-arm64:
	./scripts/build_image.sh windows arm64

darwin-amd64:
	./scripts/build_image.sh darwin amd64

darwin-arm64:
	./scripts/build_image.sh darwin arm64
