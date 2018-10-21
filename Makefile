build:
	go build -o target/ddns-update -v cmd/ddns-update/ddns-update.go

build_image:
	docker build . -t rafadc/ddns-update
