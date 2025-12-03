# 🚀 Quick Start Guide

Get up and running with my-gin-example in 5 minutes!

## Prerequisites

- Go 1.23+ installed
- Git
- curl or Postman (for testing)

## Method 1: Run Locally (Recommended for Development)

### Step 1: Clone and Setup
```bash
git clone https://github.com/wmh/my-gin-example.git
cd my-gin-example
go mod download
```

### Step 2: Start the Server
```bash
go run main.go
```

You should see:
```
{"date":"2025-12-03","datetime":"2025-12-03 10:00:00","tag":"database","message":"Database connected successfully"}
```

Server is now running on `http://localhost:8089`

### Step 3: Test It
```bash
# Health check
curl http://localhost:8089/ok

# Should return: ok
```

### Step 4: Try the Full API
```bash
# Run all examples
./examples/api_examples.sh
```

## Method 2: Using Docker

### Quick Start
```bash
docker-compose up
```

That's it! Server runs on `http://localhost:8089`

### Build Your Own Image
```bash
docker build -t my-gin-example .
docker run -p 8089:8089 my-gin-example
```

## Method 3: Using Makefile

```bash
# Install dependencies
make install

# Run the server
make run

# Run tests
make test

# Build binary
make build

# Clean up
make clean
```

## 🎯 Your First API Call

### 1. Register a User
```bash
curl -X POST http://localhost:8089/v2/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "alice",
    "email": "alice@example.com",
    "password": "securepass123",
    "full_name": "Alice Wonder"
  }'
```

**Response:**
```json
{
  "message": "User registered successfully",
  "user": {
    "id": 1,
    "username": "alice",
    "email": "alice@example.com",
    "full_name": "Alice Wonder",
    "role": "user",
    "is_active": true,
    "created_at": "2025-12-03T10:00:00Z"
  }
}
```

### 2. Login
```bash
curl -X POST http://localhost:8089/v2/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "alice",
    "password": "securepass123"
  }'
```

**Copy the token from response!**

### 3. Create a Product (Protected)
```bash
export TOKEN="<paste-your-token-here>"

curl -X POST http://localhost:8089/v2/products \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Awesome Laptop",
    "description": "High-performance laptop",
    "price": 1299.99,
    "stock": 10,
    "category": "Electronics",
    "sku": "LAP-001"
  }'
```

### 4. List Products (Public)
```bash
curl http://localhost:8089/v2/products
```

### 5. Try WebSocket
Open `examples/websocket_client.html` in your browser and click "Connect"!

## 🎮 Interactive Testing

### Using the Example Script
```bash
chmod +x examples/api_examples.sh
./examples/api_examples.sh
```

This script will:
1. ✅ Check server health
2. ✅ Register a demo user
3. ✅ Login and get token
4. ✅ Create products
5. ✅ List and search products
6. ✅ Update products
7. ✅ Test all endpoints

### Using Postman

Import these as requests:

**Base URL:** `http://localhost:8089`

**Endpoints to try:**
```
GET  /ok
POST /v2/auth/register
POST /v2/auth/login
GET  /v2/users/profile (needs token)
PUT  /v2/users/profile (needs token)
GET  /v2/products
POST /v2/products (needs token)
GET  /v2/products/1
GET  /v1/example/hello
```

## 📝 Configuration

Edit `config/app.toml`:

```toml
app_port = 8089  # Change server port

[database]
path = "./data/app.db"  # Database location

[jwt]
secret = "your-secret-key-change-in-production"
expiration_hours = 24  # Token validity
```

## 🧪 Running Tests

```bash
# All tests
go test ./...

# With coverage
go test -cover ./...

# Specific package
go test ./app/controllers/...

# Verbose
go test -v ./...
```

## 🐛 Troubleshooting

### Port Already in Use
```bash
# Change port in config/app.toml
app_port = 8090

# Or set environment variable
APP_PORT=8090 go run main.go
```

### Database Error
```bash
# Ensure data directory exists
mkdir -p data

# Remove old database
rm data/app.db

# Restart server (it will auto-create)
go run main.go
```

### Module Issues
```bash
# Clean and download
go clean -modcache
go mod download
go mod tidy
```

### Tests Failing
```bash
# Clean test cache
go clean -testcache

# Run tests again
go test ./...
```

## 📚 What's Next?

### Explore the Code
```
app/controllers/  - Request handlers
app/models/       - Data structures
app/services/     - Business logic
app/routes/       - Route definitions
```

### Read Documentation
- `README.md` - Complete API documentation
- `FEATURES.md` - Feature overview
- `CHANGELOG.md` - Version history
- `examples/README.md` - Example usage

### Customize
1. Add your own models in `app/models/`
2. Create controllers in `app/controllers/`
3. Define routes in `app/routes/`
4. Add middleware in `app/services/`

### Deploy
```bash
# Build for production
go build -o app main.go

# Run
./app

# Or use Docker
docker-compose up -d
```

## 🎓 Learning Path

1. **Day 1**: Run the server, test endpoints
2. **Day 2**: Explore authentication flow
3. **Day 3**: Study CRUD operations
4. **Day 4**: Understand middleware
5. **Day 5**: Try WebSocket
6. **Day 6**: Add your own feature
7. **Day 7**: Deploy to production

## 🆘 Need Help?

- Check `README.md` for detailed docs
- Review `examples/` for code samples
- Read inline code comments
- Check the tests for usage examples

## 🎉 Success!

You now have:
- ✅ Running web server
- ✅ REST API with auth
- ✅ Database with GORM
- ✅ WebSocket support
- ✅ Rate limiting
- ✅ Full test suite
- ✅ Docker setup

Happy coding! 🚀
