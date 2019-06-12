build:
	CGO_ENABLED=0 go build -o user-cli

docker:
	docker build -t user-cli .

run:
	docker run --rm -e MICRO_REGISTRY=mdns user-cli
