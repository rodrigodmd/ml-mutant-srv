
GOPATH:=$(shell go env GOPATH)

.PHONY: proto test docker


proto:
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/stats/stats.proto
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/elastic/elastic.proto

build: proto

	go build -o stats-srv main.go

test:
	go test -v ./... -cover

docker:
	docker build -t rodrigodmd/ml-mutant-srv:stats .
	docker push rodrigodmd/ml-mutant-srv:stats
