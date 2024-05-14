# netsly-golang

_Work in progress_

## Description

Netsly is a social network based on the idea of sharing and discussing jokes. It can be used through multiple entrypoints, such as a traditional REST/WebSocket API, SSR (server side rendering) web application, and as a gRPC API.

## Architecture

The project is built based on the Clean Architecture with the use of Domain Driven Design patterns.

Layers of the application can be shown as follows:

![architecture diagram](images/architecture_circle_diagram.png)

All the code is located in the `internal` directory. The `cmd` directory contains the entry point of the application.

## Key technologies

The following list describes technologies used in the project:

- All the main code is written in Golang.
- For configuration `.env` files are used.
- The project is containerized and can be run with Docker.
- The project is not using Makefiles, but instead, it uses [Task](https://taskfile.dev/) as a task runner.

- Domain layer:

  - [Mockery](https://github.com/vektra/mockery) is used for mock generation.
  - [Playground-Validator](https://github.com/go-playground/validator) is used for validation.

- Persistence layer:

  - [Goose](https://github.com/pressly/goose) is used as a database migration tool.
  - [SQLBoiler](https://github.com/volatiletech/sqlboiler) is used as an ORM.
  - PostgreSQL is used as a database.
  - Redis is used as a cache.

- Service layer:

  - [ElasticMail](https://elasticemail.com/) is used as an email api.
  - [Asynq](https://github.com/hibiken/asynq) library with Redis database is used as a task queue.
  - JWT tokens are used for authentication.

- Presentation layer:

  - For the REST API, the [OpenAPI](https://learn.openapis.org/) and [AsycAPI](https://www.asyncapi.com/en) specifications are use for documentation. [Ogen](https://github.com/ogen-go/ogen) library is user as the code generator for the REST API.
  - For the SSR web application, golang html/template library is used with the use of [HTMX](https://htmx.org/) and [Tailwind CSS](https://tailwindcss.com/).
  - For the gRPC API, the [gRPC framework](https://grpc.io/) is used with additional tools by [Buf](https://buf.build/).

## Development

The project is supposed to be developed inside a devcontainer, but it can be run locally as well.

### Dev Container setup

#### Requirements

Ensure you have the following installed:

- `docker`
- VSCode with the Dev Containers extension installed

#### Initialise

Firstly you need to clone the repo. After that, open the project in VSCode and click on the "Reopen in Container" button in the Remote Explorer menu.

The Dev Container will be built and you will be able to start developing.

Also, you can provide your own dotfiles by using [this](https://code.visualstudio.com/docs/devcontainers/containers#_personalizing-with-dotfile-repositories) VSCode feature.

### Local setup

#### Requirements

Ensure you have the following installed:

- `go`
- `docker`
- `task`
- `npm`

#### Initialise

Firstly you need to clone the repo. After that, you need to run the following script:

```
bash .devcontainer/post-create.sh
```

It will prepare the environment for you.
