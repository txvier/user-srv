# Foo Service

This is the Foo service

Generated with

```
micro new github.com/txvier/foo --namespace=go.micro --fqdn=com.txvier.srv.foo --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: com.txvier.srv.foo
- Type: srv
- Alias: foo

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend consul.

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
./foo-srv
```

Build a docker image
```
make docker
```