# vi: ft=make
.PHONY: run proto test benchmark

test:
	go test -v ./...

run:
	go run cloud_storage_service/main.go

proto:
	protoc -I . cloud_storage_service.proto --go_out=plugins=grpc:$$GOPATH/src

benchmark:
	go test -bench=. -benchmem -benchtime 10s ./...

docker:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/cloud_storage ./cloud_storage_service
	docker build . -t lileio/cloud_storage_service:latest
