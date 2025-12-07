# my-gin-example

A comprehensive example of a production-ready RESTful API built with Gin framework, featuring modern Go practices, security best practices, JWT authentication, file upload, WebSocket support, and complete CI/CD integration.

## ✨ Features

### Core Features
- 🚀 **Modern Go 1.24** with latest stable dependencies
- 🔐 **JWT Authentication** with role-based access control
- 📦 **CRUD Operations** for users and products
- 🗄️ **GORM Integration** with SQLite (easily switchable to MySQL/PostgreSQL)
- 📁 **File Upload Management** with security validation
- 🔌 **WebSocket Support** for real-time communication
- ✅ **Request Validation** using Gin's binding

### Quality & Security
- 🔒 **Security Hardened** - gosec verified, zero vulnerabilities
- 🧪 **Unit Tests** with testify
- 📝 **Structured Logging** with zerolog
- ⚙️ **Configuration Management** with Viper
- 🚦 **Rate Limiting** middleware
- 🔍 **Pagination & Search** for list endpoints

### DevOps & CI/CD
- ✅ **GitHub Actions** - automated testing and security scanning
- 🔄 **CircleCI** ready configuration
- 🐳 **Docker & Docker Compose** support
- 📊 **CodeQL Security Analysis**
- 🛡️ **Dependabot** integration

## 🏆 Build Status

[![Go Build and Test](https://github.com/wmh/my-gin-example/actions/workflows/go.yml/badge.svg)](https://github.com/wmh/my-gin-example/actions/workflows/go.yml)
[![CodeQL](https://github.com/wmh/my-gin-example/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/wmh/my-gin-example/actions/workflows/codeql-analysis.yml)
[![Security Scanning](https://github.com/wmh/my-gin-example/actions/workflows/security-scan.yml/badge.svg)](https://github.com/wmh/my-gin-example/actions/workflows/security-scan.yml)

## 📋 Prerequisites

- Go 1.24 or higher
- Make (optional, for convenience commands)

## 🚀 Quick Start

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

## ⚙️ Configuration

Create or edit `config/app.toml`:

```toml
app_port = 8089

[database]
driver = "sqlite"
dsn = "./data/my_gin_example.db"

[jwt]
secret = "your-secret-key-here"
expires_hours = 24

[logs]
common_log = "./data/logs/common.log"
stdout_only = false

[upload]
max_size = 10485760  # 10MB
allowed_types = ["image/jpeg", "image/png", "image/gif", "application/pdf"]
```

## 📚 API Endpoints

### Authentication
- `POST /api/v1/auth/register` - Register new user
- `POST /api/v1/auth/login` - Login and get JWT token

### Users (Protected)
- `GET /api/v1/users` - List users (with pagination)
- `GET /api/v1/users/:id` - Get user details
- `PUT /api/v1/users/:id` - Update user
- `DELETE /api/v1/users/:id` - Delete user

### Products (Protected)
- `GET /api/v1/products` - List products (with pagination & search)
- `POST /api/v1/products` - Create product
- `GET /api/v1/products/:id` - Get product details
- `PUT /api/v1/products/:id` - Update product
- `DELETE /api/v1/products/:id` - Delete product

### File Upload (Protected)
- `POST /api/v1/files/upload` - Upload single file
- `POST /api/v1/files/batch-upload` - Upload multiple files
- `GET /api/v1/files` - List uploaded files
- `GET /api/v1/files/:id` - Get file details
- `GET /api/v1/files/:id/download` - Download file
- `DELETE /api/v1/files/:id` - Delete file

### WebSocket
- `GET /ws` - WebSocket connection for real-time updates

### Health Check
- `GET /ok` - Simple health check

## 🧪 Testing

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific package tests
go test ./app/services/...
```

## 🔒 Security

This project follows security best practices:

- ✅ **Zero vulnerabilities** - verified by govulncheck
- ✅ **gosec compliant** - all security warnings resolved
- ✅ **Path traversal protection** - file operations validated
- ✅ **Secure file permissions** - 0600 for files, 0750 for directories
- ✅ **HTTP timeout configuration** - protection against slowloris attacks
- ✅ **Regular dependency updates** - via Dependabot

Run security checks locally:

```bash
# Check for vulnerabilities
govulncheck ./...

# Run security scanner
gosec ./...
```

## 🐳 Docker

```bash
# Build Docker image
docker build -t my-gin-example .

# Run with Docker Compose
docker-compose up -d
```

## 🛠️ Development

### Using Make commands

```bash
# Run the application
make run

# Run tests
make test

# Build binary
make build

# Run with hot reload (using air)
make dev

# Clean build artifacts
make clean
```

### Project Structure

```
my-gin-example/
├── app/
│   ├── controllers/     # Request handlers
│   ├── core/           # Core functionality (config, db, logger)
│   ├── models/         # Database models
│   ├── routes/         # Route definitions
│   └── services/       # Business logic
├── config/             # Configuration files
├── data/              # Database and uploads
├── scripts/           # Utility scripts
├── tests/             # Integration tests
└── main.go            # Application entry point
```

## 📖 Documentation

- [FEATURES.md](FEATURES.md) - Detailed feature documentation
- [CHANGELOG.md](CHANGELOG.md) - Version history and changes
- [QUICKSTART.md](QUICKSTART.md) - Quick start guide

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is open source and available under the [MIT License](LICENSE).

## 🙏 Acknowledgments

Built with:
- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [GORM](https://gorm.io/)
- [JWT-Go](https://github.com/golang-jwt/jwt)
- [Zerolog](https://github.com/rs/zerolog)
- [Viper](https://github.com/spf13/viper)

## 📧 Contact

For questions or feedback, please open an issue on GitHub.

---

**Version:** 2.0.0  
**Go Version:** 1.24+  
**Last Updated:** December 2025
