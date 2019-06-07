.PHONY: default build run

BIN := go-langserver-gtags
SRC := $(wildcard *.go)

default: build run

build: $(BIN)

run: build
	./$(BIN)

$(BIN): $(SRC)
	go build -o $@ $<
