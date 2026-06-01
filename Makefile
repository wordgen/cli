PREFIX ?= /usr/local
BINDIR ?= $(PREFIX)/bin
SHAREDIR ?= $(PREFIX)/share
BIN ?= wordgen

.SILENT:

.PHONY: build run install uninstall clean

build:
	go build -o ./bin/$(BIN) ./cmd/wordgen

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
