# my-gin-example

A comprehensive example of a RESTful API built with Gin framework, featuring modern Go practices, JWT authentication, CRUD operations, database integration with GORM, and more.

## Features

- 🚀 **Modern Go (1.23)** with latest dependencies
- 🔐 **JWT Authentication** with role-based access control
- 📦 **CRUD Operations** for users and products
- 🗄️ **GORM Integration** with SQLite (easily switchable to MySQL/PostgreSQL)
- ✅ **Request Validation** using Gin's binding
- 📄 **Pagination & Search** for list endpoints
- 🔒 **Rate Limiting** middleware
- 🧪 **Unit Tests** with testify
- 📝 **Structured Logging** with zerolog
- ⚙️ **Configuration Management** with Viper

## Build Status

[![CircleCI](https://circleci.com/gh/wmh/my-gin-example/tree/master.svg?style=svg)](https://circleci.com/gh/wmh/my-gin-example/tree/master)

## Prerequisites

- Go 1.23 or higher
- Make (optional)

## Installation

```bash
# Clone the repository
git clone https://github.com/wmh/my-gin-example.git
cd my-gin-example

# Install dependencies
go mod download

# Run the application
go run main.go
```

The server will start on `http://localhost:8089` by default.

## Configuration

Edit `config/app.toml` to customize:

```toml
app_port = 8089

[database]
path = "./data/app.db"

[jwt]
secret = "your-secret-key-change-in-production"
expiration_hours = 24

[logs]
disable_default_writer = false
stdout_only = true
```

## API Documentation

### Health Check

#### Check Server Status
```bash
GET /ok
```

**Response:**
```
ok
```

### Authentication (v2)

#### Register a New User
```bash
POST /v2/auth/register
Content-Type: application/json

{
  "username": "johndoe",
  "email": "john@example.com",
  "password": "securepass123",
  "full_name": "John Doe"
}
```

**Response (201):**
```json
{
  "message": "User registered successfully",
  "user": {
    "id": 1,
    "username": "johndoe",
    "email": "john@example.com",
    "full_name": "John Doe",
    "role": "user",
    "is_active": true,
    "created_at": "2025-12-03T10:00:00Z"
  }
}
```

#### Login
```bash
POST /v2/auth/login
Content-Type: application/json

{
  "username": "johndoe",
  "password": "securepass123"
}
```

**Response (200):**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": 1,
    "username": "johndoe",
    "email": "john@example.com",
    "full_name": "John Doe",
    "role": "user",
    "is_active": true,
    "created_at": "2025-12-03T10:00:00Z"
  }
}
```

### User Management (v2)

#### Get User Profile (Protected)
```bash
GET /v2/users/profile
Authorization: Bearer <token>
```

**Response (200):**
```json
{
  "id": 1,
  "username": "johndoe",
  "email": "john@example.com",
  "full_name": "John Doe",
  "role": "user",
  "is_active": true,
  "created_at": "2025-12-03T10:00:00Z"
}
```

#### Update User Profile (Protected)
```bash
PUT /v2/users/profile
Authorization: Bearer <token>
Content-Type: application/json

{
  "email": "newemail@example.com",
  "full_name": "John Updated Doe"
}
```

#### List All Users (Admin Only)
```bash
GET /v2/users
Authorization: Bearer <token>
```

### Product Management (v2)

#### List Products (Public with Pagination)
```bash
GET /v2/products?page=1&page_size=10&search=laptop&sort_by=price&order=asc
```

**Query Parameters:**
- `page` (default: 1)
- `page_size` (default: 10, max: 100)
- `search` - Search in name, description, category
- `sort_by` - Field to sort by (e.g., price, created_at)
- `order` - asc or desc (default: desc)

**Response (200):**
```json
{
  "data": [
    {
      "id": 1,
      "name": "Laptop Pro",
      "description": "High-performance laptop",
      "price": 1299.99,
      "stock": 15,
      "category": "Electronics",
      "sku": "LAP-001",
      "is_active": true,
      "created_at": "2025-12-03T10:00:00Z"
    }
  ],
  "page": 1,
  "page_size": 10,
  "total_items": 1,
  "total_pages": 1
}
```

#### Get Product by ID (Public)
```bash
GET /v2/products/1
```

#### Create Product (Protected)
```bash
POST /v2/products
Authorization: Bearer <token>
Content-Type: application/json

