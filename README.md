# Vibes RESTful API

RESTful API used to support Vibes app. Written in [Go](https://golang.org/).

# Quick Start

## Prerequisites

Define an .env file at the root directory of the project containing all the environment variables needed. You can find the keys needed for the env vars key-value pairs in the [configuration](https://github.com/electronlabs/vibes-api/blob/develop/config/config.go) file.

## Using Docker

You can use Docker to start the app locally. The [Dockerfile](https://github.com/electronlabs/vibes-api/blob/develop/Dockerfile) and the [docker-compose.yml](https://github.com/electronlabs/vibes-api/blob/develop/docker-compose.yml) are already provided for you. Navigate to project root folder and run the following command to start the server:

```
docker-compose up
```

## Using Go Tool

You will need to [download and install Go](https://golang.org/doc/install) in order to use this method. Navigate to project root folder and run the following commands to start the server:

```
go build -o api
./api
```

# Testing

## Run Tests

### Using GoConvey

Navigate to project root folder and run:

```
goconvey
```

GoConvey will run your tests and display the results in a web UI.

### Using Go Test

Navigate to project root folder and run:

```
go test ./...
```
This will run tests in all sub-folders and display the results in the console output.

## Write Tests

### Mocking

You can auto-generate mocks for your interfaces using [mockery](https://github.com/vektra/mockery) tool. Navigate to the folder that contains the interface and run:

```
mockery -name=InterfaceToMock
```

Then, you can specify expectations and verify calls are happening using Testify's [mock](https://github.com/stretchr/testify#mock-package) package.

### Assertions

GoConvey comes with a lot standard [assertions](https://github.com/smartystreets/goconvey/wiki/Assertions) you can use with `So()`.

# Architecture Overview

The app is designed to use a layered architecture. The architecture is heavily influenced by the [Clean Architecture](http://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html). You can read the [Architecture Overview](https://github.com/electronlabs/vibes-api/blob/develop/Architecture.md) document for more details.

# Routing

[Gin](https://github.com/gin-gonic/gin) web framework is used to handle routing.