# PROJECT-STARTER

Brief project description.

## Table of Contents

- [Folder Structure](#folder-structure)
- [Install Dependencies](#install-dependencies)
- [How to Run](#how-to-run)
- [How to Build](#how-to-build)

## Folder Structure

- **api**: Contains API-related logic, including request and response handling.
- **config**: Configuration-related files.
- **constants**: Constants used throughout the project.
- **controller**: Controllers that handle business logic.
- **dal**: Data Access Layer for interacting with the database.
  - `configure.go`, `db.go`: DAL files for different entities.
- **docs**: Documentation files.
  - `docs.go`: (Description of this file)
  - `swagger.json` and `swagger.yaml`: Swagger documentation files.
- **main.go**: Entry point of the application.
- **model**: Data models.
- **scripts**: Shell scripts.
- **server**: Server-related logic.
  - `server.go`: Logic related to the server.
- **utils**: Utility functions.

### Install Dependencies

```bash
➜ go mod download
```

### Install Sqlite3

```bash
➜ sudo apt install sqlite3
➜ sqlite3 --version
Output
3.31.1 2020-01-27 19:55:54 3bfa9cc97da10598521b342961df8f5f68c7388fa117345eeb516eaa837balt1
```

### How to Run

```bash
➜ go run main.go
```

### How to Build

```bash
➜ go build
```

Step 2: Run Golang Service

```bash
➜ ./project-starter
```