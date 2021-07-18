## Introduction

Welcome to my project

## Development

---

### Preresquites

- Go >=1.13

- Docker

- Dbmate

- Makefile

### Get the project

```shell
$ git clone https://github.com/nqtinh/go-gin-project.git
```

---

### Install dependencies

- With go :

```shell
$ make init
```

- With dbmate ( Homebrew ) :

```shell
$ brew install dbmate
```

---

### Environment setup

- Create a `config/app.yaml` file with content:

  ```yaml
    dsn: "postgres://postgres:postgres@localhost:5434/api?sslmode=disable"
    redis_client_host: "localhost"
    redis_client_port: 6377
  ```

---

### How to start

#### Main database setup

- Run ```docker-compose.yaml``` to prepare PostgreSQL and Redis.

  ```shell
  $ make setup-db
  ```


- Run PostgreSQL migration:

  ```shell
  $ make migration-db
  ```

---

### Develop

- Run server main.

  ```shell
  $ make dev
  ```

- Run generate mock test.

  ```shell
  $ make generate-test
  ```

- Run test.

  ```shell
  $ make test
  ```

- Generate swagger.

  ```shell
  $ make serve-swagger
  ```

---

### Requirements:

- 1. Create a new user
- 2. Login user
- 3. Create a new account