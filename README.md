# sandbox
砂場

## ディレクトリ構成

```shell
.
├── Makefile
├── aqua.yaml
├── cmd
│   └── server
│       ├── main.go
│       ├── wire.go
│       └── wire_gen.go
├── go.mod
├── go.sum
├── internal
│   ├── config
│   │   ├── env
│   │   │   ├── env.go
│   │   │   └── secret.go
│   │   ├── errs
│   │   │   └── errs.go
│   │   ├── logger
│   │   │   └── logger.go
│   │   └── tracer
│   │       └── tracer.go
│   ├── domain
│   │   ├── entity
│   │   │   ├── todos
│   │   │   └── valueobject
│   │   ├── repository
│   │   │   └── todos.go
│   │   └── service
│   │       └── todos.go
│   ├── infrastructure
│   │   ├── db
│   │   │   ├── client.go
│   │   │   ├── migrations
│   │   │   ├── models
│   │   │   ├── sql
│   │   │   └── sqlboiler.toml
│   ├── mockery.go
│   ├── presentation
│   │   ├── handler
│   │   │   └── handler.go
│   │   └── http
│   │       ├── todos
│   │       ├── errs
│   │       ├── helper
│   │       └── router.go
│   └── usecase
│       └── todos
│           ├── create_todo.go
│           ├── get_todo.go
│           ├── update_todo.go
│           └── delete_todo.go
└── lefthook.yml
```
