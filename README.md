# Crud-Golang-postgres

## About

Simple Golang Crud using the postgres library, without frameworks and with unit tests. This is a project with the purpose of training my skills with golang vanila and unit tests that I have studied and applied in my projects, In it you will see simple user authentication using JWT that must be sent in the Authentication Header.

## Requirements

- Golang 14.0 or higher
- Docker or postgres running on your machine
- Docker compose

## How works

- Clone repository

```markdown
git clone https://github.com/Leonardo404-code/crud-postgres-with-authentication.git
```

- Up postgres container with docker-compose

```docker
docker compose up -d
```

- Install dependencies

```go
go mod tidy
```

- Run the project

```go
go run main.go
```

Hope you enjoy and leave a star ⭐
