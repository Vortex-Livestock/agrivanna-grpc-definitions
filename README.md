# Agrivanna gRPC Definitions

[![Production Release](https://github.com/Vortex-Livestock/agrivanna-grpc-definitions/actions/workflows/production_release.yaml/badge.svg)](https://github.com/Vortex-Livestock/agrivanna-grpc-definitions/actions/workflows/production_release.yaml)
[![Continuous Integration](https://github.com/Vortex-Livestock/agrivanna-grpc-definitions/actions/workflows/continuous_integration.yaml/badge.svg)](https://github.com/Vortex-Livestock/agrivanna-grpc-definitions/actions/workflows/continuous_integration.yaml)

## Table of Contents

- [Contributors](#contributors)
- [Description](#description)
- [Tech Stack](#tech-stack)
- [Setup](#setup)
- [Quick Start](#quick-start)
- [Database Schemas](#database-schemas)
- [Dependencies](#dependencies)

## Contributors

- [Axel Sanchez](https://github.com/Axeloooo)

## Description

Agrivanna gRPC Definitions is a repository that contains the gRPC definitions for the Agrivanna project.

## Tech Stack

![Go](https://img.shields.io/badge/Go-00ADD8.svg?style=for-the-badge&logo=Go&logoColor=white)

## Setup

The following programs should be installed:

- [Brew](https://brew.sh/)
- [Protobuf](https://grpc.io/docs/protoc-installation/)

## Quick Start

1. Install protobuff compiler

```bash
$ brew install protobuf
```

2. Clone the repository

```bash
$ git clone git@github.com:Vortex-Livestock/agrivanna-grpc-definitions.git
```

3. Change directory to the project folder

```bash
$ cd agrivanna-grpc-definitions
```

4. Install the protobuf Go plugin

```bash
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

$ export PATH="$PATH:$(go env GOPATH)/bin"
```

5. Install the Go dependencies

```bash
$ make install
```

## Other Commands

- Generate the gRPC code

```bash
$ make proto
```

- To clean the generated code

```bash
$ make clean
```
