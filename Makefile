# vi: ft=make
.PHONY: run proto test benchmark

test:
	go test -v ./...

run:
	go run main.go

proto:
	protoc -I cloud_storage_service/ cloud_storage_service/cloud_storage_service.proto --go_out=plugins=grpc:cloud_storage_service

benchmark:
	go test -bench=./... -benchmem -benchtime 10s

docker:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o cloud_storage .
	docker build . -t lileio/cloud_storage_service:latest
