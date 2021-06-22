# Bj38 Service

This is the Bj38 service

Generated with

```
micro new --namespace=go.micro --type=service bj38
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.service.bj38
- Type: service
- Alias: bj38

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

```
# install etcd
brew install etcd

# run etcd
etcd
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./bj38-service
```

Build a docker image
```
make docker
```