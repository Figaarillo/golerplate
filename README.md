[English](./README.md) | [Spanish](./README.es.md)

# GOlerplate

GOlerplate is a template for web applications in Golang. This means it can be a starting point for web applications in Golang, designed to facilitate development with a clean and modular architecture.

## Technologies

- **Language**: Go
- **Frameworks and Libraries**: Gorilla Mux, GORM
- **Containers**: Docker, Docker Compose
- **Documentation**: Swagger

## Getting Started

### Prerequisites

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/)
- [Go](https://golang.org/doc/install) (optional if you want to use CLI)

### Clone the repository

1. Clone the repository:

```sh
git clone https://github.com/Figaarillo/golerplate.git
cd golerplate
```

2. Copy the `.env.example` file to `.env`:

```sh
cp .env.example .env
```

### Usage

- Run the server and database through Docker

```sh
make docker.run
```

- Run the server locally and the database with Docker

```sh
make run
```

### How to run tests?

#### Unit Tests

To run all unit tests:

```sh
make test.unit
```

To run a single unit test:

```sh
make test.unit.[entity_name]
```

For example, to run the unit test for category:

```sh
make test.unit.category
```

#### Integration Tests

To run all integration tests:

```sh
make test.e2e
```

To run a single integration test:

```sh
make test.e2e.[entity_name]
```

For example, to run the integration test for category:

```sh
make test.e2e.category
```

## Project Structure

```sh
.
├── cmd
│   └── api                # Main entry point of the application
├── docs                   # Swagger generated documentation
├── internal
│   ├── application
│   │   └── usecase        # Application use cases
│   ├── domain
│   │   ├── entity         # Domain entity definitions
│   │   ├── exception      # Domain exception handling
│   │   └── repository     # Repository interfaces
│   ├── infrastructure
│   │   ├── handler        # HTTP handlers
│   │   ├── middleware     # HTTP middlewares
│   │   ├── repository     # Repository implementations
│   │   └── router         # Route definitions
│   ├── setup              # Initial setup
│   ├── share
│   │   ├── config         # Shared configuration
│   │   ├── exception      # Shared exception handling
│   │   ├── utils          # Shared utilities
│   │   └── validation     # Shared validations
│   └── test               # Unit and integration tests
└── scripts                # Automation scripts
```

## Features

- REST API: CRUD implementation for entities.
- Tests: Unit and integration tests.
- Documentation: API documentation with Swagger.
- Architecture: Implements clean and hexagonal architecture principles.

## To-do

- [ ] Add authentication and authorization support using OAuth2.
- [ ] Improve performance when compiling the application with Docker.
- [ ] Improve the documentation with Swagger.
- [ ] Improve the unit and integration tests adding more test cases.
- [ ] Improve error handling, adding more exceptions, error messages, etc.
- [ ] Add goroutines to improve application performance.
- [ ] Add handling of logs.
- [ ] Integrate other server and database frameworks

<!--## Contributing-->
<!---->
<!--If you want to contribute to the project, please follow these steps:-->
<!---->
<!--1. Fork the repository.-->
<!--2. Create a branch (`git checkout -b feature/new-feature`).-->
<!--3. Make your changes (`git commit -am 'Add new feature'`).-->
<!--4. Push to the branch (`git push origin feature/new-feature`).-->
<!--5. Create a new Pull Request.-->

License

This project is licensed under the MIT License.
