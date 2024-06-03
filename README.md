# GOlerplate

GOlerpalte is a boilerplate for Golang web applications. It is meant to be used as
a starting point for new projects.

## Pre-requisites

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/)
- [Go](https://golang.org/doc/install) (optional if you want to use the CLI)

## Get started

1. Clone the repository

```bash
git clone https://github.com/figarillo/gobplate.git
cd gobplate
```

2. Copy the `.env.example` to `.env`

```bash
cp .env.example .env
```

## How to use

### Run Server And Database with Docker

```bash
make docker.run
```

### Run server from local And Database with Docker

```bash
make run
```

## Structure

```bash
.
├── cmd
│   └── api
├── docs
├── internal
│   ├── application
│   │   └── usecase
│   ├── domain
│   │   ├── entity
│   │   ├── exeption
│   │   └── repository
│   ├── infrastructure
│   │   ├── handler
│   │   ├── middleware
│   │   ├── repository
│   │   └── router
│   ├── setup
│   ├── share
│   │   ├── config
│   │   ├── exeption
│   │   ├── utils
│   │   └── validation
│   └── test
└── scripts
```
