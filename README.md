# <img src="https://uploads-ssl.webflow.com/5ea5d3315186cf5ec60c3ee4/5edf1c94ce4c859f2b188094_logo.svg" alt="Pip.Services Logo" width="200"> <br/> GRPC Components for Pip.Services in Go

This module is a part of the [Pip.Services](http://pipservices.org) polyglot microservices toolkit.

The grpc module is used to organize synchronous data exchange using calls through the gRPC protocol. It has implementations of both the server and client parts.

The module contains the following packages:

- [**Build**](https://godoc.org/github.com/pip-services3-gox/pip-services3-grpc-gox/build) - factories for creating gRPC services
- [**Clients**](https://godoc.org/github.com/pip-services3-gox/pip-services3-grpc-gox/clients) - basic client components that use the gRPC protocol and Commandable pattern through gRPC
- [**Services**](https://godoc.org/github.com/pip-services3-gox/pip-services3-grpc-gox/services) - basic service implementations for connecting via the gRPC protocol and using the Commandable pattern via gRPC

<a name="links"></a> Quick links:

* [Configuration](http://docs.pipservices.org/conceptual/configuration/component_configuration/)
* [Protocol buffer](https://github.com/pip-services3-gox/pip-services3-grpc-gox/blob/master/protos/commandable.proto)
* [API Reference](https://godoc.org/github.com/pip-services3-gox/pip-services3-grpc-gox/)
* [Change Log](CHANGELOG.md)
* [Get Help](http://docs.pipservices.org/get_help/)
* [Contribute](http://docs.pipservices.org/contribute/)


## Use

Get the package from the Github repository:
```bash
go get -u github.com/pip-services3-gox/pip-services3-grpc-gox@latest
```

## Develop

For development you shall install the following prerequisites:
* Golang v1.18+
* Visual Studio Code or another IDE of your choice
* Docker
* Git

Run automated tests:
```bash
go test -v ./test/...
```

Generate API documentation:
```bash
./docgen.ps1
```

Before committing changes run dockerized test as:
```bash
./test.ps1
./clear.ps1
```

## Contacts

The library is created and maintained by **Sergey Seroukhov** and **Levichev Dmitry**.

The documentation is written by:
- **Levichev Dmitry**