{
  "name": "Wireless Mouse",
  "description": "Ergonomic wireless mouse",
  "price": 29.99,
  "stock": 50,
  "category": "Accessories",
  "sku": "MOU-001"
}
```

**Response (201):**
```json
{
  "id": 2,
  "name": "Wireless Mouse",
  "description": "Ergonomic wireless mouse",
  "price": 29.99,
  "stock": 50,
  "category": "Accessories",
  "sku": "MOU-001",
  "is_active": true,
  "created_at": "2025-12-03T10:00:00Z"
}
```

#### Update Product (Protected)
```bash
PUT /v2/products/2
Authorization: Bearer <token>
Content-Type: application/json

{
  "price": 24.99,
  "stock": 45
}
```

#### Delete Product (Admin Only)
```bash
DELETE /v2/products/2
Authorization: Bearer <token>
```

### Legacy API (v1)

#### Hello World
```bash
GET /v1/example/hello
```

#### Long Request (5s delay)
```bash
GET /v1/example/longRequest
```

#### Authenticated Hello (Custom Auth)
```bash
GET /v1/example/auth/example
X-Auth: PASS
```

## Project Structure

```
.
├── app/
│   ├── controllers/      # Request handlers
│   │   ├── example.go
│   │   ├── ok.go
│   │   ├── product.go
│   │   └── user.go
│   ├── core/            # Core utilities
│   │   ├── config.go    # Configuration management
│   │   ├── database.go  # Database connection
│   │   ├── logger.go    # Logging utilities
│   │   └── shortcuts.go
│   ├── models/          # Data models
│   │   ├── product.go
│   │   └── user.go
│   ├── routes/          # Route definitions
│   │   ├── common.go
│   │   ├── example.go
│   │   ├── product.go
│   │   └── user.go
│   └── services/        # Business logic & middleware
│       ├── auth.go
│       ├── jwt.go
│       ├── jwt_middleware.go
│       └── rate_limiter.go
├── config/
│   └── app.toml        # Configuration file
├── tests/
│   └── main_test.go    # Integration tests
├── go.mod
└── main.go             # Application entry point
```

## Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with verbose output
go test -v ./...

# Using Makefile
make test
make test-cover
```

## Docker Support

```bash
# Build and run with docker-compose
docker-compose up -d

# Build manually
docker build -t my-gin-example .

# Run manually
docker run -p 8089:8089 -v $(pwd)/data:/root/data my-gin-example
```

## Testing with cURL

### Complete User Flow
```bash
# 1. Register a user
curl -X POST http://localhost:8089/v2/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "email": "test@example.com",
    "password": "password123",
    "full_name": "Test User"
  }'

# 2. Login and get token
TOKEN=$(curl -X POST http://localhost:8089/v2/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "password123"
  }' | jq -r '.token')

# 3. Get profile
curl -X GET http://localhost:8089/v2/users/profile \
  -H "Authorization: Bearer $TOKEN"

# 4. Create a product
curl -X POST http://localhost:8089/v2/products \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Test Product",
    "description": "A test product",
    "price": 99.99,
    "stock": 10,
    "category": "Test",
    "sku": "TEST-001"
  }'

# 5. List products
curl -X GET "http://localhost:8089/v2/products?page=1&page_size=10"
```

### WebSocket Support

#### Connect to WebSocket
```bash
# Use the provided HTML client
open examples/websocket_client.html

# Or connect programmatically
wscat -c ws://localhost:8089/ws/connect
```

The WebSocket server:
- Sends a welcome message on connection
- Echoes back any JSON messages received
- Sends periodic ping messages every 5 seconds
- Handles graceful disconnection

**Example WebSocket Message:**
```json
{
  "type": "message",
  "data": {
    "text": "Hello WebSocket!",
    "sender": "client"
  },
  "timestamp": 1701432000
}
```

## Key Technologies

- **[Gin](https://github.com/gin-gonic/gin)** v1.10.0 - Web framework
- **[GORM](https://gorm.io/)** v1.25.12 - ORM library
- **[JWT](https://github.com/golang-jwt/jwt)** v5.2.1 - JWT implementation
- **[Viper](https://github.com/spf13/viper)** v1.19.0 - Configuration management
- **[Zerolog](https://github.com/rs/zerolog)** v1.33.0 - Structured logging
- **[Testify](https://github.com/stretchr/testify)** v1.9.0 - Testing toolkit
- **[Gorilla WebSocket](https://github.com/gorilla/websocket)** v1.5.3 - WebSocket support

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is open source and available under the MIT License.