# Agrivanna gRPC Definitions

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

1. Clone the repository

```bash
$ git clone git@github.com:Vortex-Livestock/agrivanna-grpc-definitions.git
```

2. Change directory to the project folder

```bash
$ cd agrivanna-grpc-definitions
```

3. Install protobuff compiler

```bash
$ brew install protobuf
```

4. Install the Go dependencies

```bash
$ make install
```

5. Generate the gRPC code

```bash
$ make proto
```
