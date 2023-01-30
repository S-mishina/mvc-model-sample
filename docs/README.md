[日本語](README-jp.md)
# introduction

This repository is a Sample WEB Application created in the MVC model. This program was created using the WEB Application Framework Gin of the Go language.

## What is MVC?
![mvc model](image/mvc-model.png)

## Directory Structures

### Overall directory organization

```terminal:Overall directory structure
├── application/
├── infra/
```

`application/`: The code related to Web Application is stored.

`infra/`: Codes related to infra are stored.

### directory structure of application

```terminal:Directory Structures
├── application/
│   ├── config/
│   ├── controller/
│   ├── docs/
│   ├── migration/
│   ├── model/
│   │   └── orm/
│   │       └── ent/
│   ├── tools/
│   └── view/
```

This section describes the directories under application.

`config/`: It contains connection information and connection-related code for the BACKEND system (DB, etc.).

`docs/`: Contains documentation data about the application, ex) ER diagrams, Go doc, etc.

`migration/`: The migration file is stored.

`tools/`: It contains tools related to the application.

`model/`:The code related to M (Model) in the MVC model is stored. ORM code is also included in the model.

`view/`: The code related to V (View) in the MVC model is stored.

`controller/`: The code (router, controller) related to (Controller) in the MVC model is stored.

The files located directly under the application directory are basically the files needed to configure a Go application.

ex) go.mod, go.sum, main.go, Dockerfile, docker-compose.yaml, etc.

### infra directory structure

```terminal:infra directory structure
├── infra/
    ├── database/
```

`database/`: The files necessary to start the database (mysql) container are located here. ex) mysql.conf, Dockerfile

## How to execute this program

### When launching an application using docker

```terminal:terminal
> docker-compose build
> docker-compose up -d
```

### To launch an application without docker

<br>
(Remarks.)

When running the application locally, you will need to load the .env file.

Create a `.env` file under `/application`.

```:.env
MYSQL_ROOT_ID=root
MYSQL_ROOT_PASSWORD=pass
MYSQL_HOST=localhost
MYSQL_PORT=3306
MYSQL_DATABASE=test
```

### Activation Method

1. DB is started by docker container.

```terminal:terminal
> docker-compose -f infra/mysql.yaml build
> docker-compose -f infra/mysql.yaml up -d
```

2. Launch the application.

```terminal:Launch the application
> go run main.go
```

### DB Migration

When the application is run for the first time, it is necessary to perform migration to the DB.

```terminal:
# migrate to the application directory.
> application/tools/run_migration.sh -v <migration version> -e root:pass@localhost:3306/test -f migration/
```

## References

to be filled in later
