# Agrivanna gRPC Definitions

[![Production Release](https://github.com/Vortex-Livestock/agrivanna-grpc-definitions/actions/workflows/production_release.yaml/badge.svg)](https://github.com/Vortex-Livestock/agrivanna-grpc-definitions/actions/workflows/production_release.yaml)
[![Continuous Integration](https://github.com/Vortex-Livestock/agrivanna-grpc-definitions/actions/workflows/continuous_integration.yaml/badge.svg)](https://github.com/Vortex-Livestock/agrivanna-grpc-definitions/actions/workflows/continuous_integration.yaml)

Agrivanna gRPC Definitions is a repository that contains the gRPC definitions for the Agrivanna project.

## Documentation

- [Setup](#setup)
- [Quick Start](#quick-start)
- [Other Commands](#other-commands)
- [Release](#release)
- [Contributors](#contributors)
- [License](#license)

## Setup

The following programs should be installed:

> [!IMPORTANT]
>
> If you are using Windows or Linux, you need to install the following programs using the package manager of your operating system. The following instructions use Homebrew, which is a package manager for Mac OS.

- [Homebrew](https://brew.sh/)
- [Go](https://golang.org/doc/install)
- [Protocol Buffer Compiler](https://grpc.io/docs/protoc-installation/)
- [Go Plugins](https://grpc.io/docs/languages/go/quickstart/)

## Quick Start

1. Clone the repository

```bash
$ git clone git@github.com:Vortex-Livestock/agrivanna-grpc-definitions.git
```

2. Change directory to the project folder

```bash
$ cd agrivanna-grpc-definitions
```

3. Install the Go dependencies

```bash
$ make install
```

## Other Commands

- Generate the gRPC code

> [!TIP]
>
> Only run this command if you have made changes to the `.proto` files.

```bash
$ make proto
```

- To clean the generated code

```bash
$ make clean
```

## Release

- [v0.0.2](https://github.com/Vortex-Livestock/agrivanna-grpc-definitions/releases/tag/v0.0.2)

## Contributors

- [Axel Sanchez](https://github.com/Axeloooo)

## License

[MIT](https://opensource.org/licenses/MIT)
