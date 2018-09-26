
GOPATH:=$(shell go env GOPATH)

.PHONY: proto test docker


build:
	go build -o gateway-srv main.go plugin.go

test:
	go test -v ./... -cover

docker:
	docker build -t rodrigodmd/ml-mutant-srv:gateway .
	docker push rodrigodmd/ml-mutant-srv:gateway
