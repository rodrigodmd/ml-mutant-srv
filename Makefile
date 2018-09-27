GOPATH:=$(shell go env GOPATH)

.PHONY: resolve docker-build docker-push

resolve:
	dep ensure

docker-build:
	docker build -t rodrigodmd/ml-mutant-srv:base .
	
docker-push: docker-build
	docker push rodrigodmd/ml-mutant-srv:base