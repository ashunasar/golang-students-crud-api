# 🛠️ Students API — Development Guide (Step-by-Step)

This guide walks you through the process of building a Students CRUD API in Go using SQLite. You can follow it without needing to refer to the original code.

---

## 1. Project Setup

- Initialize the project folder `students-api`
- Create the main entry file: `cmd/students-api/main.go`
- Initialize Git and add `.gitignore`

```bash
git init
touch .gitignore
```

## 2. Configuration

- Create a config file: `config/local.yaml`

```yaml
env: 'dev'
storage_path: 'storage/storage.db'
http_server:
  address: 'localhost:8082'
```

- Create a folder `internal/config/` and define a struct matching the config fields.
- Use the [cleanenv](https://github.com/ilyakaznacheev/cleanenv) package to load config into the struct.
- In `main.go`, call the `loadConfig()` function.

## 3. Basic Server Setup

- Use `http.NewServeMux()` to register a basic route:

```go
GET / → returns "Hello world from Go"
```

- Start the HTTP server in `main.go`
- Run the app with:

```bash
go run cmd/students-api/main.go -config config/local.yaml
```

- Use `slog` for logging
- Test basic route in Postman

## 4. Build HTTP Handlers

- Create the folder: `internal/http/handlers/student/`
- In `student.go`, define a handler for `POST /api/student` to create a student.

## 5. Define Student Model

- Create `internal/models/models.go`
- Define a struct:

```go
type Student struct {
    Id    int    `json:"id"`
    Name  string `json:"name" validate:"required"`
    Email string `json:"email" validate:"required,email"`
    Age   int    `json:"age" validate:"required"`
}
```

## 6. Response Utility

- Create `internal/utils/response/response.go`
- Add a `WriteJSON()` function to format all responses
- Format error responses as:

```json
{
  "status": "Error",
  "error": "[Error Message]"
}
```

## 7. Validation

- Install validator:

```bash
go get github.com/go-playground/validator/v10
```

- Add a `ValidationError()` helper in `response.go` to format field errors

## 8. Storage Abstraction

- Create `internal/storage/storage.go`
- Define a `Storage` interface with method `CreateStudent(Student) (int64, error)`

## 9. SQLite Implementation

- Create folder: `internal/storage/sqlite/`
- In `sqlite.go`, define a struct `Sqlite` with a field `db *sql.DB`
- Use the SQLite3 driver:

```go
import _ "github.com/mattn/go-sqlite3"
```

- Create a `New()` function to open the database using the path from config

## 10. Plug Storage into Main

- In `main.go`, create the SQLite instance and pass it to handlers

## 11. Implement Handlers

- Use the `Sqlite` instance in `student.go`
- Implement:

  - `CreateStudent`
  - `GetStudentById`
  - `GetAllStudents`
  - `UpdateStudent`
  - `DeleteStudent`

## ✅ Done!

You now have a fully working Students CRUD API using Go + SQLite.

---

You can now test all routes using Postman or any REST client.

Happy Coding 🚀
