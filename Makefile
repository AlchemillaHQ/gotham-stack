.PHONY: test build deps

APP_NAME=gotham-stack
APP_VERSION=0.0.2
TEST_DIRS := $(shell find . -name '*_test.go' -exec dirname {} \; | sort -u)

deps:
	go mod download
	npm install

test:
ifeq ($(strip $(TEST_DIRS)),)
	@echo "No test files found."
else
	@for dir in $(TEST_DIRS); do \
		echo "Running tests in $$dir"; \
		go test $$dir; \
	done
endif

build: clean
	GOOS=linux GOARCH=amd64 go build -o dist/$(APP_NAME)_$(APP_VERSION)_linux_amd64 main.go
	GOOS=darwin GOARCH=amd64 go build -o dist/$(APP_NAME)_$(APP_VERSION)_darwin_amd64 main.go
	GOOS=windows GOARCH=amd64 go build -o dist/$(APP_NAME)_$(APP_VERSION)_windows_amd64.exe main.go
	GOOS=freebsd GOARCH=amd64 go build -o dist/$(APP_NAME)_$(APP_VERSION)_freebsd_amd64 main.go

clean:
	rm -rf dist
