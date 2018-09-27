# Go Gateway Microservice

## Getting Started

If you haven't done it yet, follow the [getting started](https://github.com/rodrigodmd/ml-mutant-srv#getting-started) guide

# Development process

Run the services with docker-compose:

    cd $GOPATH/src/github.com/rodrigodmd/ml-mutant-srv/srv
    docker-compose up

Run each service manually:

### gateway

    cd $GOPATH/src/github.com/rodrigodmd/ml-mutant-srv/srv/gateway
    go run main.go

### stats

    cd $GOPATH/src/github.com/rodrigodmd/ml-mutant-srv/srv/stats
    go run main.go

### mutant

    cd $GOPATH/src/github.com/rodrigodmd/ml-mutant-srv/srv/mutant
    go run main.go


