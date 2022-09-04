.PHONY: build clean help run

all: build

build:clean
	@cd cmd; go build -o ../potato -v .

clean:
	@echo "clean"
	@rm -f potato potato.sqlite3

run:
	@./potato

help:
	@echo "make: compile packages and dependencies"
	@echo "make tool: run specified go tool"
	@echo "make lint: golint ./..."
	@echo "make clean: remove object files and cached files"
