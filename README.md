# 📘 Students API — CRUD App in Go + SQLite

A simple and clean **RESTful API** built with **Go** and **SQLite**, designed to manage students using full CRUD operations. Great for beginners learning Go's modular project layout, routing, SQLite, and input validation.

GitHub Repo 👉 [github.com/ashunasar/golang-students-crud-api](https://github.com/ashunasar/golang-students-crud-api)

---

## 📁 Project Structure

```
students-api/
├── cmd/
│   └── students-api/          # Main entry point (main.go)
├── config/
│   └── local.yaml             # App config
├── internal/
│   ├── config/                # Loads config (cleanenv)
│   ├── http/handlers/student/ # All student API handlers
│   ├── models/                # Student model definitions
│   ├── storage/               # Storage interface + SQLite impl
│   └── utils/response/        # Response formatting & validation errors
├── storage/
│   └── storage.db             # SQLite DB file
├── go.mod
├── go.sum
├── README.md
└── students-api.postman_collection.json  # Postman collection for testing
```

---

## 🔧 Configuration (`config/local.yaml`)

```yaml
env: 'dev'
storage_path: 'storage/storage.db'
http_server:
  address: 'localhost:8082'
```

---

## 🚀 Getting Started

### 1. Clone and Initialize

```bash
git clone https://github.com/ashunasar/golang-students-crud-api.git
cd golang-students-crud-api
go mod tidy
```

### 2. Install Required SQLite Driver

The project uses the Go SQLite3 driver:

```bash
go get github.com/mattn/go-sqlite3
```

### 3. Run the Application

```bash
go run cmd/students-api/main.go -config config/local.yaml
```

This will:

- Load config
- Open SQLite DB (`storage.db`)
- Register HTTP routes
- Start server at `localhost:8082`

---

## 🤖 API Endpoints

| Method | Endpoint             | Description       |
| ------ | -------------------- | ----------------- |
| POST   | `/api/student`       | Create a student  |
| GET    | `/api/students`      | Get all students  |
| GET    | `/api/students/{id}` | Get student by ID |
| PUT    | `/api/students`      | Update student    |
| DELETE | `/api/students/{id}` | Delete student    |

---

## 📢 Sample Requests (Postman Examples)

### ✅ Create Student

**POST** `/api/student`

```json
{
  "name": "hello",
  "email": "example@xample.com",
  "Age": 18
}
```

---

### ✅ Get All Students

**GET** `/api/students`

---

### ✅ Get Student by ID

**GET** `/api/students/1`

---

### ✅ Update Student

**PUT** `/api/students`

```json
{
  "id": 1,
  "name": "updated name",
  "email": "hello2@xample.com",
  "Age": 19
}
```

---

### ✅ Delete Student

**DELETE** `/api/students/1`

---

## ⚙️ Tech Used

- 🐹 Go
- 🗓 SQLite
- 📌 [`github.com/mattn/go-sqlite3`](https://github.com/mattn/go-sqlite3) for SQLite driver
- 📦 [go-playground/validator](https://github.com/go-playground/validator/v10) for validation
- 📘 [cleanenv](https://github.com/ilyakaznacheev/cleanenv) for config
- 📡 Native `net/http`
- 📄 JSON REST API

---

## ✅ JSON Response Format

All responses follow this standard:

### On Success:

```json
{
  "status": "OK",
  "data": {
    "id": 1,
    "name": "John"
  }
}
```

### On Error:

```json
{
  "status": "Error",
  "error": "student not found"
}
```

---

## 🔍 Notes for Developers

- Logging is handled using Go’s `slog`
- Routes are handled manually with `http.NewServeMux`
- SQLite storage is modular and implements an interface
- Validations use tags like `validate:"required,email"`

---

## 📅 Postman Collection

Use the provided file:

```
students-api.postman_collection.json
```

You can import this into Postman to test all API routes quickly.

---

## ✅ TODOs for Future Enhancements

- Add pagination
- Add search/filter by name or email
- Add Swagger or Redoc API docs
- Dockerize the application
- Add unit tests

---

## 👏 Credits

Created by [ashunasar](https://github.com/ashunasar)
Feel free to ⭐️ star the repo, fork it, and contribute!

---
