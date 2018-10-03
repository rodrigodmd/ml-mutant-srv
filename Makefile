GOPATH:=$(shell go env GOPATH)

.PHONY: test resolve docker-build docker-push


docker-up:
	docker-compose pull
	docker-compose up -d
	echo "Waiting for server to start..."
	sleep 5

docker-down:
	docker-compose down

integration-test: docker-up
	go test ./test/... -v
	docker-compose down

resolve:
	dep ensure

docker-build:
	docker build -t rodrigodmd/ml-mutant-srv:base .
	
docker-push: docker-build
	docker push rodrigodmd/ml-mutant-srv:base