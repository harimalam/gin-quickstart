<p align="center">
  <img src="https://raw.githubusercontent.com/gin-gonic/logo/master/color.png" alt="Gin Logo" width="200"/>
</p>

<h1 align="center">🚀 Gin Quickstart</h1>

<p align="center">
  <strong>A production-ready RESTful API boilerplate built with Go, Gin Framework, PostgreSQL, and JWT Authentication</strong>
</p>

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.25-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go Version"/>
  <img src="https://img.shields.io/badge/Gin-1.11-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Gin Version"/>
  <img src="https://img.shields.io/badge/PostgreSQL-Latest-336791?style=for-the-badge&logo=postgresql&logoColor=white" alt="PostgreSQL"/>
  <img src="https://img.shields.io/badge/Docker-Ready-2496ED?style=for-the-badge&logo=docker&logoColor=white" alt="Docker"/>
  <img src="https://img.shields.io/badge/JWT-Auth-000000?style=for-the-badge&logo=jsonwebtokens&logoColor=white" alt="JWT"/>
</p>

---

## 📋 Table of Contents

- [Features](#-features)
- [Architecture](#-architecture)
- [Project Structure](#-project-structure)
- [Prerequisites](#-prerequisites)
- [Quick Start](#-quick-start)
- [Configuration](#-configuration)
- [API Endpoints](#-api-endpoints)
- [Authentication](#-authentication)
- [API Usage Examples](#-api-usage-examples)
- [Docker Deployment](#-docker-deployment)
- [Development](#-development)
- [Tech Stack](#-tech-stack)
- [License](#-license)

---

## ✨ Features

| Feature                         | Description                                                      |
| ------------------------------- | ---------------------------------------------------------------- |
| 🏗️ **Clean Architecture**       | Repository-Service-Handler pattern for maintainable code         |
| 🔐 **JWT Authentication**       | Secure token-based authentication with role-based access control |
| 👥 **Role-Based Authorization** | Admin and User roles with different permissions                  |
| 🗄️ **PostgreSQL + GORM**        | Robust database with auto-migrations                             |
| 🐳 **Docker Ready**             | Multi-stage Dockerfile for optimized production builds           |
| ⚙️ **Environment Config**       | Flexible configuration via `.env` or environment variables       |
| 🔒 **Password Hashing**         | Secure bcrypt password hashing                                   |
| 📝 **CRUD Operations**          | Complete Create, Read, Update, Delete functionality              |

---

## 🏛️ Architecture

This project follows a **Clean Architecture** pattern with clear separation of concerns:

```
┌─────────────────────────────────────────────────────────────┐
│                        HTTP Layer                           │
│                    (Gin Router + Handlers)                  │
├─────────────────────────────────────────────────────────────┤
│                      Middleware Layer                       │
│              (Auth Middleware + Authorization)              │
├─────────────────────────────────────────────────────────────┤
│                       Service Layer                         │
│                    (Business Logic)                         │
├─────────────────────────────────────────────────────────────┤
│                      Repository Layer                       │
│                     (Data Access)                           │
├─────────────────────────────────────────────────────────────┤
│                       Database Layer                        │
│                  (PostgreSQL + GORM)                        │
└─────────────────────────────────────────────────────────────┘
```

---

## 📁 Project Structure

```
gin-quickstart/
├── cmd/
│   └── main.go                 # Application entry point
├── internal/
│   ├── albums/                 # Albums feature module
│   │   ├── handler.go          # HTTP handlers (controllers)
│   │   ├── model.go            # Data models & DTOs
│   │   ├── repository.go       # Database operations
│   │   └── service.go          # Business logic
│   ├── auth/                   # Authentication module
│   │   ├── handler.go          # Auth HTTP handlers
│   │   ├── model.go            # User model & request DTOs
│   │   ├── repository.go       # User data operations
│   │   ├── service.go          # Auth business logic
│   │   └── token.go            # JWT token utilities
│   ├── config/
│   │   └── config.go           # Configuration management
│   ├── db/
│   │   └── db.go               # Database initialization
│   └── middleware/
│       └── auth_middleware.go  # JWT & authorization middleware
├── pkg/                        # Shared utilities (if any)
├── .env                        # Environment variables (create this)
├── docker-compose.yaml         # Docker Compose configuration
├── Dockerfile                  # Multi-stage Docker build
├── go.mod                      # Go module dependencies
├── go.sum                      # Dependency checksums
└── readme.md                   # Project documentation
```

---

## 📋 Prerequisites

- **Go** 1.21 or higher
- **PostgreSQL** 13 or higher
- **Docker & Docker Compose** (optional, for containerized deployment)

---

## 🚀 Quick Start

### Option 1: Run Locally

**1. Clone the repository:**

```bash
git clone https://github.com/yourusername/gin-quickstart.git
cd gin-quickstart
```

**2. Create environment file:**

```bash
cat > .env << EOF
APP_PORT=8080
JWT_SECRET=your_32_byte_secret_key_here_abcd

DB_HOST=localhost
DB_PORT=5432
DB_USER=root
DB_PASSWORD=password
DB_NAME=gin_db
SSL_MODE=disable
EOF
```

**3. Install dependencies:**

```bash
go mod download
```

**4. Run the application:**

```bash
go run cmd/main.go
```

### Option 2: Run with Docker

```bash
# Start the application
docker compose up -d

# View logs
docker compose logs -f

# Stop the application
docker compose down
```

---

## ⚙️ Configuration

The application uses environment variables for configuration. Create a `.env` file in the root directory:

| Variable      | Description                               | Default     |
| ------------- | ----------------------------------------- | ----------- |
| `APP_PORT`    | Server port                               | `8080`      |
| `JWT_SECRET`  | Secret key for JWT signing (min 32 chars) | Required    |
| `DB_HOST`     | PostgreSQL host                           | `localhost` |
| `DB_PORT`     | PostgreSQL port                           | `5432`      |
| `DB_USER`     | Database username                         | Required    |
| `DB_PASSWORD` | Database password                         | Required    |
| `DB_NAME`     | Database name                             | Required    |
| `SSL_MODE`    | PostgreSQL SSL mode                       | `disable`   |

---

## 📡 API Endpoints

### Authentication Routes (Public)

| Method | Endpoint              | Description             |
| ------ | --------------------- | ----------------------- |
| `POST` | `/api/v1/auth/signup` | Register a new user     |
| `POST` | `/api/v1/auth/login`  | Login and get JWT token |

### Album Routes (Protected)

| Method   | Endpoint             | Description      | Role Required   |
| -------- | -------------------- | ---------------- | --------------- |
| `GET`    | `/api/v1/albums/`    | Get all albums   | `user`, `admin` |
| `GET`    | `/api/v1/albums/:id` | Get album by ID  | `user`, `admin` |
| `POST`   | `/api/v1/albums/`    | Create new album | `admin` only    |
| `PUT`    | `/api/v1/albums/:id` | Update album     | `admin` only    |
| `DELETE` | `/api/v1/albums/:id` | Delete album     | `admin` only    |

---

## 🔐 Authentication

This API uses **JWT (JSON Web Tokens)** for authentication with **role-based access control**.

### Roles

| Role    | Permissions                |
| ------- | -------------------------- |
| `user`  | Read-only access to albums |
| `admin` | Full CRUD access to albums |

### Token Structure

```json
{
  "id": 1,
  "role": "admin",
  "exp": 1764330628
}
```

> ⏰ Tokens expire after **24 hours**

---

## 📖 API Usage Examples

### 1. Register a New User

```bash
# Register as admin
curl -X POST http://localhost:8080/api/v1/auth/signup \
  -H "Content-Type: application/json" \
  -d '{
    "username": "admin",
    "password": "admin123",
    "role": "admin"
  }'

# Register as regular user
curl -X POST http://localhost:8080/api/v1/auth/signup \
  -H "Content-Type: application/json" \
  -d '{
    "username": "john",
    "password": "john123",
    "role": "user"
  }'
```

**Response:**

```json
{
  "data": {
    "user": {
      "id": 1,
      "username": "admin",
      "role": "admin"
    }
  },
  "message": "User registered successfully"
}
```

### 2. Login

```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "admin",
    "password": "admin123"
  }'
```

**Response:**

```json
{
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  },
  "message": "Login successful"
}
```

### 3. Get All Albums (Protected)

```bash
curl http://localhost:8080/api/v1/albums/ \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

**Response:**

```json
{
  "data": {
    "albums": [
      {
        "ID": 1,
        "title": "Blue Train",
        "artist": "John Coltrane",
        "CreatedAt": "2024-01-01T00:00:00Z",
        "UpdatedAt": "2024-01-01T00:00:00Z"
      }
    ]
  },
  "message": "Albums retrieved successfully"
}
```

### 4. Create Album (Admin Only)

```bash
curl -X POST http://localhost:8080/api/v1/albums/ \
  -H "Authorization: Bearer YOUR_ADMIN_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Kind of Blue",
    "artist": "Miles Davis"
  }'
```

**Response:**

```json
{
  "data": {
    "album": {
      "ID": 2,
      "title": "Kind of Blue",
      "artist": "Miles Davis",
      "CreatedAt": "2024-01-01T00:00:00Z",
      "UpdatedAt": "2024-01-01T00:00:00Z"
    }
  },
  "message": "Album created successfully"
}
```

### 5. Get Album by ID

```bash
curl http://localhost:8080/api/v1/albums/1 \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### 6. Update Album (Admin Only)

```bash
curl -X PUT http://localhost:8080/api/v1/albums/1 \
  -H "Authorization: Bearer YOUR_ADMIN_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Updated Title",
    "artist": "Updated Artist"
  }'
```

### 7. Delete Album (Admin Only)

```bash
curl -X DELETE http://localhost:8080/api/v1/albums/1 \
  -H "Authorization: Bearer YOUR_ADMIN_TOKEN"
```

---

## 🐳 Docker Deployment

### Multi-Stage Dockerfile

The project includes an optimized multi-stage Dockerfile:

- **Stage 1 (Builder)**: Compiles the Go binary with optimizations
- **Stage 2 (Production)**: Minimal Alpine image (~15MB)

### Build & Run

```bash
# Build the image
docker build -t gin-quickstart .

# Run the container
docker run -p 8080:8080 \
  -e DB_HOST=your-db-host \
  -e DB_PORT=5432 \
  -e DB_USER=root \
  -e DB_PASSWORD=password \
  -e DB_NAME=gin_db \
  -e SSL_MODE=disable \
  -e JWT_SECRET=your_32_byte_secret_key_here_abcd \
  -e APP_PORT=8080 \
  gin-quickstart
```

### Docker Compose

```bash
# Start all services
docker compose up -d

# View logs
docker compose logs -f app

# Stop all services
docker compose down

# Rebuild and restart
docker compose up -d --build
```

---

## 🛠️ Development

### Running Tests

```bash
go test ./...
```

### Building for Production

```bash
# Build optimized binary
CGO_ENABLED=0 go build -ldflags='-w -s' -o goapp ./cmd/main.go
```

### Code Structure Conventions

- **Handlers**: HTTP request/response handling
- **Services**: Business logic implementation
- **Repositories**: Database operations abstraction
- **Models**: Data structures and DTOs
- **Middleware**: Request processing (auth, logging, etc.)

---

## 🛠️ Tech Stack

| Technology                                              | Purpose                  |
| ------------------------------------------------------- | ------------------------ |
| [Go](https://golang.org/)                               | Programming Language     |
| [Gin](https://gin-gonic.com/)                           | Web Framework            |
| [GORM](https://gorm.io/)                                | ORM Library              |
| [PostgreSQL](https://www.postgresql.org/)               | Database                 |
| [JWT](https://jwt.io/)                                  | Authentication           |
| [Viper](https://github.com/spf13/viper)                 | Configuration Management |
| [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) | Password Hashing         |
| [Docker](https://www.docker.com/)                       | Containerization         |

---

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

<p align="center">
  Made with ❤️ and Go
</p>

<p align="center">
  <a href="#-gin-quickstart">⬆️ Back to Top</a>
</p>
