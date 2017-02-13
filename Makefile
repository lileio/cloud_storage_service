# vi: ft=make
.PHONY: run proto test benchmark
run:
	go run main.go

proto:
	protoc -I cloud_storage_service/ cloud_storage_service/cloud_storage_service.proto --go_out=plugins=grpc:cloud_storage_service

test:
	go test -v ./...

benchmark:
	go test -bench=./... -benchmem -benchtime 10s
