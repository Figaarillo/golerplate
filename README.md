[English](./README.md) | [Spanish](./README.es.md)

# GOlerplate

GOlerplate is a template for web applications in Golang. This means it can be a starting point for web applications in Golang, designed to facilitate development with a clean and modular architecture.

## Technologies

- **[gorilla/mux](https://github.com/gorilla/mux)**: HTTP router for building Go web servers with 🦍
- **[gorm](https://gorm.io)**: ORM for Go
- **[Docker](https://www.docker.com) & [Docker Compose](https://docs.docker.com/compose/)**: For the containerization
- **[Swagger](https://swagger.io)**: API documentation

## Features

- **REST API**: CRUD implementation for entities.
- **Tests**: Unit and integration tests, to validate the correct operation of the application.
- **Documentation**: API documentation generated with Swagger.
- **Clean Architecture**: Implements clean and hexagonal architecture principles.
- **Dockerization**: Dockerization of the application and the database, to run the application in a Docker container.
- **Automation**: Automation of common tasks, such as running run the application and the database in Docker, or running unit and integration tests.


## Getting Started

### Prerequisites

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/)
- [Go](https://golang.org/doc/install) (optional if you want to use the project locally)

### Clone the repository

1. Clone the repository:

```sh
git clone https://github.com/Figaarillo/golerplate.git
cd golerplate
```

2. Set up environment variables. Copy the `.env.example` file to `.env`:

```sh
cp .env.example .env
```

3. Install dependencies:

```sh
go mod download
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
- Run the build of the server and database through Docker

```sh
make run.build
```

### How to run tests?

#### Unit Tests

- To run all unit tests:

```sh
make test.unit
```

- If you want to see the coverage of the unit tests:

```sh
make test.unit.cover
```

- To run a single unit test:

```sh
make test.unit.[entity_name]
```

- For example, to run the unit test for category:

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

### Documentation

The documentation is generated with [Swagger](https://swagger.io/docs/2.0/).

To generate the documentation, run:

```sh
make docs
```

And now you can access the documentation at *http://localhost:8080/api/docs/*

## Project Structure

For building the project, I decided to use an **architecture hexagonal** based on the principles of Clean Architecture.
In this architecture, I have three main layeres for each entity: Domain, Application, and Infrastructure 
The Domain layer contains the domain entities, domain exceptions, respository interfaces, and all logic related to the domain.  The Application layer contains the use cases of the Application, and the Infrastructure layer contains the HTTP handlers, middlewares, repositories,  routes, and all three party dependencies.

```sh
.
├── cmd
│   └── api                # Main entry point of the application
├── docs                   # Swagger generated documentation
├── internal
│   ├── application
│   │   └── usecase        # Application use cases
│   ├── bootstrap          # Bootstrapper for each entity
│   ├── domain
│   │   ├── entity         # Domain entity definitions and unit tests
│   │   ├── exception      # Domain exception handling
│   │   └── repository     # Repository interfaces
│   ├── infrastructure
│   │   ├── handler        # HTTP handlers
│   │   ├── middleware     # HTTP middlewares
│   │   ├── repository     # Repository implementations
│   │   └── router         # Route definitions
│   ├── share
│   │   ├── config         # Shared configuration
│   │   ├── exception      # Shared exception handling
│   │   ├── utils          # Shared utilities
│   │   └── validation     # Shared validations
│   └── test               # Integration tests
└── scripts                # Automation scripts
```

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

## License

This project is licensed under the MIT License.
