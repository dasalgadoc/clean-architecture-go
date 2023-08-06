# clean-architecture-go
Explore ports and adapters architecture in go

## Folder structure

```
.
├── cmd
│   ├── api
│   │   ├── bootstrap
│   │   │   └── bootstrap.go
│   │   └── main.go
│   ├── internal
│   │   ├── platform
│   │   │   ├── server
│   │   │   │   ├── hanlder
│   │   │   │   │   ├── <resource_1>
│   │   │   │   │   │   └── <resource_1>.go
│   │   │   │   │   ├── <resource_2>
│   │   │   │   │      └── <resource_2>.go
│   │   │   │   └── server.go
.
```

|                   Folder                    | Description                                   |
|:-------------------------------------------:|-----------------------------------------------|
|                     cmd                     | Contains all entry points to the application  |
|                   cmd/api                   | Define a concrete API entry point             |
|              cmd/api/bootstrap              | Groups functionally to setup application      |
|                  internal                   | Private application and library code.         |
|              internal/platform              | Infrastructure (I/O) code                     |
|          internal/platform/server           | Server code                                   |
|      internal/platform/server/handler       | Bunch of each server entry point              |
| internal/platform/server/handler/<resource> | Entry point code for a functionally `reource` |
