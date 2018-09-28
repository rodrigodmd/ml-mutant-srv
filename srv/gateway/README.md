# Go Gateway Microservice

## Getting Started

If you haven't done it yet, follow the [getting started](https://github.com/rodrigodmd/ml-mutant-srv#getting-started) guide

# Development process

We are going to start all the process, except for the one/s we want to work on. Comment out the gateway service from:
*$GOPATH/src/github.com/rodrigodmd/ml-mutant-srv*
```yml
gateway:
    image: rodrigodmd/ml-mutant-srv:gateway
    build: ./srv/gateway
    command: --registry_address=consul:8500 --register_interval=5 --register_ttl=10
    links:
      - consul
```
Start services with docker-compose:

    cd $GOPATH/src/github.com/rodrigodmd/ml-mutant-srv
    docker-compose up

Run the service with go:

    cd srv/gateway
    make run
