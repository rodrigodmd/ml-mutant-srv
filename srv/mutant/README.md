# Mutant Service

This is the Mutant service

Generated with

```
micro new github.com/rodrigodmd/mutant --namespace=go.micro --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.mutant
- Type: srv
- Alias: mutant

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
./mutant-srv
```

Build a docker image
```
make docker
```