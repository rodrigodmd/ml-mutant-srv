# Stats Service

This is the Stats service

Generated with

```
micro new github.com/rodrigodmd/ml-mutant-srv/stats --namespace=go.micro --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.stats
- Type: srv
- Alias: stats

## Dependencies

Micro services depend on service discovery. The default is consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./stats-srv
```

Build a docker image
```
make docker
```