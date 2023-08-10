<h1 align="center">
  🚀 🐹 Clean Architecture in Go 🐹 🚀
</h1>

<p align="center">
    <a href="#"><img src="https://img.shields.io/badge/technology-go-blue.svg" alt="Go"/></a>
</p>
<p align="center">
  This repository explore ports and adapters architecture in Go.
</p>

## 🧲 Environment Setup

### 🛠️ Needed tools

1. Go 1.20 (recommended)
2. Docker and docker compose are recommended but not required

### 🏃🏻 Application execution

1. Make sure to download all Needed tools
2. Clone the repository
```bash
git clone https://github.com/dasalgadoc/clean-architecture-go.git
```
3. Build up go project
```bash
go get .
```

4. Run the API

Build the project
```bash
docker-compose build
```

Run the project
```bash
docker-compose up -d
```

An util command to see a docker
```bash
docker exec -it <NAME> bash
```

## 🧳 Project explanation

A clean architecture is an architectural pattern that separates the elements of a software application into different layers, so that it is easier to manage and develop the application. The main idea is that the application should be independent of the database, the UI, the framework, etc. This allows you to test the application without any of these elements. The clean architecture is based on the SOLID principles.

In Go, the clean architecture is based on the following elements:

### Folder structure

```
.
├── cmd
│   └── api
│       ├── bootstrap
│       │   └── bootstrap.go
│       └── main.go
├── internal
│   └── platform
│       ├── bus
│       │   └── <bus_implementation>
│       │       └── <bus_implementation>.go
│       ├── server
│       │   ├── hanlder
│       │   │   ├── <resource_1>
│       │   │   │   └── <resource_1>.go
│       │   │   └── <resource_2>
│       │   │       └── <resource_2>.go
│       │   ├── middleware
│       │   │   └── <concrete_middleware>
│       │   │       └── <concrete_middleware>.go
│       │   └── server.go
│       └── storage
│           └── <storage_implementation>
│               └── <storage_implementation>.go
├── kit
.
```

|           Root           |   Folder   | Description                                        |
|:------------------------:|:----------:|----------------------------------------------------|
|            .             |    cmd     | Contains all entry points to the application       |
|           cmd            |    api     | Define a concrete API entry point                  |
|         cmd/api          | bootstrap  | Groups functionally to setup application           |
|            .             |  internal  | Private application and library code.              |
|            .             |    kit     | All elements `shared` in the project               |
|         internal         |  platform  | Infrastructure (I/O) code                          |
|    internal/platform     |    bus     | Implementations of buses (commands, query, events) |
|    internal/platform     |   server   | Server code                                        |
| internal/platform/server |  handler   | Bunch of each server entry point                   |
| internal/platform/server | middleware | Middleware functions                               |
|    internal/platform     |  storage   | Storage code, means all persistence tool           |


# 📕Recommended libraries

- [Envconfig](https://github.com/kelseyhightower/envconfig) for environment variables
- [Wire](https://github.com/google/wire) for dependency injection by Google
- [fx](https://github.com/uber-go/fx) for dependency injection by Uber
