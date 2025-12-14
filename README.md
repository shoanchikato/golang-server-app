# Golang Server App

A RESTful API server built with **Go (Golang)** following **Clean Architecture principles**.
This project demonstrates a structured, scalable backend with authentication, validation, Swagger documentation, and SQLite persistence.

---

## ✨ Features

* Clean Architecture–inspired structure
* RESTful API using `net/http`
* Authentication (login, tokens)
* SQLite database
* Swagger (OpenAPI) documentation
* Request validation
* Centralized error handling
* Manual dependency injection
* MIT licensed

---

## 📁 Project Structure

```text
cmd/
  server/
    main.go          # Application entry point

internal/ (or pkg/)
  handler/           # HTTP handlers (controllers)
  repo/              # Database access layer
  stmt/              # Raw SQL statements
  validation/        # Request validation logic
  di/                # Dependency injection setup

docs/
  swagger.yaml       # OpenAPI specification

test/
  ...                # Tests

go.mod
go.sum
Makefile
```

> ⚠️ Note: If this project is an **application** (not a reusable library), it is recommended to use `internal/` instead of `pkg/`.

---

## 🚀 Getting Started

### Prerequisites

* Go **1.20+**
* SQLite installed (or bundled with Go driver)

---

### Clone the repository

```bash
git clone https://github.com/shoanchikato/golang-server-app.git
cd golang-server-app
```

---

### Install dependencies

```bash
go mod tidy
```

---

### Run the server

```bash
go run cmd/server/main.go
```

Or using Makefile (if available):

```bash
make run
```

The server will start on:

```
http://localhost:3000
```

---

## 📘 API Documentation (Swagger)

Swagger UI is available once the server is running:

```
http://localhost:3000/swagger/index.html
```

### Regenerate Swagger docs

If you modify annotations:

```bash
go install github.com/swaggo/swag/cmd/swag@latest
swag init
```

---

## 🔐 Authentication Example

### Login

**POST** `/login`

```json
{
  "username": "john_doe",
  "password": "password1"
}
```

**Response**

```json
{
  "access_token": "jwt-access-token",
  "refresh_token": "jwt-refresh-token"
}
```

---

## ⚙️ Configuration

Configuration is handled via environment variables.

Example `.env` file:

```env
PORT=3000
DB_PATH=small.db
JWT_SECRET=your-secret-key
```

> ⚠️ Do **not** commit `.env` or database files to version control.

## 🧪 Testing

Run tests with:

```bash
go test ./...
```
---

## 📄 License

This project is licensed under the **MIT License**.

---

## 👤 Author

**shoanchikato**
GitHub: [https://github.com/shoanchikato](https://github.com/shoanchikato)
