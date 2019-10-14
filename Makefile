.PHONY: build

.ONESHELL:
build:
	mkdir -p build
	cd cmd/ddns-update
	go build -v -o ../../build/ddns-update .
