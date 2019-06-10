build:
	protoc -I. --go_out=plugins=micro:. proto/vessel/vessel.proto
	CGO_ENABLED=0 go build -o vessel-service main.go

docker:
	docker build -t vessel-service .

run:
	docker run --rm -p 50052:50051 \
		-e MICRO_SERVER_ADDRESS=:50051 \
		-e MICRO_REGISTRY=mdns \
		vessel-service
