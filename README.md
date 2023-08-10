# clean-architecture-go
Explore ports and adapters architecture in go

## Folder structure

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

```bash
docker-compose build
```

```bash
docker-compose up -d
```

```bash
docker exec -it cqrs-java-mysql-container-1 bash
```

# 📕Recommended libraries

- [Envconfig](https://github.com/kelseyhightower/envconfig) for environment variables
