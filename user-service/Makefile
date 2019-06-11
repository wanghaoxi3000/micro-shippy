build:
	CGO_ENABLED=0 go build -o user-service

.PHONY: proto
proto:
	protoc -I. --go_out=plugins=micro:. proto/user/user.proto

all: proto build

docker:
	docker build -t user-service .

run:
	docker run --rm -p 50053:50051 \
		-e MICRO_SERVER_ADDRESS=:50051 \
		-e MICRO_REGISTRY=mdns \
		user-service
