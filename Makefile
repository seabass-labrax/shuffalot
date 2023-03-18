# SPDX-FileCopyrightText: 2023 Sebastian Crane <seabass-labrax@gmx.com>
# SPDX-License-Identifier: AGPL-3.0-or-later

.POSIX:
.DELETE_ON_ERROR:
.SHELLFLAGS := -eu -c

PREFIX ?= /usr/local
MAN_DIR ?= $(PREFIX)/man
OUTPUT_DIR ?= _build

VERSION := $(shell git describe --tags --always --dirty)

build: $(OUTPUT_DIR)/shuffalot.1

$(OUTPUT_DIR)/shuffalot.1: shuffalot.1
	mkdir -p $(OUTPUT_DIR)
	sed "s/VERSION/$(VERSION)/" shuffalot.1 > $@

install: $(OUTPUT_DIR)/shuffalot.1
	install -D $(OUTPUT_DIR)/shuffalot.1 $(MAN_DIR)/man1/shuffalot.1

uninstall:
	rm -f $(MAN_DIR)/man1/shuffalot.1

test:
	go test -v ./...

clean:
	rm -rf $(OUTPUT_DIR)

.PHONY: test build install uninstall clean
