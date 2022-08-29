.PHONY: build clean   help

all: build

build:
	@go build -v .

clean:
	@echo "clean"

help:
	@echo "make: compile packages and dependencies"
	@echo "make tool: run specified go tool"
	@echo "make lint: golint ./..."
	@echo "make clean: remove object files and cached files"