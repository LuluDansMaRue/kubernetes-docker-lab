## Bobba-api

This is the api that is used by the bobba-vue

## Requirements

- Go

## Installation

Run the command 

```go get ./...```

## Run the project

There are 2 ways for running this projet.
Either use the ```go run```` command by runnin: 

```shell
go run api/main.go api/router.go api/routines.go
```

or by using the ```CompileDaemon``` like so

```shell
CompileDaemon -command="./api"
```

## Build the project for production mode

Run the command inside the folder ```api```

```shell
go build
```

You should then be able to run the project like this

```shell
./api
```