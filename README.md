# Go Microservice Restful API  # ML Mutant excercise [![Go Report Card](https://goreportcard.com/badge/github.com/rodrigodmd/ml-mutant-srv)](https://goreportcard.com/report/github.com/rodrigodmd/ml-mutant-srv) [![GoDoc](https://godoc.org/github.com/micro/go-api?status.svg)](https://godoc.org/github.com/rodrigodmd/ml-mutant-srv) [![Build Status](https://travis-ci.org/rodrigodmd/ml-mutant-srv.svg?branch=master)](https://travis-ci.org/rodrigodmd/ml-mutant-srv) [![codecov](https://codecov.io/gh/rodrigodmd/ml-mutant-srv/branch/master/graph/badge.svg)](https://codecov.io/gh/rodrigodmd/ml-mutant-srv)

This is an example of how to serve REST behind the API using go-restful

## Getting Started

### Run the Micro API

```
$ micro api --handler=http
```

### Run the Greeter Service

```
$ go run greeter/server/main.go
```

### Run the Greeter API

```
$ go run go-restful.go
Listening on [::]:64738
```

### Curl the API

Test the index
```
curl http://localhost:8080/greeter
{
  "message": "Hi, this is the Greeter API"
}
```

Test a resource
```
 curl http://localhost:8080/greeter/asim
{
  "msg": "Hello asim"
}
```
