
GOPATH:=$(shell go env GOPATH)

.PHONY: run proto build test docker-build docker-push

run:
	go run main.go

proto:
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/mutant/mutant.proto

build: proto

	go build -o mutant-srv main.go

test:
	go test -v -coverpkg=github.com/rodrigodmd/ml-mutant-srv/srv/mutant ./...

docker-build:
	docker build -t rodrigodmd/ml-mutant-srv:mutant .

docker-push: docker-build
	docker push rodrigodmd/ml-mutant-srv:mutant

