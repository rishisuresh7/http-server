# http-server
## Basic Http server to interact with a GRPC server

## API Documentation: [api](api)

### Prerequisites
    - make
    - go

###### Config Usage(to be set at env)

- `PORT` -  Port number the application will start on
- `TOKEN` - Token for proxy authentication
- `GRPC_URI` - Host and port for GRPC server

#### Clean binaries
```shell
$ make clean
```

#### Run all tests
```shell
$ make test
```

#### Build application
```shell
$ make build
```

#### Run application with default config
```shell
$ make run
```

#### Build application for linux
```shell
$ make build-linux
```
