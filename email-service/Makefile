build:
	CGO_ENABLED=0 go build -o email-service

docker:
	docker build -t user-cli .

run:
	docker run --rm -e MICRO_REGISTRY=mdns email-service
