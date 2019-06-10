build:
	CGO_ENABLED=0 go build -o consignment-cli cli.go

docker:
	docker build -t consignment-cli .

run:
	docker run --rm -e MICRO_REGISTRY=mdns consignment-cli