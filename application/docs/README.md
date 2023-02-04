# directory structure of application

[日本語](README-jp.md)

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
