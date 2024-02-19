# REST API file storage

<img title="api for storage" alt="api" src="/docs/img.png">

## Methods

- signup
- signin
- verifymail
- whoami
- upload
- list
- download
- delete
- logout

## Features

- mail registration
- OTP token for authorization
- JWT session
- persistent on postgresql
- caching active user
- limit of upload size
- UI (signup, signin, list, upload, logout)
- swagger annotation
- metrics for gin and custom length of cache

## Build & Run

### Prerequisites

- go 1.22
- docker & docker-compose
- [golangci-lint](https://github.com/golangci/golangci-lint) (<i>optional</i>, used to run code checks)
- [swag](https://github.com/swaggo/swag) (<i>optional</i>, used to re-generate swagger documentation)
- [mockgen](https://github.com/golang/mock) (<i>optional</i>, used to re-generate mock)

### Installation

Change default credentials on _docker-compose.yml_, _storage-server.toml_, replacement example certs and `make run` to build & run project
