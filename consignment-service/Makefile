build:
	CGO_ENABLED=0 go build -o consignment-service

.PHONY: proto
proto:
	protoc -I. --go_out=plugins=micro:. proto/consignment/consignment.proto

all: proto build

docker:
	docker build -t consignment-service .

run:
	docker run --rm -p 50051:50051 \
	 -e MICRO_SERVER_ADDRESS=:50051 \
	 -e MICRO_REGISTRY=mdns \
	 consignment-service
